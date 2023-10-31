package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type ports struct {
	//All metrics have port in your metrics
	meSession       *MeMetrics
	up              descMétrica //descMétrica is 1, and your labels is port, portType, media
	controller      descMétrica
	portType        descMétrica
	status          descMétrica
	actualSpeed     descMétrica
	configuredSpeed descMétrica
	health          descMétrica
	logger          log.Logger
	//IscsiPort
	sfpStatus  descMétrica
	sfpPresent descMétrica
}

func init() {
	registerCollector("ports", NewPortsCollector)
}

func NewPortsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &ports{
		meSession: me,
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "up"),
				"Port literals", []string{"port", "portType", "media"}),
		},
		controller: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "controller_numeric"),
				"Port controller numeric", []string{"port"}),
		},
		portType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "port_type_numeric"),
				"Port type numeric", []string{"port"}),
		},
		status: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "status_numeric"),
				"Port status numeric", []string{"port"}),
		},
		actualSpeed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "actual_speed_numeric"),
				"Port actual speed numeric", []string{"port"}),
		},
		configuredSpeed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "configured_speed_numeric"),
				"Port configured speed numeric", []string{"port"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "health_numeric"),
				"Port health numeric", []string{"port", "health"}),
		},
		//IscsiPort
		sfpStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "sfp_status"),
				"Port sfp status", []string{"port", "sfp_status"}),
		},
		sfpPresent: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "sfp_present"),
				"Port sfp present", []string{"port", "sfp_present"}),
		},
		logger: logger,
	}, nil
}

func (p ports) Update(ch chan<- prometheus.Metric) error {
	if err := p.meSession.Ports(); err != nil {
		return err
	}

	for _, port := range p.meSession.ports {
		ch <- p.up.constMetric(1, port.Port, port.PortType, port.Media)
		ch <- p.controller.constMetric(float64(port.ControllerNumeric), port.Port)
		ch <- p.portType.constMetric(float64(port.PortTypeNumeric), port.Port)
		ch <- p.status.constMetric(float64(port.StatusNumeric), port.Port)
		ch <- p.actualSpeed.constMetric(float64(port.ActualSpeedNumeric), port.Port)
		ch <- p.configuredSpeed.constMetric(float64(port.ConfiguredSpeedNumeric), port.Port)
		ch <- p.health.constMetric(float64(port.HealthNumeric), port.Port, port.Health)
		//IscsiPort
		for _, iscsiPort := range port.IscsiPort {
			ch <- p.sfpStatus.constMetric(1, port.Port, iscsiPort.SfpStatus)
			ch <- p.sfpPresent.constMetric(1, port.Port, iscsiPort.SfpPresent)
		}
	}

	return nil
}
