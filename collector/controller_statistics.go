package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type controllerStatisticsCollector struct {
	meSession *MeMetrics
	cpuLoad   descMétrica
	logger    log.Logger
}

func init() {
	registerCollector("controller", NewControllerStatisticsCollector)
}

func NewControllerStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &controllerStatisticsCollector{
		meSession: me,
		cpuLoad: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "cpu_load"),
				"CPU load on the controller", []string{"controller"}),
		},
		logger: logger,
	}, nil
}

func (c *controllerStatisticsCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.meSession.ControllerStatistics(); err != nil {
		return err
	}

	stats := c.meSession.controllerStatistics
	ch <- prometheus.MustNewConstMetric(
		c.cpuLoad.desc, c.cpuLoad.tipo, float64(stats.CPULoad), stats.DurableID,
	)

	return nil
}
