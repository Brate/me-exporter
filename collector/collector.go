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
	diskGroups           []Me.Disk
	diskStatistic        []Me.DiskStatistic
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
		_ = level.Error(logger).Log("msg", "Login error on ME",
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
	_ = c.Update(ch)
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

func (meMetrics *MeMetrics) Login(instance config.AuthEntry) (err error) {
	if meMetrics.sessionKey != "" {
		return
	}

	url := fmt.Sprintf("%v/login/%v", meMetrics.baseUrl,
		instance.Hash)
	req, err := meMetrics.Me4Request(url)
	if err != nil {
		_ = level.Error(meMetrics.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	client := meMetrics.NewClient()
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
		meMetrics.sessionKey = httpLogin.Status[0].Response
	case 1:
		// TODO: Testar problemas de login
		meMetrics.sessionKey = ""
		err = fmt.Errorf("erro ao logar em %v", instance.Instance)
	}

	return
}
func (meMetrics *MeMetrics) ServiceTag() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/service-tag-info", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	st, err := Me.NewMe4ServiceTagInfoFrom(body)
	if err == nil {
		meMetrics.serviceTag = st[0]
	}
	return
}
func (meMetrics *MeMetrics) CacheSettings() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/cache-parameters", meMetrics.baseUrl)

	body, err := meMetrics.ClientDo(url)

	cs, err := Me.NewMe4CacheSettingsFrom(body)
	if err == nil {
		meMetrics.cacheSettings = cs[0]
	}
	return
}
func (meMetrics *MeMetrics) DiskGroupStatistics() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disk-group-statistics", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	stats, err := Me.NewMe4DiskGroupStatisticsFrom(body)
	if err == nil {
		meMetrics.diskGroupsStatistics = stats
	}
	return
}
func (meMetrics *MeMetrics) DiskGroups() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disk-groups", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	dg, err := Me.NewMe4DiskGroupsFrom(body)
	if err == nil {
		meMetrics.diskGroups = dg
	}
	return
}
func (meMetrics *MeMetrics) DiskStatistics() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disk-statistics", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	ds, err := Me.NewMe4DiskStatisticsFrom(body)
	if err == nil {
		meMetrics.diskStatistic = ds
	}
	return
}
func (meMetrics *MeMetrics) Disks() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/disks", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	disks, err := Me.NewMe4DisksFrom(body)
	if err == nil {
		meMetrics.disks = disks[0]
	}
	return
}
func (meMetrics *MeMetrics) Enclosures() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/enclosures", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	enclosures, err := Me.NewMe4EnclosuresFrom(body)
	if err == nil {
		meMetrics.enclosures = enclosures[0]
	}
	return
}
func (meMetrics *MeMetrics) ExpanderStatus() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/expander-status", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	expanders, err := Me.NewMe4ExpanderStatusFrom(body)
	if err == nil {
		meMetrics.expanderStatus = expanders[0]
	}
	return
}
func (meMetrics *MeMetrics) Fans() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/fans", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	fans, err := Me.NewMe4FansFrom(body)
	if err == nil {
		meMetrics.fans = fans[0]
	}
	return
}
func (meMetrics *MeMetrics) Frus() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/frus", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	frus, err := Me.NewMe4FrusFrom(body)
	if err == nil {
		meMetrics.frus = frus[0]
	}
	return
}
func (meMetrics *MeMetrics) Pools() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/pools", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	// Corrige JSON mal formado:
	// duas ocorrencias de disk-groups no mesmo objeto

	regex := regexp.MustCompile(`"disk-groups":(\s*\d+)`)
	body = regex.ReplaceAll(body, []byte(`"disk-groups-count":$1`))

	pools, err := Me.NewMe4PoolsFrom(body)
	if err == nil {
		meMetrics.pools = pools[0]
	}
	return
}
func (meMetrics *MeMetrics) PoolsStatistics() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/pool-statistics", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	stats, err := Me.NewMe4ShowPoolStatisticsFrom(body)
	if err == nil {
		meMetrics.poolStatistics = stats[0]
	}
	return
}
func (meMetrics *MeMetrics) Ports() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/ports", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	ports, err := Me.NewMe4PortsFrom(body)
	if err == nil {
		meMetrics.ports = ports[0]
	}
	return
}
func (meMetrics *MeMetrics) SensorStatus() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/sensor-status", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	status, err := Me.NewMe4SensorStatusFrom(body)
	if err == nil {
		meMetrics.sensorStatus = status[0]
	}
	return
}
func (meMetrics *MeMetrics) ControllerStatistics() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/controller-statistics", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	stats, err := Me.NewMe4ControllerStatisticsFrom(body)
	if err == nil {
		meMetrics.controllerStatistics = stats[0]
	}
	return
}
func (meMetrics *MeMetrics) Volumes() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/volumes", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	volumes, err := Me.NewMe4VolumesFrom(body)
	if err == nil {
		meMetrics.volumes = volumes[0]
	}
	return
}
func (meMetrics *MeMetrics) VolumeStatistics() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/volume-statistics", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	stats, err := Me.NewMe4VolumeStatisticsFrom(body)
	if err == nil {
		meMetrics.volumeStatistics = stats[0]
	}
	return
}
func (meMetrics *MeMetrics) TierStatistics() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/tier-statistics/tier/all", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	stats, err := Me.NewMe4TierStatisticsFrom(body)
	if err == nil {
		meMetrics.tierStatistics = stats[0]
	}
	return
}
func (meMetrics *MeMetrics) Tiers() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/tiers/tier/all", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	tiers, err := Me.NewMe4TiersFrom(body)
	if err == nil {
		meMetrics.tiers = tiers[0]
	}
	return
}
func (meMetrics *MeMetrics) UnwritableCache() (err error) {
	if meMetrics.sessionKey == "" {
		return fmt.Errorf("invalid session")
	}

	url := fmt.Sprintf("%v/show/unwritable-cache", meMetrics.baseUrl)
	body, err := meMetrics.ClientDo(url)

	cache, err := Me.NewMe4UnwritableCacheFrom(body)
	if err == nil {
		meMetrics.unwritableCache = cache[0]
	}
	return
}

// Me4Request _-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
// Helpers
func (meMetrics *MeMetrics) Me4Request(url string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("dataType", "json")
	if meMetrics.sessionKey != "" {
		req.Header.Add("sessionKey", meMetrics.sessionKey)
	}
	return
}
func (meMetrics *MeMetrics) NewClient() (client *http.Client) {
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 5 * time.Second,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return
}
func (meMetrics *MeMetrics) ClientDo(url string) (body []byte, err error) {
	client := meMetrics.NewClient()
	req, err := meMetrics.Me4Request(url)
	if err != nil {
		_ = level.Error(meMetrics.logger).Log("msg", "Erro ao criar request", "error", err)
		return
	}

	msg := fmt.Sprintf("requesting %v", url)
	_ = level.Info(meMetrics.logger).Log("msg", msg)
	resp, err := client.Do(req)
	if err != nil {
		_ = level.Error(meMetrics.logger).Log("msg", "request error", "error", err)
		return
	}

	body, _ = io.ReadAll(resp.Body)
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			_ = level.Error(meMetrics.logger).Log("msg", "body.Close error", "error", err)
		}
	}()

	return
}
