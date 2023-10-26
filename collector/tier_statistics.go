package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type tierStatistics struct {
	meSession              *MeMetrics
	up                     descMétrica
	tierNumeric            descMétrica
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
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "up"),
				"Was the last query of tier_statistics successful.", []string{"pool", "serial_number"}),
		},
		tierNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "tier_numeric"),
				"Tier numeric", []string{"pool", "num_tier"}),
		},
		pagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "pages_alloc_per_minute"),
				"Pages alloc per minute", []string{"pool"}),
		},
		pagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "pages_dealloc_per_minute"),
				"Pages dealloc per minute", []string{"pool"}),
		},
		pagesReclaimed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "pages_reclaimed"),
				"Pages reclaimed", []string{"pool"}),
		},
		numPagesUnmapPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "num_pages_unmap_per_minute"),
				"Num pages unmap per minute", []string{"pool"}),
		},
		logger: logger,
	}, nil
}

func (t tierStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := t.meSession.TierStatistics(); err != nil {
		return err
	}

	for _, tierS := range t.meSession.tierStatistics {
		ch <- t.up.constMetric(1, tierS.Pool, tierS.Tier, tierS.SerialNumber)
		ch <- t.tierNumeric.constMetric(float64(tierS.TierNumeric), tierS.Pool, tierS.Tier)
		ch <- t.pagesAllocPerMinute.constMetric(float64(tierS.PagesAllocPerMinute), tierS.Pool)
		ch <- t.pagesDeallocPerMinute.constMetric(float64(tierS.PagesDeallocPerMinute), tierS.Pool)
		ch <- t.pagesReclaimed.constMetric(float64(tierS.PagesReclaimed), tierS.Pool)
		ch <- t.numPagesUnmapPerMinute.constMetric(float64(tierS.NumPagesUnmapPerMinute), tierS.Pool)
	}

	return nil
}
