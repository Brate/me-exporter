package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type controllerStatisticsCollector struct {
	meSession              *MeMetrics
	cpuLoad                descMétrica
	powerOnTime            descMétrica
	writeCacheUsed         descMétrica
	bytesPerSecondNumeric  descMétrica
	iops                   descMétrica
	numberOfReads          descMétrica
	readCacheHits          descMétrica
	readCacheMisses        descMétrica
	numberOfWrites         descMétrica
	writeCacheHits         descMétrica
	writeCacheMisses       descMétrica
	dataReadNumeric        descMétrica
	dataWrittenNumeric     descMétrica
	numForwardedCmds       descMétrica
	resetTimeNumeric       descMétrica
	startSampleTimeNumeric descMétrica
	stopSampleTimeNumeric  descMétrica
	totalPowerOnHours      descMétrica
	logger                 log.Logger
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
		powerOnTime: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "power_on_time"),
				"Power on time on the controller", []string{"controller"}),
		},
		writeCacheUsed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "write_cache_used"),
				"Write cache used on the controller", []string{"controller"}),
		},
		bytesPerSecondNumeric: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("controller", "bytes_per_second_numeric"),
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
		dataReadNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "data_read_numeric"),
				"Data read numeric on the controller", []string{"controller"}),
		},
		dataWrittenNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "data_written_numeric"),
				"Data written numeric on the controller", []string{"controller"}),
		},
		numForwardedCmds: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "num_forwarded_cmds"),
				"Number of forwarded commands on the controller", []string{"controller"}),
		},
		resetTimeNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "reset_time_numeric"),
				"Reset time numeric on the controller", []string{"controller"}),
		},
		startSampleTimeNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "start_sample_time_numeric"),
				"Start sample time numeric on the controller", []string{"controller"}),
		},
		stopSampleTimeNumeric: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("controller", "stop_sample_time_numeric"),
				"Stop sample time numeric on the controller", []string{"controller"}),
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
	ch <- prometheus.MustNewConstMetric(c.bytesPerSecondNumeric.desc, c.bytesPerSecondNumeric.tipo, float64(stats.BytesPerSecondNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.iops.desc, c.iops.tipo, float64(stats.Iops), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.numberOfReads.desc, c.numberOfReads.tipo, float64(stats.NumberOfReads), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.readCacheHits.desc, c.readCacheHits.tipo, float64(stats.ReadCacheHits), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.readCacheMisses.desc, c.readCacheMisses.tipo, float64(stats.ReadCacheMisses), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.numberOfWrites.desc, c.numberOfWrites.tipo, float64(stats.NumberOfWrites), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.writeCacheHits.desc, c.writeCacheHits.tipo, float64(stats.WriteCacheHits), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.writeCacheMisses.desc, c.writeCacheMisses.tipo, float64(stats.WriteCacheMisses), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.dataReadNumeric.desc, c.dataReadNumeric.tipo, float64(stats.DataReadNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.dataWrittenNumeric.desc, c.dataWrittenNumeric.tipo, float64(stats.DataWrittenNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.numForwardedCmds.desc, c.numForwardedCmds.tipo, float64(stats.NumForwardedCmds), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.resetTimeNumeric.desc, c.resetTimeNumeric.tipo, float64(stats.ResetTimeNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.startSampleTimeNumeric.desc, c.startSampleTimeNumeric.tipo, float64(stats.StartSampleTimeNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.stopSampleTimeNumeric.desc, c.stopSampleTimeNumeric.tipo, float64(stats.StopSampleTimeNumeric), stats.DurableID)
	ch <- prometheus.MustNewConstMetric(c.totalPowerOnHours.desc, c.totalPowerOnHours.tipo, stats.TotalPowerOnHoursNumeric, stats.DurableID)
	return nil
}
