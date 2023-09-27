package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type diskStatistics struct {
	meSession               *MeMetrics
	powerOnHours            descMétrica
	bytesPerSecond          descMétrica
	numberOfReads           descMétrica
	dataRead                descMétrica
	dataWritten             descMétrica
	lifetimeDataRead        descMétrica
	lifetimeDataWritten     descMétrica
	queueDepth              descMétrica
	resetTime               descMétrica
	startSampleTime         descMétrica
	stopSampleTime          descMétrica
	smartCount1             descMétrica
	smartCount2             descMétrica
	ioTimeoutCount1         descMétrica
	ioTimeoutCount2         descMétrica
	noResponseCount1        descMétrica
	noResponseCount2        descMétrica
	spinupRetryCount1       descMétrica
	spinupRetryCount2       descMétrica
	numberOfMediaErrors1    descMétrica
	numberOfMediaErrors2    descMétrica
	numberOfNonmediaErrors1 descMétrica
	numberOfNonmediaErrors2 descMétrica
	numberOfBlockReassigns1 descMétrica
	numberOfBlockReassigns2 descMétrica
	numberOfBadBlocks1      descMétrica
	numberOfBadBlocks2      descMétrica
	logger                  log.Logger
}

func init() {
	registerCollector("disk_statistics", NewDiskStatistics)
}

func NewDiskStatistics(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &diskStatistics{
		meSession: me,
		powerOnHours: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "power_on_hours"),
				"Power on hours of the disk", []string{"disk"}),
		},
		bytesPerSecond: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "bytes_per_second"),
				"Bytes per second of the disk", []string{"disk"}),
		},
		numberOfReads: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_reads"),
				"Number of reads of the disk", []string{"disk"}),
		},
		dataRead: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "data_read"),
				"Data read of the disk", []string{"disk"}),
		},
		dataWritten: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "data_written"),
				"Data written of the disk", []string{"disk"}),
		},
		lifetimeDataRead: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "lifetime_data_read"),
				"Lifetime data read of the disk", []string{"disk"}),
		},
		lifetimeDataWritten: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "lifetime_data_written"),
				"Lifetime data written of the disk", []string{"disk"}),
		},
		queueDepth: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "queue_depth"),
				"Queue depth of the disk", []string{"disk"}),
		},
		resetTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "reset_time"),
				"Reset time of the disk", []string{"disk"}),
		},
		startSampleTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "start_sample_time"),
				"Start sample time of the disk", []string{"disk"}),
		},
		stopSampleTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "stop_sample_time"),
				"Stop sample time of the disk", []string{"disk"}),
		},
		smartCount1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "smart_count_1"),
				"Smart count 1 of the disk", []string{"disk"}),
		},
		smartCount2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "smart_count_2"),
				"Smart count 2 of the disk", []string{"disk"}),
		},
		ioTimeoutCount1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "io_timeout_count_1"),
				"Io timeout count 1 of the disk", []string{"disk"}),
		},
		ioTimeoutCount2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "io_timeout_count_2"),
				"Io timeout count 2 of the disk", []string{"disk"}),
		},
		noResponseCount1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "no_response_count_1"),
				"No response count 1 of the disk", []string{"disk"}),
		},
		noResponseCount2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "no_response_count_2"),
				"No response count 2 of the disk", []string{"disk"}),
		},
		spinupRetryCount1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "spinup_retry_count_1"),
				"Spinup retry count 1 of the disk", []string{"disk"}),
		},
		spinupRetryCount2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "spinup_retry_count_2"),
				"Spinup retry count 2 of the disk", []string{"disk"}),
		},
		numberOfMediaErrors1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_media_errors_1"),
				"Number of media errors 1 of the disk", []string{"disk"}),
		},
		numberOfMediaErrors2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_media_errors_2"),
				"Number of media errors 2 of the disk", []string{"disk"}),
		},
		numberOfNonmediaErrors1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_nonmedia_errors_1"),
				"Number of nonmedia errors 1 of the disk", []string{"disk"}),
		},
		numberOfNonmediaErrors2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_nonmedia_errors_2"),
				"Number of nonmedia errors 2 of the disk", []string{"disk"}),
		},
		numberOfBlockReassigns1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_block_reassigns_1"),
				"Number of block reassigns 1 of the disk", []string{"disk"}),
		},
		numberOfBlockReassigns2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_block_reassigns_2"),
				"Number of block reassigns 2 of the disk", []string{"disk"}),
		},
		numberOfBadBlocks1: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_bad_blocks_1"),
				"Number of bad blocks 1 of the disk", []string{"disk"}),
		},
		numberOfBadBlocks2: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_bad_blocks_2"),
				"Number of bad blocks 2 of the disk", []string{"disk"}),
		},
		logger: logger,
	}, nil
}

func (ds diskStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := ds.meSession.DiskStatistics(); err != nil {
		return err
	}

	for _, disk := range ds.meSession.diskStatistic {
		ch <- prometheus.MustNewConstMetric(ds.powerOnHours.desc, ds.powerOnHours.tipo, float64(disk.PowerOnHours), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.bytesPerSecond.desc, ds.bytesPerSecond.tipo, float64(disk.BytesPerSecondNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfReads.desc, ds.numberOfReads.tipo, float64(disk.NumberOfReads), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.dataRead.desc, ds.dataRead.tipo, float64(disk.DataReadNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.dataWritten.desc, ds.dataWritten.tipo, float64(disk.DataWrittenNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.lifetimeDataRead.desc, ds.lifetimeDataRead.tipo, float64(disk.LifetimeDataReadNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.lifetimeDataWritten.desc, ds.lifetimeDataWritten.tipo, float64(disk.LifetimeDataWrittenNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.queueDepth.desc, ds.queueDepth.tipo, float64(disk.QueueDepth), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.resetTime.desc, ds.resetTime.tipo, float64(disk.ResetTimeNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.startSampleTime.desc, ds.startSampleTime.tipo, float64(disk.StartSampleTimeNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.stopSampleTime.desc, ds.stopSampleTime.tipo, float64(disk.StopSampleTimeNumeric), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.smartCount1.desc, ds.smartCount1.tipo, float64(disk.SmartCount1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.smartCount2.desc, ds.smartCount2.tipo, float64(disk.SmartCount2), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.ioTimeoutCount1.desc, ds.ioTimeoutCount1.tipo, float64(disk.IoTimeoutCount1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.ioTimeoutCount2.desc, ds.ioTimeoutCount2.tipo, float64(disk.IoTimeoutCount2), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.noResponseCount1.desc, ds.noResponseCount1.tipo, float64(disk.NoResponseCount1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.noResponseCount2.desc, ds.noResponseCount2.tipo, float64(disk.NoResponseCount2), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.spinupRetryCount1.desc, ds.spinupRetryCount1.tipo, float64(disk.SpinupRetryCount1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.spinupRetryCount2.desc, ds.spinupRetryCount2.tipo, float64(disk.SpinupRetryCount2), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfMediaErrors1.desc, ds.numberOfMediaErrors1.tipo, float64(disk.NumberOfMediaErrors1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfMediaErrors2.desc, ds.numberOfMediaErrors2.tipo, float64(disk.NumberOfMediaErrors2), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfNonmediaErrors1.desc, ds.numberOfNonmediaErrors1.tipo, float64(disk.NumberOfNonmediaErrors1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfNonmediaErrors2.desc, ds.numberOfNonmediaErrors2.tipo, float64(disk.NumberOfNonmediaErrors2), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfBlockReassigns1.desc, ds.numberOfBlockReassigns1.tipo, float64(disk.NumberOfBlockReassigns1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfBlockReassigns2.desc, ds.numberOfBlockReassigns2.tipo, float64(disk.NumberOfBlockReassigns2), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfBadBlocks1.desc, ds.numberOfBadBlocks1.tipo, float64(disk.NumberOfBadBlocks1), disk.DurableID)
		ch <- prometheus.MustNewConstMetric(ds.numberOfBadBlocks2.desc, ds.numberOfBadBlocks2.tipo, float64(disk.NumberOfBadBlocks2), disk.DurableID)
	}

	return nil
}
