package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type ports struct {
	meSession *MeMetrics

	up              descMétrica
	controller      descMétrica
	portType        descMétrica
	status          descMétrica
	actualSpeed     descMétrica
	configuredSpeed descMétrica
	health          descMétrica

	//IscsiPort
	sfpStatus  descMétrica
	sfpPresent descMétrica

	logger log.Logger
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
				NomeMetrica("port", "controller"),
				"Port controller", []string{"port", "controller"}),
		},
		portType: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "port_type"),
				"Port type", []string{"port", "portType"}),
		},
		status: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "status"),
				"Port status", []string{"port", "status"}),
		},
		actualSpeed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "actual_speed"),
				"Port actual speed", []string{"port", "speed"}),
		},
		configuredSpeed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "configured_speed"),
				"Port configured speed", []string{"port", "speed"}),
		},
		health: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "health"),
				"Port health", []string{"port", "health"}),
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
		ch <- p.controller.constMetric(float64(port.ControllerNumeric), port.Port, port.Controller)
		ch <- p.portType.constMetric(float64(port.PortTypeNumeric), port.Port, port.PortType)
		ch <- p.status.constMetric(float64(port.StatusNumeric), port.Port, port.Status)
		ch <- p.actualSpeed.constMetric(float64(port.ActualSpeedNumeric), port.Port, port.ActualSpeed)
		ch <- p.configuredSpeed.constMetric(float64(port.ConfiguredSpeedNumeric), port.Port, port.ConfiguredSpeed)
		ch <- p.health.constMetric(float64(port.HealthNumeric), port.Port, port.Health)
		//IscsiPort
		for _, iscsiPort := range port.IscsiPort {
			ch <- p.sfpStatus.constMetric(1, port.Port, iscsiPort.SfpStatus)
			ch <- p.sfpPresent.constMetric(1, port.Port, iscsiPort.SfpPresent)
		}
	}

	return nil
}
