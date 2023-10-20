package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"me_exporter/Me"
)

type poolStatistics struct {
	meSession                     *MeMetrics
	numPageAllocations            descMétrica
	numPageDeallocations          descMétrica
	numPageUnmaps                 descMétrica
	numPagePromotionsToSsdBlocked descMétrica
	numHotPageMoves               descMétrica
	numColdPageMoves              descMétrica

	// TierStatistics
	// TODO: Add Response time metrics on TierStatistics
	Tier           descMétrica
	pagesReclaimed descMétrica

	logger log.Logger
}

func init() {
	registerCollector("pool_statistics", NewPoolStatisticsCollector)
}

func NewPoolStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &poolStatistics{
		meSession: me,
		numPageAllocations: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "page_allocations"),
				"Num page allocations", []string{"pool"}),
		},
		numPageDeallocations: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "page_deallocations"),
				"Num page deallocations", []string{"pool"}),
		},
		numPageUnmaps: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "page_unmaps"),
				"Num page unmaps", []string{"pool"}),
		},
		numPagePromotionsToSsdBlocked: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "blocked_ssd_page_promotions"),
				"Num page promotions to ssd blocked", []string{"pool"}),
		},
		numHotPageMoves: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "hot_page_moves"),
				"Num hot page moves", []string{"pool"}),
		},
		numColdPageMoves: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "cold_page_moves"),
				"Num cold page moves", []string{"pool"}),
		},

		// TierStatistics
		Tier: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "tier"),
				"Tier", []string{"pool", "tier"}),
		},
		pagesReclaimed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_reclaimed"),
				"Pages reclaimed", []string{"pool", "tier"}),
		},
		logger: logger,
	}, nil
}

func (p poolStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := p.meSession.PoolsStatistics(); err != nil {
		return err
	}

	for _, pool := range p.meSession.poolStatistics {
		ch <- p.numPageAllocations.constMetric(float64(pool.NumPageAllocations), pool.Pool)
		ch <- p.numPageDeallocations.constMetric(float64(pool.NumPageDeallocations), pool.Pool)
		ch <- p.numPageUnmaps.constMetric(float64(pool.NumPageUnmaps), pool.Pool)
		ch <- p.numPagePromotionsToSsdBlocked.constMetric(float64(pool.NumPagePromotionsToSsdBlocked), pool.Pool)
		ch <- p.numHotPageMoves.constMetric(float64(pool.NumHotPageMoves), pool.Pool)
		ch <- p.numColdPageMoves.constMetric(float64(pool.NumColdPageMoves), pool.Pool)

		p.collectTierStatistics(ch, pool)

	}

	return nil
}

func (p poolStatistics) collectTierStatistics(ch chan<- prometheus.Metric, pool Me.PoolStatistics) {
	for _, tier := range pool.TierStatistics {
		ch <- p.Tier.constMetric(float64(tier.TierNumeric), pool.Pool, tier.Tier)
		ch <- p.pagesReclaimed.constMetric(float64(tier.PagesReclaimed), pool.Pool, tier.Tier)
	}
}
