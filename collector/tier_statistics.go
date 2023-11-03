package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type tierStatistics struct {
	meSession *MeMetrics

	pagesAllocPerMinute    descMétrica
	pagesDeallocPerMinute  descMétrica
	pagesReclaimed         descMétrica
	numPagesUnmapPerMinute descMétrica
	logger                 log.Logger
}

func init() {
	registerCollector("tier_statistics", NewTierStatisticsCollector)
}

func NewTierStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &tierStatistics{
		meSession: me,
		pagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "pages_allocated_per_minute"),
				"Pages allocated per minute", []string{"pool", "tier"}),
		},
		pagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "pages_deallocated_per_minute"),
				"Pages deallocated per minute", []string{"pool", "tier"}),
		},
		pagesReclaimed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "pages_reclaimed"),
				"Pages reclaimed", []string{"pool", "tier"}),
		},
		numPagesUnmapPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "num_pages_unmap_per_minute"),
				"Num pages unmap per minute", []string{"pool", "tier"}),
		},
		logger: logger,
	}, nil
}

func (t tierStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := t.meSession.TierStatistics(); err != nil {
		return err
	}

	for _, tierS := range t.meSession.tierStatistics {
		ch <- t.pagesAllocPerMinute.constMetric(float64(tierS.PagesAllocPerMinute), tierS.Pool, tierS.Tier)
		ch <- t.pagesDeallocPerMinute.constMetric(float64(tierS.PagesDeallocPerMinute), tierS.Pool, tierS.Tier)
		ch <- t.pagesReclaimed.constMetric(float64(tierS.PagesReclaimed), tierS.Pool, tierS.Tier)
		ch <- t.numPagesUnmapPerMinute.constMetric(float64(tierS.NumPagesUnmapPerMinute), tierS.Pool, tierS.Tier)
	}

	return nil
}
