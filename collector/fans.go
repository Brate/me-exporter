package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type fans struct {
	meSession *MeMetrics
	up        descMétrica
	statusSes descMétrica
	status    descMétrica
	speed     descMétrica
	health    descMétrica
	logger    log.Logger
}

func init() {
	registerCollector("fans", NewFansCollector)
}

func NewFansCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &fans{
		meSession: me,
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "up"),
				"Up", []string{"id", "name", "location"}),
		},
		statusSes: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "status_ses"),
				"Status SES", []string{"id", "status"}),
		},
		status: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "status"),
				"Status", []string{"id", "status"}),
		},
		speed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "speed"),
				"Speed in RPM", []string{"id"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "health"),
				"Health", []string{"id", "health"}),
		},
		logger: logger,
	}, nil
}

func (f fans) Update(ch chan<- prometheus.Metric) error {
	err := f.meSession.Fans()
	if err != nil {
		return err
	}

	for _, fan := range f.meSession.fans {
		ch <- f.up.constMetric(1, fan.DurableID, fan.Name, fan.Location)
		ch <- f.status.constMetric(float64(fan.StatusNumeric), fan.DurableID, fan.Status)
		ch <- f.statusSes.constMetric(float64(fan.StatusSesNumeric), fan.DurableID, fan.StatusSes)
		ch <- f.speed.constMetric(float64(fan.Speed), fan.DurableID)
		ch <- f.health.constMetric(float64(fan.HealthNumeric), fan.DurableID, fan.Health)
	}

	return nil
}
