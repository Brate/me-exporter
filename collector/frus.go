package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	h "me_exporter/app/helpers"
)

type fru struct {
	meSession *MeMetrics
	status    descMétrica
	logger    log.Logger
}

func init() {
	registerCollector("fru", NewFruCollector)
}

func NewFruCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &fru{
		meSession: me,
		status: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("fru", "status"),
				"Status", []string{"name", "status", "location", "enclosure"}),
		},
		logger: logger,
	}, nil
}
func (f fru) Update(ch chan<- prometheus.Metric) error {
	if err := f.meSession.Frus(); err != nil {
		return err
	}

	is := h.IntToString

	for _, fru := range f.meSession.frus {
		ch <- prometheus.MustNewConstMetric(f.status.desc, f.status.tipo,
			float64(fru.FruStatusNumeric), fru.Name, fru.FruStatus, fru.FruLocation, is(fru.EnclosureID))
	}

	return nil
}
