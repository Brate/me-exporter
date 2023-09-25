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
	"regexp"
	"sync"
	"time"
)

const (
	namespace = "me"
)

var (
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
	diskGroupsStatistics []Me.DiskGroupStatistics
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

	pool := make(chan struct{}, 4) // 4 workers
	for name, coletor := range coletores {
		go func(name string, col Coletor) {
			pool <- struct{}{}
			execute(col, ch)
			<-pool // release a worker
			wg.Done()
		}(name, coletor)
	}
	wg.Wait()
}

func (c *MeCollector) Describe(ch chan<- *prometheus.Desc) {

}

func execute(c Coletor, ch chan<- prometheus.Metric) {
	c.Update(ch)
}

// _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

//func (me *MeMetrics) FetchMetrics() (instance config.AuthEntry, err error) {
//	err = me.Login(instance)
//
//	var metricas []func() error
//	metricas = append(metricas, me.ServiceTag, me.CacheSettings, me.DiskGroupStatistics,
//		me.DiskGroups, me.DiskStatistics, me.Disks, me.Enclosures, me.ExpanderStatus,
//		me.Fans, me.Frus, me.Pools, me.PoolsStatistics, me.Ports, me.SensorStatus,
//		me.ControllerStatistics, me.Volumes, me.VolumeStatistics, me.Tiers,
//		me.TierStatistics, me.UnwritableCache)
//
//	for _, metrica := range metricas {
//		err = metrica()
//		if err != nil {
//			return
//		}
//	}
//
//	return
//}

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

	url := fmt.Sprintf("%v/show/service-tag-info", me.baseUrl)
	body, err := me.ClientDo(url)

	st, err := Me.NewMe4ServiceTagInfoFrom(body)
	if err == nil {
		me.serviceTag = st[0]
	}
	return
}
func (me *MeMetrics) CacheSettings() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/cache-parameters", me.baseUrl)

	body, err := me.ClientDo(url)

	cs, err := Me.NewMe4CacheSettingsFrom(body)
	if err == nil {
		me.cacheSettings = cs[0]
	}
	return
}
func (me *MeMetrics) DiskGroupStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disk-group-statistics", me.baseUrl)
	body, err := me.ClientDo(url)

	stats, err := Me.NewMe4DiskGroupStatisticsFrom(body)
	if err == nil {
		me.diskGroupsStatistics = stats
	}
	return
}
func (me *MeMetrics) DiskGroups() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disk-groups", me.baseUrl)
	body, err := me.ClientDo(url)

	dg, err := Me.NewMe4DiskGroupsFrom(body)
	if err == nil {
		me.diskGroups = dg[0]
	}
	return
}
func (me *MeMetrics) DiskStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disk-statistics", me.baseUrl)
	body, err := me.ClientDo(url)

	ds, err := Me.NewMe4DiskStatisticsFrom(body)
	if err == nil {
		me.diskStatistic = ds[0]
	}
	return
}
func (me *MeMetrics) Disks() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disks", me.baseUrl)
	body, err := me.ClientDo(url)

	disks, err := Me.NewMe4DisksFrom(body)
	if err == nil {
		me.disks = disks[0]
	}
	return
}
func (me *MeMetrics) Enclosures() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/enclosures", me.baseUrl)
	body, err := me.ClientDo(url)

	enclosures, err := Me.NewMe4EnclosuresFrom(body)
	if err == nil {
		me.enclosures = enclosures[0]
	}
	return
}
func (me *MeMetrics) ExpanderStatus() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/expander-status", me.baseUrl)
	body, err := me.ClientDo(url)

	expanders, err := Me.NewMe4ExpanderStatusFrom(body)
	if err == nil {
		me.expanderStatus = expanders[0]
	}
	return
}
func (me *MeMetrics) Fans() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/fans", me.baseUrl)
	body, err := me.ClientDo(url)

	fans, err := Me.NewMe4FansFrom(body)
	if err == nil {
		me.fans = fans[0]
	}
	return
}
func (me *MeMetrics) Frus() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/frus", me.baseUrl)
	body, err := me.ClientDo(url)

	frus, err := Me.NewMe4FrusFrom(body)
	if err == nil {
		me.frus = frus[0]
	}
	return
}
func (me *MeMetrics) Pools() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/pools", me.baseUrl)
	body, err := me.ClientDo(url)

	// Corrige JSON mal formado:
	// duas ocorrencias de disk-groups no mesmo objeto

	regex := regexp.MustCompile(`"disk-groups":(\s*\d+)`)
	body = regex.ReplaceAll(body, []byte(`"disk-groups-count":$1`))

	pools, err := Me.NewMe4PoolsFrom(body)
	if err == nil {
		me.pools = pools[0]
	}
	return
}
func (me *MeMetrics) PoolsStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/pool-statistics", me.baseUrl)
	body, err := me.ClientDo(url)

	stats, err := Me.NewMe4ShowPoolStatisticsFrom(body)
	if err == nil {
		me.poolStatistics = stats[0]
	}
	return
}
func (me *MeMetrics) Ports() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/ports", me.baseUrl)
	body, err := me.ClientDo(url)

	ports, err := Me.NewMe4PortsFrom(body)
	if err == nil {
		me.ports = ports[0]
	}
	return
}
func (me *MeMetrics) SensorStatus() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/sensor-status", me.baseUrl)
	body, err := me.ClientDo(url)

	status, err := Me.NewMe4SensorStatusFrom(body)
	if err == nil {
		me.sensorStatus = status[0]
	}
	return
}
func (me *MeMetrics) ControllerStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/controller-statistics", me.baseUrl)
	body, err := me.ClientDo(url)

	stats, err := Me.NewMe4ControllerStatisticsFrom(body)
	if err == nil {
		me.controllerStatistics = stats[0]
	}
	return
}
func (me *MeMetrics) Volumes() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/volumes", me.baseUrl)
	body, err := me.ClientDo(url)

	volumes, err := Me.NewMe4VolumesFrom(body)
	if err == nil {
		me.volumes = volumes[0]
	}
	return
}
func (me *MeMetrics) VolumeStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/volume-statistics", me.baseUrl)
	body, err := me.ClientDo(url)

	stats, err := Me.NewMe4VolumeStatisticsFrom(body)
	if err == nil {
		me.volumeStatistics = stats[0]
	}
	return
}
func (me *MeMetrics) TierStatistics() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/tier-statistics/tier/all", me.baseUrl)
	body, err := me.ClientDo(url)

	stats, err := Me.NewMe4TierStatisticsFrom(body)
	if err == nil {
		me.tierStatistics = stats[0]
	}
	return
}
func (me *MeMetrics) Tiers() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/tiers/tier/all", me.baseUrl)
	body, err := me.ClientDo(url)

	tiers, err := Me.NewMe4TiersFrom(body)
	if err == nil {
		me.tiers = tiers[0]
	}
	return
}
func (me *MeMetrics) UnwritableCache() (err error) {
	if me.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/unwritable-cache", me.baseUrl)
	body, err := me.ClientDo(url)

	cache, err := Me.NewMe4UnwritableCacheFrom(body)
	if err == nil {
		me.unwritableCache = cache[0]
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
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return
}
func (me *MeMetrics) ClientDo(url string) (body []byte, err error) {
	client := me.NewClient()
	req, err := me.Me4Request(url)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	msg := fmt.Sprintf("requesting %v", url)
	_ = level.Info(me.logger).Log("msg", msg)
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(me.logger).Log("msg", "request error", "error", err)
		return
	}

	body, _ = io.ReadAll(resp.Body)
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			_ = level.Error(me.logger).Log("msg", "body.Close error", "error", err)
		}
	}()

	return
}
