package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type volumeStatistics struct {
	//All labels volume-name
	meSession *MeMetrics
	//up        descMétrica
	//bytesPerSecond         descMétrica
	numberOfReads       descMétrica
	numberOfWrites      descMétrica
	dataRead            descMétrica
	dataWritten         descMétrica
	allocatedPages      descMétrica
	percentTierSsd      descMétrica
	percentTierSas      descMétrica
	percentTierSata     descMétrica
	percentAllocatedRfc descMétrica
	//pagesAllocPerMinute    descMétrica
	//pagesDeallocPerMinute  descMétrica
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
	//resetTime              descMétrica
	//startSampleTime        descMétrica
	//stopSampleTime         descMétrica
	logger log.Logger
}

func init() {
	registerCollector("volume_statistics", NewVolumeStatisticsCollector)
}

func NewVolumeStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &volumeStatistics{
		meSession: me,
		// TODO: Mover para NewVolume
		//up: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("volume_statistics", "up"),
		//		"Up", []string{"volume_name", "serial_number"}),
		//},
		numberOfReads: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "read_count"),
				"Number of reads", []string{"volume_name"}),
		},
		numberOfWrites: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "write_count"),
				"Number of writes", []string{"volume_name"}),
		},
		dataRead: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "data_read_bytes"),
				"Data read", []string{"volume_name"}),
		},
		dataWritten: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "data_written_bytes"),
				"Data written", []string{"volume_name"}),
		},
		allocatedPages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "allocated_pages"),
				"number of pages allocated to the volume", []string{"volume_name"}),
		},
		percentTierSsd: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "percent_tier_ssd"),
				"percentage of volume capacity occupied by data in the Performance tier", []string{"volume_name"}),
		},
		percentTierSas: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "percent_tier_sas"),
				"percentage of volume capacity occupied by data in the Standard tier", []string{"volume_name"}),
		},
		percentTierSata: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "percent_tier_sata"),
				"percentage of volume capacity occupied by data in the Archive tier", []string{"volume_name"}),
		},
		percentAllocatedRfc: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "percent_rfc"),
				"percentage of volume capacity occupied by data in read cache", []string{"volume_name"}),
		},
		sharedPages: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "shared_pages"),
				"number of pages that are shared between this volume and any other volumes", []string{"volume_name"}),
		},
		writeCacheHits: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "write_cache_hits"),
				"number of times the block written to is found in cache", []string{"volume_name"}),
		},
		writeCacheMisses: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "write_cache_misses"),
				"number of times the block written to is not found in cache", []string{"volume_name"}),
		},
		readCacheHits: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "read_cache_hits"),
				"number of times the block to be read is found in cache", []string{"volume_name"}),
		},
		readCacheMisses: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "read_cache_misses"),
				"number of times the block to be read is not found in cache", []string{"volume_name"}),
		},
		smallDestages: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "destage_small_count"),
				"number of times flush from cache to disk is not a full stripe", []string{"volume_name"}),
		},
		fullStripeWriteDestage: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "destage_full_stripe_count"),
				"number of times flush from cache to disk is a full stripe", []string{"volume_name"}),
		},
		readAheadOperations: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("volume", "read_ahead_operations"),
				"number of read pre-fetch or anticipatory-read operations", []string{"volume_name"}),
		},
		writeCacheSpace: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "write_cache_space"),
				"cache size used on behalf of the volume", []string{"volume_name"}),
		},
		writeCachePercent: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("volume", "write_cache_percent"),
				"percentage of cache used on behalf of the volume", []string{"volume_name"}),
		},
		logger: logger,
	}, nil
}

func (v volumeStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := v.meSession.VolumeStatistics(); err != nil {
		return err
	}

	for _, volume := range v.meSession.volumeStatistics {
		ch <- v.numberOfReads.constMetric(float64(volume.NumberOfReads), volume.VolumeName)
		ch <- v.numberOfWrites.constMetric(float64(volume.NumberOfWrites), volume.VolumeName)
		ch <- v.dataRead.constMetric(float64(volume.DataReadNumeric), volume.VolumeName)
		ch <- v.dataWritten.constMetric(float64(volume.DataWrittenNumeric), volume.VolumeName)
		ch <- v.allocatedPages.constMetric(float64(volume.AllocatedPages), volume.VolumeName)
		ch <- v.percentTierSsd.constMetric(float64(volume.PercentTierSsd), volume.VolumeName)
		ch <- v.percentTierSas.constMetric(float64(volume.PercentTierSas), volume.VolumeName)
		ch <- v.percentTierSata.constMetric(float64(volume.PercentTierSata), volume.VolumeName)
		ch <- v.percentAllocatedRfc.constMetric(float64(volume.PercentAllocatedRfc), volume.VolumeName)
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
	}

	return nil
}
