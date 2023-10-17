package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"me_exporter/Me"
)

type poolStatistics struct {
	//Where have perMinute or PerHour is prometheus.GaugeValue
	meSession                        *MeMetrics
	sampleTime                       descMétrica
	pagesAllocPerMinute              descMétrica
	pagesAllocPerHour                descMétrica
	pagesDeallocPerMinute            descMétrica
	pagesDeallocPerHour              descMétrica
	pagesUnmapPerMinute              descMétrica
	pagesUnmapPerHour                descMétrica
	numBlockedSsdPromotionsPerMinute descMétrica
	numBlockedSsdPromotionsPerHour   descMétrica
	numPageAllocations               descMétrica
	numPageDeallocations             descMétrica
	numPageUnmaps                    descMétrica
	numPagePromotionsToSsdBlocked    descMétrica
	numHotPageMoves                  descMétrica
	numColdPageMoves                 descMétrica

	// TierStatistics
	Tier                      descMétrica
	tierPagesAllocPerMinute   descMétrica
	tierPagesDeallocPerMinute descMétrica
	pagesReclaimed            descMétrica
	numPagesUnmapPerMinute    descMétrica

	logger log.Logger
}

func init() {
	registerCollector("pool_statistics", NewPoolStatisticsCollector)
}

func NewPoolStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &poolStatistics{
		meSession: me,
		sampleTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "sample_time"),
				"Sample time", []string{"pool", "sample_time"}),
		},
		pagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_alloc_per_minute"),
				"Pages alloc per minute", []string{"pool"}),
		},
		pagesAllocPerHour: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_alloc_per_hour"),
				"Pages alloc per hour", []string{"pool"}),
		},
		pagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_dealloc_per_minute"),
				"Pages dealloc per minute", []string{"pool"}),
		},
		pagesDeallocPerHour: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_dealloc_per_hour"),
				"Pages dealloc per hour", []string{"pool"}),
		},
		pagesUnmapPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_unmap_per_minute"),
				"Pages unmap per minute", []string{"pool"}),
		},
		pagesUnmapPerHour: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_unmap_per_hour"),
				"Pages unmap per hour", []string{"pool"}),
		},
		numBlockedSsdPromotionsPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_blocked_ssd_promotions_per_minute"),
				"Num blocked ssd promotions per minute", []string{"pool"}),
		},
		numBlockedSsdPromotionsPerHour: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_blocked_ssd_promotions_per_hour"),
				"Num blocked ssd promotions per hour", []string{"pool"}),
		},
		numPageAllocations: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_page_allocations"),
				"Num page allocations", []string{"pool"}),
		},
		numPageDeallocations: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_page_deallocations"),
				"Num page deallocations", []string{"pool"}),
		},
		numPageUnmaps: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_page_unmaps"),
				"Num page unmaps", []string{"pool"}),
		},
		numPagePromotionsToSsdBlocked: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_page_promotions_to_ssd_blocked"),
				"Num page promotions to ssd blocked", []string{"pool"}),
		},
		numHotPageMoves: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_hot_page_moves"),
				"Num hot page moves", []string{"pool"}),
		},
		numColdPageMoves: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_cold_page_moves"),
				"Num cold page moves", []string{"pool"}),
		},
		//TierStatistics

		Tier: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "tier"),
				"Tier", []string{"pool", "tier"}),
		},
		tierPagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "tier_pages_alloc_per_minute"),
				"Tier pages alloc per minute", []string{"pool", "tier"}),
		},
		tierPagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "tier_pages_dealloc_per_minute"),
				"Tier pages dealloc per minute", []string{"pool", "tier"}),
		},
		pagesReclaimed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "pages_reclaimed"),
				"Pages reclaimed", []string{"pool", "tier"}),
		},
		numPagesUnmapPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("pool_statistics", "num_pages_unmap_per_minute"),
				"Num pages unmap per minute", []string{"pool", "tier"}),
		},
		logger: logger,
	}, nil
}

func (p poolStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := p.meSession.PoolsStatistics(); err != nil {
		return err
	}

	for _, pool := range p.meSession.poolStatistics {
		ch <- p.sampleTime.constMetric(float64(pool.SampleTimeNumeric), pool.Pool, pool.SampleTime)
		ch <- p.pagesAllocPerMinute.constMetric(float64(pool.PagesAllocPerMinute), pool.Pool)
		ch <- p.pagesAllocPerHour.constMetric(float64(pool.PagesAllocPerHour), pool.Pool)
		ch <- p.pagesDeallocPerMinute.constMetric(float64(pool.PagesDeallocPerMinute), pool.Pool)
		ch <- p.pagesDeallocPerHour.constMetric(float64(pool.PagesDeallocPerHour), pool.Pool)
		ch <- p.pagesUnmapPerMinute.constMetric(float64(pool.PagesUnmapPerMinute), pool.Pool)
		ch <- p.pagesUnmapPerHour.constMetric(float64(pool.PagesUnmapPerHour), pool.Pool)
		ch <- p.numBlockedSsdPromotionsPerMinute.constMetric(float64(pool.NumBlockedSsdPromotionsPerMinute), pool.Pool)
		ch <- p.numBlockedSsdPromotionsPerHour.constMetric(float64(pool.NumBlockedSsdPromotionsPerHour), pool.Pool)
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
		ch <- p.tierPagesAllocPerMinute.constMetric(float64(tier.PagesAllocPerMinute), pool.Pool, tier.Tier)
		ch <- p.tierPagesDeallocPerMinute.constMetric(float64(tier.PagesDeallocPerMinute), pool.Pool, tier.Tier)
		ch <- p.pagesReclaimed.constMetric(float64(tier.PagesReclaimed), pool.Pool, tier.Tier)
		ch <- p.numPagesUnmapPerMinute.constMetric(float64(tier.NumPagesUnmapPerMinute), pool.Pool, tier.Tier)
	}
}
