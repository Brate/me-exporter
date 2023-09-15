package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type cacheSettingsController struct {
	meSession                 *MeMetrics
	cacheBlockSize            descMétrica
	controllerCacheParameters descMétrica
	operationMode             descMétrica

	// Controller Cache Parameters
	CompactFlashStatus descMétrica
	CompactFlashHealth descMétrica
	CacheFlush         descMétrica
	WriteBackStatus    descMétrica
	logger             log.Logger
}

func init() {
	registerCollector("cache_settings_controller", NewCacheSettingsControllerCollector)
}

func NewCacheSettingsControllerCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &cacheSettingsController{
		meSession: me,
		cacheBlockSize: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings", "block_size_kbytes"),
				"Cache block size", nil),
		},
		controllerCacheParameters: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "controller_cache_parameters"),
				"Controller cache parameters", []string{"controller"}),
		},
		operationMode: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_redundancy", "operation_mode"),
				"Operation mode on the controller", []string{"mode"}),
		},
		// Controller Cache Parameters
		CompactFlashStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "compact_flash_status"),
				"Compact flash status ", []string{"controller"}),
		},
		CompactFlashHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "compact_flash_health"),
				"Compact flash health ", []string{"controller"}),
		},
		CacheFlush: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "cache_flush"),
				"Cache flush ", []string{"controller"}),
		},
		WriteBackStatus: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "write_back_status"),
				"Write back status ", []string{"controller"}),
		},
		logger: logger,
	}, nil
}

func (c *cacheSettingsController) Update(ch chan<- prometheus.Metric) error {
	if err := c.meSession.CacheSettings(); err != nil {
		return err
	}

	s := c.meSession.cacheSettings

	ch <- prometheus.MustNewConstMetric(c.operationMode.desc, c.operationMode.tipo, float64(s.OperationModeNumeric), s.OperationMode)
	ch <- prometheus.MustNewConstMetric(c.cacheBlockSize.desc, c.cacheBlockSize.tipo, float64(s.CacheBlockSize))

	return nil
}
