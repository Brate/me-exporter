package collector

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"io"
	"me_exporter/Me"
	"me_exporter/config"
	"net/http"
	"sync"
)

const (
	namespace = "me"
)

var (
	//metrics            MeMetrics
	factories          = make(map[string]factoryType)
	coletores          = make(map[string]Coletor)
	mutexInitColetores = sync.Mutex{}
)

type factoryType func(me *MeMetrics, logger log.Logger) (Coletor, error)
type Coletor interface {
	Update(ch chan<- prometheus.Metric) error
}

type descMétrica struct {
	tipo prometheus.ValueType
	desc *prometheus.Desc
}

func registerCollector(name string, factory factoryType) {
	factories[name] = factory
}

func (d *descMétrica) mustNewConstMetric(v float64, labels ...string) prometheus.Metric {
	m, err := prometheus.NewConstMetric(d.desc, d.tipo, v, labels...)
	if err != nil {
		panic(err)
	}
	return m
}

func NewDescritor(name string, help string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(name, help, labels, nil)
}

func NomeMetrica(subsystem string, name string) string {
	return prometheus.BuildFQName(namespace, subsystem, name)
}

type MeMetrics struct {
	baseUrl    string
	sessionKey string
	serviceTag Me.ServiceTagInfo

	controllerStatistics Me.ControllerStatistics
	cacheSettings        Me.SystemCacheSettings
	diskGroupStatistics  Me.DiskGroupStatistics
	diskGroups           Me.Disk
	diskStatistic        Me.DiskStatistic
	disks                Me.Drives
	enclosures           Me.Enclosures
	expanderStatus       Me.SasStatusControllerA
	fans                 Me.Fans
	frus                 Me.EnclosureFru
	pools                Me.Pools
	poolStatistics       Me.PoolStatistics
	ports                Me.Ports
	sensorStatus         Me.SensorStatus
	volumes              Me.Volumes
	volumeStatistics     Me.VolumeStatistics
	tierStatistics       Me.TierStatistics
	tiers                Me.Tiers
	unwritableCache      Me.UnwritableCache

	logger log.Logger
}

func NewMeMetrics(instance string, logger log.Logger) (me *MeMetrics) {
	me = &MeMetrics{logger: logger}

	aEntry := config.ExporterConfig.FindAuthByInstance(instance)
	if aEntry == nil {
		return
	}

	me.baseUrl = aEntry.Url
	if err := me.Login(*aEntry); err != nil {
		level.Error(logger).Log("msg", "Login error on ME",
			"instance", instance, "error", err)
		return
	}

	return
}

type MeCollector struct {
	Coletores map[string]Coletor
	logger    log.Logger
}

func NewMECollector(me *MeMetrics, logger log.Logger) (*MeCollector, error) {
	mutexInitColetores.Lock()
	defer mutexInitColetores.Unlock()

	lColetores := make(map[string]Coletor)
	for name, factory := range factories {
		if col, ok := coletores[name]; ok {
			lColetores[name] = col
		} else {
			c, err := factory(me, log.With(logger, "coletor", name))
			if err != nil {
				return nil, err
			}
			coletores[name] = c
			lColetores[name] = c
		}
	}

	return &MeCollector{Coletores: lColetores, logger: logger}, nil
}

func FlushMECollectors() {
	mutexInitColetores.Lock()
	defer mutexInitColetores.Unlock()

	coletores = make(map[string]Coletor)
}

func (c *MeCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(factories))
	for name, coletor := range coletores {
		go func(name string, col Coletor) {
			execute(name, col, ch, c.logger)
			wg.Done()
		}(name, coletor)
	}
	wg.Wait()
}

func (c *MeCollector) Describe(ch chan<- *prometheus.Desc) {

}

func execute(name string, c Coletor, ch chan<- prometheus.Metric, logger log.Logger) {
	c.Update(ch)
}

// _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

func (me *MeMetrics) FetchMetrics() (instance config.AuthEntry, err error) {
	err = me.Login(instance)

	var metricas []func() error
	metricas = append(metricas, me.ServiceTag, me.CacheSettings, me.DiskGroupStatistics,
		me.DiskGroups, me.DiskStatistics, me.Disks, me.Enclosures, me.ExpanderStatus,
		me.Fans, me.Frus, me.Pools, me.PoolsStatistics, me.Ports, me.SensorStatus,
		me.ControllerStatistics, me.Volumes, me.VolumeStatistics, me.Tiers,
		me.TierStatistics, me.UnwritableCache)

	for _, metrica := range metricas {
		err = metrica()
		if err != nil {
			return
		}
	}

	return
}

func (me *MeMetrics) Login(instance config.AuthEntry) (err error) {
	if me.sessionKey != "" {
		return
	}

	url := fmt.Sprintf("%v/login/%v", me.baseUrl,
		instance.Hash)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	client := me.NewClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	httpLogin := struct {
		Status []Me.Status `json:"status"`
	}{}
	err = json.Unmarshal(body, &httpLogin)
	if err != nil {
		fmt.Printf("Erro ao deserializar %v", err)
		return err
	}

	switch httpLogin.Status[0].ResponseTypeNumeric {
	case 0:
		me.sessionKey = httpLogin.Status[0].Response
	case 1:
		// TODO: Testar problemas de login
		me.sessionKey = ""
		err = fmt.Errorf("erro ao logar em %v", instance.Instance)
	}

	return
}
func (me *MeMetrics) ServiceTag() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/service-tag-info", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4ServiceTagInfoFromRequest(client, req)
	if err == nil {
		me.serviceTag = xxx[0]
	}
	return
}
func (me *MeMetrics) CacheSettings() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/cache-parameters", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4CacheSettingsFromRequest(client, req)
	if err == nil {
		me.cacheSettings = xxx[0]
	}
	return
}
func (me *MeMetrics) DiskGroupStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/disk-group-statistics", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4DiskGroupStatisticsFromRequest(client, req)
	if err == nil {
		me.diskGroupStatistics = xxx[0]
	}
	return
}
func (me *MeMetrics) DiskGroups() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/disk-groups", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4DiskGroupsFromRequest(client, req)
	if err == nil {
		me.diskGroups = xxx[0]
	}
	return
}
func (me *MeMetrics) DiskStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/disk-statistics", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4DiskStatisticFromRequest(client, req)
	if err == nil {
		me.diskStatistic = xxx[0]
	}
	return
}
func (me *MeMetrics) Disks() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/disks", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4DisksFromRequest(client, req)
	if err == nil {
		me.disks = xxx[0]
	}
	return
}
func (me *MeMetrics) Enclosures() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/enclosures", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4EnclosuresFromRequest(client, req)
	if err == nil {
		me.enclosures = xxx[0]
	}
	return
}
func (me *MeMetrics) ExpanderStatus() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/expander-status", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4ExpanderStatusFromRequest(client, req)
	if err == nil {
		me.expanderStatus = xxx[0]
	}
	return
}
func (me *MeMetrics) Fans() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/fans", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4FansFromRequest(client, req)
	if err == nil {
		me.fans = xxx[0]
	}
	return
}
func (me *MeMetrics) Frus() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/frus", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	xxx, err := Me.NewMe4FrusFromRequest(client, req)
	if err == nil {
		me.frus = xxx[0]
	}
	return
}
func (me *MeMetrics) Pools() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/pools", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	xxx, err := Me.NewMe4PoolsFromRequest(client, req)
	if err == nil {
		me.pools = xxx[0]
	}
	return
}
func (me *MeMetrics) PoolsStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/pool-statistics", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	xxx, err := Me.NewMe4ShowPoolStatisticsFromRequest(client, req)
	if err == nil {
		me.poolStatistics = xxx[0]
	}
	return
}
func (me *MeMetrics) Ports() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/ports", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	xxx, err := Me.NewMe4PortsFromRequest(client, req)
	if err == nil {
		me.ports = xxx[0]
	}
	return
}
func (me *MeMetrics) SensorStatus() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/sensor-status", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	xxx, err := Me.NewMe4SensorStatusFromRequest(client, req)
	if err == nil {
		me.sensorStatus = xxx[0]
	}
	return
}
func (me *MeMetrics) ControllerStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	//url := fmt.Sprintf("%v/show/controller-statistics", me.baseUrl)
	url := fmt.Sprintf("%v/show/controller-statistics", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	xxx, err := Me.NewMe4ControllerStatisticsFromRequest(client, req)
	if err == nil {
		me.controllerStatistics = xxx[0]
	}
	return
}
func (me *MeMetrics) Volumes() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/volumes", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	xxx, err := Me.NewMe4VolumesFromRequest(client, req)
	if err == nil {
		me.volumes = xxx[0]
	}
	return
}
func (me *MeMetrics) VolumeStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	//url := fmt.Sprintf("%v/volume-statistics", me.baseUrl)
	url := fmt.Sprintf("%v/show/volume-statistics", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	//xxx, err := Me.NewMe4VolumeStatisticsFromRequest(client, req)
	xxx, err := Me.NewMe4VolumeStatisticsFromRequest(client, req)
	if err == nil {
		me.volumeStatistics = xxx[0]
	}
	return
}
func (me *MeMetrics) TierStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	//url := fmt.Sprintf("%v/tier-statistics", me.baseUrl)
	url := fmt.Sprintf("%v/show/tier-statistics/tier/all", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	//xxx, err := Me.NewMe4TierStatisticsFromRequest(client, req)
	xxx, err := Me.NewMe4TierStatisticsFromRequest(client, req)
	if err == nil {
		me.tierStatistics = xxx[0]
	}
	return
}
func (me *MeMetrics) Tiers() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	//url := fmt.Sprintf("%v/tiers", me.baseUrl)
	url := fmt.Sprintf("%v/show/tiers/tier/all", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	//xxx, err := Me.NewMe4TiersFromRequest(client, req)
	xxx, err := Me.NewMe4TiersFromRequest(client, req)
	if err == nil {
		me.tiers = xxx[0]
	}
	return
}
func (me *MeMetrics) UnwritableCache() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	client := me.NewClient()
	url := fmt.Sprintf("%v/show/unwritable-cache", me.baseUrl)
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
	}

	xxx, err := Me.NewMe4UnwritableCacheFromRequest(client, req)
	if err == nil {
		me.unwritableCache = xxx[0]
	}
	return
}

// _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
// Helpers
func (me *MeMetrics) Me4Request(url string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("dataType", "json")
	if me.sessionKey != "" {
		req.Header.Add("sessionKey", me.sessionKey)
	}
	return
}
func (me *MeMetrics) NewClient() (client *http.Client) {
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return
}
