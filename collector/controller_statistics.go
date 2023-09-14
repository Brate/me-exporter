package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type controllerStatisticsCollector struct {
	meSession         *MeMetrics
	cpuLoad           descMétrica
	powerOnTime       descMétrica
	writeCacheUsed    descMétrica
	bytesPerSecond    descMétrica
	iops              descMétrica
	numberOfReads     descMétrica
	readCacheHits     descMétrica
	readCacheMisses   descMétrica
	numberOfWrites    descMétrica
	writeCacheHits    descMétrica
	writeCacheMisses  descMétrica
	dataRead          descMétrica
	dataWritten       descMétrica
	numForwardedCmds  descMétrica
	statsResetTime    descMétrica
	totalPowerOnHours descMétrica
	logger            log.Logger
}

func init() {
	registerCollector("controller", NewControllerStatisticsCollector)
}

func NewControllerStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &controllerStatisticsCollector{
		meSession: me,
		cpuLoad: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "cpu_load"),
				"CPU load on the controller", []string{"controller"}),
		},
		powerOnTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "power_on_time_seconds"),
				"Power on time on the controller", []string{"controller"}),
		},
		writeCacheUsed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "write_cache_used"),
				"Write cache used on the controller", []string{"controller"}),
		},
		bytesPerSecond: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "bytes_per_second"),
				"Bytes per second on the controller", []string{"controller"}),
		},
		iops: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "iops"),
				"IOPS on the controller", []string{"controller"}),
		},
		numberOfReads: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "number_of_reads"),
				"Number of reads on the controller", []string{"controller"}),
		},
		readCacheHits: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "read_cache_hits"),
				"Read cache hits on the controller", []string{"controller"}),
		},
		readCacheMisses: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "read_cache_misses"),
				"Read cache misses on the controller", []string{"controller"}),
		},
		numberOfWrites: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "number_of_writes"),
				"Number of writes on the controller", []string{"controller"}),
		},
		writeCacheHits: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "write_cache_hits"),
				"Write cache hits on the controller", []string{"controller"}),
		},
		writeCacheMisses: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "write_cache_misses"),
				"Write cache misses on the controller", []string{"controller"}),
		},
		dataRead: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "data_read"),
				"Data read numeric on the controller", []string{"controller"}),
		},
		dataWritten: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "data_written"),
				"Data written numeric on the controller", []string{"controller"}),
		},
		numForwardedCmds: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "num_forwarded_cmds"),
				"Number of forwarded commands on the controller", []string{"controller"}),
		},
		statsResetTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "time_since_stats_reset_seconds"),
				"Time elapsed since statistics reset on the controller", []string{"controller"}),
		},
		totalPowerOnHours: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "total_power_on_hours"),
				"Total power on hours on the controller", []string{"controller"}),
		},
		logger: logger,
	}, nil
}

func (c *controllerStatisticsCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.meSession.ControllerStatistics(); err != nil {
		return err
	}

	stats := c.meSession.controllerStatistics
	ch <- prometheus.MustNewConstMetric(c.cpuLoad.desc, c.cpuLoad.tipo, float64(stats.CPULoad), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.powerOnTime.desc, c.powerOnTime.tipo, float64(stats.PowerOnTime), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.writeCacheUsed.desc, c.writeCacheUsed.tipo, float64(stats.WriteCacheUsed), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.bytesPerSecond.desc, c.bytesPerSecond.tipo, float64(stats.BytesPerSecondNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.iops.desc, c.iops.tipo, float64(stats.Iops), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.numberOfReads.desc, c.numberOfReads.tipo, float64(stats.NumberOfReads), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.readCacheHits.desc, c.readCacheHits.tipo, float64(stats.ReadCacheHits), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.readCacheMisses.desc, c.readCacheMisses.tipo, float64(stats.ReadCacheMisses), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.numberOfWrites.desc, c.numberOfWrites.tipo, float64(stats.NumberOfWrites), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.writeCacheHits.desc, c.writeCacheHits.tipo, float64(stats.WriteCacheHits), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.writeCacheMisses.desc, c.writeCacheMisses.tipo, float64(stats.WriteCacheMisses), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.dataRead.desc, c.dataRead.tipo, float64(stats.DataReadNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.dataWritten.desc, c.dataWritten.tipo, float64(stats.DataWrittenNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.numForwardedCmds.desc, c.numForwardedCmds.tipo, float64(stats.NumForwardedCmds), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.statsResetTime.desc, c.statsResetTime.tipo, float64(stats.TimeSinceStatsReset), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.totalPowerOnHours.desc, c.totalPowerOnHours.tipo, stats.TotalPowerOnHoursNumeric, stats.DurableID)
	return nil
}
