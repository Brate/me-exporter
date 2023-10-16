package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	h "me_exporter/app/helpers"
)

type fans struct {
	//All used name in your labels
	meSession        *MeMetrics
	up               descMétrica // descMétrica 1, labels name, location and ExtendedStatus
	statusSesNumeric descMétrica // Enum, with labels name, StatusSes
	statusNumeric    descMétrica // Enum, with labels name, Status
	speed            descMétrica // descMétrica speed, descipiton is speed in RPM, labels name, Speed
	positionNumeric  descMétrica // Enum, with labels name, Position
	health           descMétrica // Enum, with labels name, Health, HealthReason and HealthRecommendation
	logger           log.Logger
}

func init() {
	registerCollector("fans", NewFansCollector)
}

func NewFansCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &fans{
		meSession: me,
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "fan_up"),
				"Up", []string{"name", "location", "extended_status"}),
		},
		statusSesNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "status_ses_numeric"),
				"Status SES", []string{"name", "status_ses"}),
		},
		statusNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "status_numeric"),
				"Status", []string{"name", "status"}),
		},
		speed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "speed"),
				"Speed in RPM", []string{"name", "speed"}),
		},
		positionNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "position_numeric"),
				"Position", []string{"name", "position"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fan", "health"),
				"Health", []string{"name", "health", "health_reason", "health_recommendation"}),
		},
		logger: logger,
	}, nil
}

func (f fans) Update(ch chan<- prometheus.Metric) error {
	err := f.meSession.Fans()
	if err != nil {
		return err
	}
	is := h.IntToString

	//TODO revisar o uso do constMetric que está no enclosure

	for _, fan := range f.meSession.fans {
		ch <- f.up.constMetric(1, fan.Name, fan.Location, fan.ExtendedStatus)
		ch <- f.statusSesNumeric.constMetric(float64(fan.StatusSesNumeric), fan.Name, fan.StatusSes)
		ch <- f.statusNumeric.constMetric(float64(fan.StatusNumeric), fan.Name, fan.Status)
		//Preciso dessa label fan_Speed? acho que não
		ch <- f.speed.constMetric(float64(fan.Speed), fan.Name, is(fan.Speed))
		ch <- f.positionNumeric.constMetric(float64(fan.PositionNumeric), fan.Name, fan.Position)
		ch <- f.health.constMetric(float64(fan.HealthNumeric), fan.Name, fan.Health, fan.HealthReason, fan.HealthRecommendation)
	}

	return nil
}
