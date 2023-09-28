package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type diskStatistics struct {
	meSession    *MeMetrics
	powerOnHours descMétrica
	//bytesPerSecond      descMétrica
	readCount           descMétrica
	writeCount          descMétrica
	dataRead            descMétrica
	dataWritten         descMétrica
	lifetimeDataRead    descMétrica
	lifetimeDataWritten descMétrica
	queueDepth          descMétrica
	//resetTime              descMétrica
	//startSampleTime        descMétrica
	//stopSampleTime         descMétrica
	smartCount             descMétrica
	ioTimeoutCount         descMétrica
	noResponseCount        descMétrica
	spinupRetryCount       descMétrica
	numberOfMediaErrors    descMétrica
	numberOfNonmediaErrors descMétrica
	numberOfBlockReassigns descMétrica
	numberOfBadBlocks      descMétrica
	logger                 log.Logger
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
		//bytesPerSecond: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disk_statistics", "bytes_per_second"),
		//		"Bytes per second of the disk", []string{"disk"}),
		//},
		readCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "read_count"),
				"How many time a read command was issued to this disk", []string{"disk"}),
		},
		writeCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "write_count"),
				"How many time a write command was issued to this disk", []string{"disk"}),
		},
		dataRead: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "data_read_bytes"),
				"Data read of the disk", []string{"disk"}),
		},
		dataWritten: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "data_written_bytes"),
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
		//resetTime: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disk_statistics", "reset_time"),
		//		"Reset time of the disk", []string{"disk"}),
		//},
		//startSampleTime: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disk_statistics", "start_sample_time"),
		//		"Start sample time of the disk", []string{"disk"}),
		//},
		//stopSampleTime: descMétrica{prometheus.GaugeValue,
		//	NewDescritor(
		//		NomeMetrica("disk_statistics", "stop_sample_time"),
		//		"Stop sample time of the disk", []string{"disk"}),
		//},
		smartCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "smart_count"),
				"Smart count  of the disk", []string{"disk", "port"}),
		},
		ioTimeoutCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "io_timeout_count"),
				"Io timeout count  of the disk", []string{"disk", "port"}),
		},
		noResponseCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "no_response_count"),
				"No response count  of the disk", []string{"disk", "port"}),
		},
		spinupRetryCount: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "spinup_retry_count"),
				"Spinup retry count  of the disk", []string{"disk", "port"}),
		},
		numberOfMediaErrors: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_media_errors"),
				"Number of media errors  of the disk", []string{"disk", "port"}),
		},
		numberOfNonmediaErrors: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_nonmedia_errors"),
				"Number of nonmedia errors  of the disk", []string{"disk", "port"}),
		},
		numberOfBlockReassigns: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_block_reassigns"),
				"Number of block reassigns  of the disk", []string{"disk", "port"}),
		},
		numberOfBadBlocks: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_statistics", "number_of_bad_blocks"),
				"Number of bad blocks  of the disk", []string{"disk", "port"}),
		},
		logger: logger,
	}, nil
}

func (ds diskStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := ds.meSession.DiskStatistics(); err != nil {
		return err
	}

	for _, disk := range ds.meSession.diskStatistic {
		ch <- prometheus.MustNewConstMetric(ds.powerOnHours.desc, ds.powerOnHours.tipo, float64(disk.PowerOnHours), disk.Location)
		//ch <- prometheus.MustNewConstMetric(ds.bytesPerSecond.desc, ds.bytesPerSecond.tipo, float64(disk.BytesPerSecondNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.readCount.desc, ds.readCount.tipo, float64(disk.NumberOfReads), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.writeCount.desc, ds.writeCount.tipo, float64(disk.NumberOfWrites), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.dataRead.desc, ds.dataRead.tipo, float64(disk.DataReadNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.dataWritten.desc, ds.dataWritten.tipo, float64(disk.DataWrittenNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.lifetimeDataRead.desc, ds.lifetimeDataRead.tipo, float64(disk.LifetimeDataReadNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.lifetimeDataWritten.desc, ds.lifetimeDataWritten.tipo, float64(disk.LifetimeDataWrittenNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.queueDepth.desc, ds.queueDepth.tipo, float64(disk.QueueDepth), disk.Location)
		//ch <- prometheus.MustNewConstMetric(ds.resetTime.desc, ds.resetTime.tipo, float64(disk.ResetTimeNumeric), disk.Location)
		//ch <- prometheus.MustNewConstMetric(ds.startSampleTime.desc, ds.startSampleTime.tipo, float64(disk.StartSampleTimeNumeric), disk.Location)
		//ch <- prometheus.MustNewConstMetric(ds.stopSampleTime.desc, ds.stopSampleTime.tipo, float64(disk.StopSampleTimeNumeric), disk.Location)
		ch <- prometheus.MustNewConstMetric(ds.smartCount.desc, ds.smartCount.tipo, float64(disk.SmartCount1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.smartCount.desc, ds.smartCount.tipo, float64(disk.SmartCount2), disk.Location, "2")
		ch <- prometheus.MustNewConstMetric(ds.ioTimeoutCount.desc, ds.ioTimeoutCount.tipo, float64(disk.IoTimeoutCount1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.ioTimeoutCount.desc, ds.ioTimeoutCount.tipo, float64(disk.IoTimeoutCount2), disk.Location, "2")
		ch <- prometheus.MustNewConstMetric(ds.noResponseCount.desc, ds.noResponseCount.tipo, float64(disk.NoResponseCount1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.noResponseCount.desc, ds.noResponseCount.tipo, float64(disk.NoResponseCount2), disk.Location, "2")
		ch <- prometheus.MustNewConstMetric(ds.spinupRetryCount.desc, ds.spinupRetryCount.tipo, float64(disk.SpinupRetryCount1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.spinupRetryCount.desc, ds.spinupRetryCount.tipo, float64(disk.SpinupRetryCount2), disk.Location, "2")
		ch <- prometheus.MustNewConstMetric(ds.numberOfMediaErrors.desc, ds.numberOfMediaErrors.tipo, float64(disk.NumberOfMediaErrors1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.numberOfMediaErrors.desc, ds.numberOfMediaErrors.tipo, float64(disk.NumberOfMediaErrors2), disk.Location, "2")
		ch <- prometheus.MustNewConstMetric(ds.numberOfNonmediaErrors.desc, ds.numberOfNonmediaErrors.tipo, float64(disk.NumberOfNonmediaErrors1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.numberOfNonmediaErrors.desc, ds.numberOfNonmediaErrors.tipo, float64(disk.NumberOfNonmediaErrors2), disk.Location, "2")
		ch <- prometheus.MustNewConstMetric(ds.numberOfBlockReassigns.desc, ds.numberOfBlockReassigns.tipo, float64(disk.NumberOfBlockReassigns1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.numberOfBlockReassigns.desc, ds.numberOfBlockReassigns.tipo, float64(disk.NumberOfBlockReassigns2), disk.Location, "2")
		ch <- prometheus.MustNewConstMetric(ds.numberOfBadBlocks.desc, ds.numberOfBadBlocks.tipo, float64(disk.NumberOfBadBlocks1), disk.Location, "1")
		ch <- prometheus.MustNewConstMetric(ds.numberOfBadBlocks.desc, ds.numberOfBadBlocks.tipo, float64(disk.NumberOfBadBlocks2), disk.Location, "2")
	}

	return nil
}
