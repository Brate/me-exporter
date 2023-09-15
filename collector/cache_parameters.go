package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type cacheSettingsController struct {
	meSession      *MeMetrics
	cacheBlockSize descMétrica
	operationMode  descMétrica

	// Controller Cache Parameters
	CompactFlash       descMétrica
	CompactFlashHealth descMétrica
	CacheFlush         descMétrica
	WriteBack          descMétrica
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
		operationMode: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_redundancy", "operation_mode"),
				"Operation mode on the controller", []string{"mode"}),
		},
		// Controller Cache Parameters
		CompactFlash: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "compact_flash_status"),
				"Compact flash status ", []string{"controller", "status"}),
		},
		CompactFlashHealth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "compact_flash_health"),
				"Compact flash health ", []string{"controller", "status"}),
		},
		CacheFlush: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "cache_flush"),
				"Cache flush ", []string{"controller"}),
		},
		WriteBack: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("cache_settings_controller", "write_back_status"),
				"Write back status ", []string{"controller", "status"}),
		},
		logger: logger,
	}, nil
}

func (c *cacheSettingsController) Update(ch chan<- prometheus.Metric) error {
	if err := c.meSession.CacheSettings(); err != nil {
		return err
	}

	s := c.meSession.cacheSettings

	ch <- prometheus.MustNewConstMetric(c.operationMode.desc, c.operationMode.tipo,
		float64(s.OperationModeNumeric), s.OperationMode)
	ch <- prometheus.MustNewConstMetric(c.cacheBlockSize.desc, c.cacheBlockSize.tipo,
		float64(s.CacheBlockSize))

	for _, controller := range s.ControllerCacheParameters {
		ch <- prometheus.MustNewConstMetric(c.CompactFlash.desc, c.CompactFlash.tipo,
			float64(controller.CompactFlashStatusNumeric), controller.ControllerID, controller.CompactFlashStatus)
		ch <- prometheus.MustNewConstMetric(c.CompactFlashHealth.desc, c.CompactFlashHealth.tipo,
			float64(controller.CompactFlashHealthNumeric), controller.ControllerID, controller.CompactFlashHealth)
		ch <- prometheus.MustNewConstMetric(c.CacheFlush.desc, c.CacheFlush.tipo,
			float64(controller.CacheFlushNumeric), controller.ControllerID)
		ch <- prometheus.MustNewConstMetric(c.WriteBack.desc, c.WriteBack.tipo,
			float64(controller.WriteBackStatusNumeric), controller.ControllerID, controller.WriteBackStatus)
	}

	return nil
}
