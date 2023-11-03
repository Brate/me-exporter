package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type tier struct {
	meSession *MeMetrics

	up             descMétrica
	tier           descMétrica
	poolPercentage descMétrica
	diskCount      descMétrica
	rawSize        descMétrica
	totalSize      descMétrica
	allocatedSize  descMétrica
	availableSize  descMétrica
	affinitySize   descMétrica
	logger         log.Logger
}

func init() {
	registerCollector("tier", NewTierCollector)
}

func NewTierCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &tier{
		meSession: me,
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "up"),
				"Was the last query of tier successful.", []string{"pool", "tier", "serial_number"}),
		},
		tier: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "tier"),
				"Tier numeric", []string{"pool", "tier"}),
		},
		poolPercentage: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "pool_percentage"),
				"Pool percentage", []string{"pool", "tier"}),
		},
		diskCount: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "diskCount"),
				"Diskcount", []string{"pool", "tier"}),
		},
		rawSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "raw_size_blocks"),
				"Raw size in blocks", []string{"pool", "tier"}),
		},
		totalSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "total_size"),
				"Total size in blocks", []string{"pool", "tier"}),
		},
		allocatedSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "allocated_size"),
				"Allocated size in blocks", []string{"pool", "tier"}),
		},
		availableSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "available_size"),
				"Available size in blocks", []string{"pool", "tier"}),
		},
		affinitySize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "affinity_size"),
				"Affinity size in Bytes", []string{"pool", "tier"}),
		},
		logger: logger,
	}, nil
}

func (t tier) Update(ch chan<- prometheus.Metric) error {
	if err := t.meSession.Tiers(); err != nil {
		return err
	}

	for _, tier := range t.meSession.tiers {
		ch <- t.up.constMetric(1, tier.Pool, tier.Tier, tier.SerialNumber)
		ch <- t.tier.constMetric(float64(tier.TierNumeric), tier.Pool, tier.Tier)
		ch <- t.poolPercentage.constMetric(float64(tier.PoolPercentage), tier.Pool, tier.Tier)
		ch <- t.diskCount.constMetric(float64(tier.Diskcount), tier.Pool, tier.Tier)
		ch <- t.rawSize.constMetric(float64(tier.RawSizeNumeric), tier.Pool, tier.Tier)
		ch <- t.totalSize.constMetric(float64(tier.TotalSizeNumeric), tier.Pool, tier.Tier)
		ch <- t.allocatedSize.constMetric(float64(tier.AllocatedSizeNumeric), tier.Pool, tier.Tier)
		ch <- t.availableSize.constMetric(float64(tier.AvailableSizeNumeric), tier.Pool, tier.Tier)
		ch <- t.affinitySize.constMetric(float64(tier.AffinitySizeNumeric), tier.Pool, tier.Tier)
	}

	return nil
}
