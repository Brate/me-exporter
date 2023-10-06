package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type serviceTagInfo struct {
	meSession   *MeMetrics
	enclosureID descMétrica
	logger      log.Logger
}

func init() {
	registerCollector("service_tag_info", NewServiceTagInfoCollector)
}

func NewServiceTagInfoCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &serviceTagInfo{
		meSession: me,
		enclosureID: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("service_tag_info_enclosure_id", "service_tag_info"),
				"Enclosure ID", []string{"service_tag"},
			),
		},
		logger: logger,
	}, nil

}

func (s serviceTagInfo) Update(ch chan<- prometheus.Metric) error {
	if err := s.meSession.ServiceTag(); err != nil {
		return err
	}

	for _, sti := range s.meSession.serviceTag {
		ch <- prometheus.MustNewConstMetric(s.enclosureID.desc, s.enclosureID.tipo,
			float64(sti.EnclosureID), sti.ServiceTag)
	}

	return nil
}
