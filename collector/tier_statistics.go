package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type tierStatistics struct {
	meSession              *MeMetrics
	up                     descMétrica
	TierNumeric            descMétrica
	PagesAllocPerMinute    descMétrica
	PagesDeallocPerMinute  descMétrica
	PagesReclaimed         descMétrica
	NumPagesUnmapPerMinute descMétrica
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
		TierNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "tier_numeric"),
				"Tier numeric", []string{"pool", "num_tier"}),
		},
		PagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "pages_alloc_per_minute"),
				"Pages alloc per minute", []string{"pool"}),
		},
		PagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "pages_dealloc_per_minute"),
				"Pages dealloc per minute", []string{"pool"}),
		},
		PagesReclaimed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier_statistics", "pages_reclaimed"),
				"Pages reclaimed", []string{"pool"}),
		},
		NumPagesUnmapPerMinute: descMétrica{prometheus.GaugeValue,
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

	for _, tier := range t.meSession.tierStatistics {
		ch <- t.up.constMetric(1, tier.Pool, tier.Tier, tier.SerialNumber)
		ch <- t.TierNumeric.constMetric(float64(tier.TierNumeric), tier.Pool, tier.Tier)
		ch <- t.PagesAllocPerMinute.constMetric(float64(tier.PagesAllocPerMinute), tier.Pool)
		ch <- t.PagesDeallocPerMinute.constMetric(float64(tier.PagesDeallocPerMinute), tier.Pool)
		ch <- t.PagesReclaimed.constMetric(float64(tier.PagesReclaimed), tier.Pool)
		ch <- t.NumPagesUnmapPerMinute.constMetric(float64(tier.NumPagesUnmapPerMinute), tier.Pool)
	}

	return nil
}
