package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type volumeStatistics struct {
	//All labels volume-name
	meSession              *MeMetrics
	up                     descMétrica
	bytesPerSecond         descMétrica
	numberOfReads          descMétrica
	numberOfWrites         descMétrica
	dataRead               descMétrica
	dataWritten            descMétrica
	allocatedPages         descMétrica
	percentTierSsd         descMétrica
	percentTierSas         descMétrica
	percentTierSata        descMétrica
	percentAllocatedRfc    descMétrica
	pagesAllocPerMinute    descMétrica
	pagesDeallocPerMinute  descMétrica
	sharedPages            descMétrica
	writeCacheHits         descMétrica
	writeCacheMisses       descMétrica
	readCacheHits          descMétrica
	readCacheMisses        descMétrica
	smallDestages          descMétrica
	fullStripeWriteDestage descMétrica
	readAheadOperations    descMétrica
	writeCacheSpace        descMétrica
	writeCachePercent      descMétrica
	resetTime              descMétrica
	startSampleTime        descMétrica
	stopSampleTime         descMétrica
	logger                 log.Logger
}

func init() {
	registerCollector("volume_statistics", NewVolumeStatisticsCollector)
}

func NewVolumeStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &volumeStatistics{
		meSession: me,
		up: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "up"),
				"Up", []string{"volume_name", "serial_number"}),
		},
		bytesPerSecond: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "bytes_per_second"),
				"Bytes per second", []string{"volume_name", "bytes_per_second"}),
		},
		numberOfReads: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "number_of_reads"),
				"Number of reads", []string{"volume_name"}),
		},
		numberOfWrites: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "number_of_writes"),
				"Number of writes", []string{"volume_name"}),
		},
		dataRead: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "data_read"),
				"Data read", []string{"volume_name"}),
		},
		dataWritten: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "data_written"),
				"Data written", []string{"volume_name"}),
		},
		allocatedPages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "allocated_pages"),
				"Allocated pages", []string{"volume_name"}),
		},
		percentTierSsd: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "percent_tier_ssd"),
				"Percent tier ssd", []string{"volume_name"}),
		},
		percentTierSas: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "percent_tier_sas"),
				"Percent tier sas", []string{"volume_name"}),
		},
		percentTierSata: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "percent_tier_sata"),
				"Percent tier sata", []string{"volume_name"}),
		},
		percentAllocatedRfc: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "percent_allocated_rfc"),
				"Percent allocated rfc", []string{"volume_name"}),
		},
		pagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "pages_alloc_per_minute"),
				"Pages alloc per minute", []string{"volume_name"}),
		},
		pagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "pages_dealloc_per_minute"),
				"Pages dealloc per minute", []string{"volume_name"}),
		},
		sharedPages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "shared_pages"),
				"Shared pages", []string{"volume_name"}),
		},
		writeCacheHits: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "write_cache_hits"),
				"Write cache hits", []string{"volume_name"}),
		},
		writeCacheMisses: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "write_cache_misses"),
				"Write cache misses", []string{"volume_name"}),
		},
		readCacheHits: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "read_cache_hits"),
				"Read cache hits", []string{"volume_name"}),
		},
		readCacheMisses: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "read_cache_misses"),
				"Read cache misses", []string{"volume_name"}),
		},
		smallDestages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "small_destages"),
				"Small destages", []string{"volume_name"}),
		},
		fullStripeWriteDestage: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "full_stripe_write_destage"),
				"Full stripe write destage", []string{"volume_name"}),
		},
		readAheadOperations: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "read_ahead_operations"),
				"Read ahead operations", []string{"volume_name"}),
		},
		writeCacheSpace: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "write_cache_space"),
				"Write cache space", []string{"volume_name"}),
		},
		writeCachePercent: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "write_cache_percent"),
				"Write cache percent", []string{"volume_name"}),
		},
		resetTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "reset_time"),
				"Reset time in epoch", []string{"volume_name", "reset_time"}),
		},
		startSampleTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "start_sample_time"),
				"Start sample time in epoch", []string{"volume_name", "start_sample_time"}),
		},
		stopSampleTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume_statistics", "stop_sample_time"),
				"Stop sample time in epoch", []string{"volume_name", "stop_sample_time"}),
		},
		logger: logger,
	}, nil
}

func (v volumeStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := v.meSession.VolumeStatistics(); err != nil {
		return err
	}

	for _, volume := range v.meSession.volumeStatistics {
		ch <- v.up.constMetric(1, volume.VolumeName, volume.SerialNumber)
		ch <- v.bytesPerSecond.constMetric(float64(volume.BytesPerSecondNumeric), volume.VolumeName, volume.BytesPerSecond)
		ch <- v.numberOfReads.constMetric(float64(volume.NumberOfReads), volume.VolumeName)
		ch <- v.numberOfWrites.constMetric(float64(volume.NumberOfWrites), volume.VolumeName)
		ch <- v.dataRead.constMetric(float64(volume.DataReadNumeric), volume.VolumeName)
		ch <- v.dataWritten.constMetric(float64(volume.DataWrittenNumeric), volume.VolumeName)
		ch <- v.allocatedPages.constMetric(float64(volume.AllocatedPages), volume.VolumeName)
		ch <- v.percentTierSsd.constMetric(float64(volume.PercentTierSsd), volume.VolumeName)
		ch <- v.percentTierSas.constMetric(float64(volume.PercentTierSas), volume.VolumeName)
		ch <- v.percentTierSata.constMetric(float64(volume.PercentTierSata), volume.VolumeName)
		ch <- v.percentAllocatedRfc.constMetric(float64(volume.PercentAllocatedRfc), volume.VolumeName)
		ch <- v.pagesAllocPerMinute.constMetric(float64(volume.PagesAllocPerMinute), volume.VolumeName)
		ch <- v.pagesDeallocPerMinute.constMetric(float64(volume.PagesDeallocPerMinute), volume.VolumeName)
		ch <- v.sharedPages.constMetric(float64(volume.SharedPages), volume.VolumeName)
		ch <- v.writeCacheHits.constMetric(float64(volume.WriteCacheHits), volume.VolumeName)
		ch <- v.writeCacheMisses.constMetric(float64(volume.WriteCacheMisses), volume.VolumeName)
		ch <- v.readCacheHits.constMetric(float64(volume.ReadCacheHits), volume.VolumeName)
		ch <- v.readCacheMisses.constMetric(float64(volume.ReadCacheMisses), volume.VolumeName)
		ch <- v.smallDestages.constMetric(float64(volume.SmallDestages), volume.VolumeName)
		ch <- v.fullStripeWriteDestage.constMetric(float64(volume.FullStripeWriteDestages), volume.VolumeName)
		ch <- v.readAheadOperations.constMetric(float64(volume.ReadAheadOperations), volume.VolumeName)
		ch <- v.writeCacheSpace.constMetric(float64(volume.WriteCacheSpace), volume.VolumeName)
		ch <- v.writeCachePercent.constMetric(float64(volume.WriteCachePercent), volume.VolumeName)
		ch <- v.resetTime.constMetric(float64(volume.ResetTimeNumeric), volume.VolumeName, volume.ResetTime)
		ch <- v.startSampleTime.constMetric(float64(volume.StartSampleTimeNumeric), volume.VolumeName, volume.StartSampleTime)
		ch <- v.stopSampleTime.constMetric(float64(volume.StopSampleTimeNumeric), volume.VolumeName, volume.StopSampleTime)
	}

	return nil
}
