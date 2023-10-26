package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type tier struct {
	// All labels have pool
	meSession      *MeMetrics
	up             descMétrica // desMétrica is 1, and your labels is pool and serial_number
	tier           descMétrica
	poolPercentage descMétrica
	diskcount      descMétrica
	rawSize        descMétrica
	totalSize      descMétrica
	allocatedSize  descMétrica
	availableSize  descMétrica
	affinitySize   descMétrica
	logger         log.Logger
}

func (t tier) Update(ch chan<- prometheus.Metric) error {
	//TODO implement me
	panic("implement me")
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
				"Was the last query of tier successful.", []string{"pool"}),
		},
		tier: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "tier_numeric"),
				"Tier numeric", []string{"pool", "num_tier"}),
		},
		poolPercentage: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "pool_percentage"),
				"Pool percentage", []string{"pool"}),
		},
		diskcount: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "diskcount"),
				"Diskcount", []string{"pool"}),
		},
		rawSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "raw_size_numeric"),
				"Raw size in blocks", []string{"pool"}),
		},
		totalSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "total_size_numeric"),
				"Total size in blocks", []string{"pool"}),
		},
		allocatedSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "allocated_size_numeric"),
				"Allocated size in blocks", []string{"pool"}),
		},
		availableSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "available_size_numeric"),
				"Available size in blocks", []string{"pool"}),
		},
		affinitySize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("tier", "affinity_size_numeric"),
				"Affinity size in Bytes", []string{"pool"}),
		},
		logger: logger,
	}, nil
}
