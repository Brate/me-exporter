package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type unwritableCache struct {
	meSession             *MeMetrics
	unwritableAPercentage descMétrica
	unwritableBPercentage descMétrica
	logger                log.Logger
}

func init() {
	registerCollector("unwritable_cache", NewUnwritableCacheCollector)
}

func NewUnwritableCacheCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &unwritableCache{
		meSession: me,
		unwritableAPercentage: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("unwritable_cache", "unwritable_a_percentage"),
				"Unwritable A Percentage", []string{}),
		},
		unwritableBPercentage: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("unwritable_cache", "unwritable_b_percentage"),
				"Unwritable B Percentage", []string{}),
		},
		logger: logger,
	}, nil
}

func (u unwritableCache) Update(ch chan<- prometheus.Metric) error {
	if err := u.meSession.UnwritableCache(); err != nil {
		return err
	}

	for _, unwrtCache := range u.meSession.unwritableCache {
		ch <- u.unwritableAPercentage.constMetric(float64(unwrtCache.UnwritableAPercentage))
		ch <- u.unwritableBPercentage.constMetric(float64(unwrtCache.UnwritableBPercentage))
	}

	return nil
}
