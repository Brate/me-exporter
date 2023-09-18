package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type diskGroupStatistics struct {
	meSession       *MeMetrics
	timeSinceReset  descMétrica
	timeSinceSample descMétrica
	numberOfReads   descMétrica
	numberOfWrites  descMétrica
	dataRead        descMétrica
	dataWritten     descMétrica
	bytesPerSecond  descMétrica
	iops            descMétrica
	avgRspTime      descMétrica
	avgReadRspTime  descMétrica
	AvgWriteRspTime descMétrica

	// Disk Group Statistics Paged
	dgdpserialNumber       descMétrica
	pagesAllocPerMinute    descMétrica
	pagesDeallocPerMinute  descMétrica
	pagesReclaimed         descMétrica
	numPagesUnmapPerMinute descMétrica
	logger                 log.Logger
}

func init() {
	registerCollector("disk_group_statistics", NewDiskGroupStatisticsCollector)
}

func NewDiskGroupStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &diskGroupStatistics{
		meSession: me,
		timeSinceReset: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "time_since_reset"),
				"Time since reset", []string{"time_since_reset"}),
		},
		timeSinceSample: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "time_since_sample"),
				"Time since sample", []string{"time_since_sample"}),
		},
		numberOfReads: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "number_of_reads"),
				"Number of reads", []string{"number_of_reads"}),
		},
		numberOfWrites: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "number_of_writes"),
				"Number of writes", []string{"number_of_writes"}),
		},
		dataRead: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "data_read"),
				"Data read numeric", []string{"data_read"}),
		},
		dataWritten: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "data_written"),
				"Data written numeric", []string{"data_written"}),
		},
		bytesPerSecond: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "bytes_per_second"),
				"Bytes per second numeric", []string{"bytes_per_second"}),
		},
		iops: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "iops"),
				"Iops", []string{"iops"}),
		},
		avgRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "avg_rsp_time"),
				"Avg rsp time", []string{"avg_rsp_time"}),
		},
		avgReadRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "avg_read_rsp_time"),
				"Avg read rsp time", []string{"avg_read_rsp_time"}),
		},
		AvgWriteRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics", "avg_write_rsp_time"),
				"Avg write rsp time", []string{"avg_write_rsp_time"}),
		},
		// Disk Group Statistics Paged
		dgdpserialNumber: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics_paged", "serial_number"),
				"Serial number of the disk group", []string{"serial_number"}),
		},
		pagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics_paged", "pages_alloc_per_minute"),
				"Pages alloc per minute", []string{"pages_alloc_per_minute"}),
		},
		pagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics_paged", "pages_dealloc_per_minute"),
				"Pages dealloc per minute", []string{"pages_dealloc_per_minute"}),
		},
		pagesReclaimed: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics_paged", "pages_reclaimed"),
				"Pages reclaimed", []string{"pages_reclaimed"}),
		},
		numPagesUnmapPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group_statistics_paged", "num_pages_unmap_per_minute"),
				"Num pages unmap per minute", []string{"num_pages_unmap_per_minute"}),
		},
		logger: logger,
	}, nil
}

func (dk *diskGroupStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := dk.meSession.DiskGroupStatistics(); err != nil {
		return err
	}

	s := dk.meSession.diskGroupStatistics

	ch <- prometheus.MustNewConstMetric(dk.timeSinceReset.desc, dk.timeSinceReset.tipo, float64(s.TimeSinceReset))
	ch <- prometheus.MustNewConstMetric(dk.timeSinceSample.desc, dk.timeSinceSample.tipo, float64(s.TimeSinceSample))
	ch <- prometheus.MustNewConstMetric(dk.numberOfReads.desc, dk.numberOfReads.tipo, float64(s.NumberOfReads))
	ch <- prometheus.MustNewConstMetric(dk.numberOfWrites.desc, dk.numberOfWrites.tipo, float64(s.NumberOfWrites))
	ch <- prometheus.MustNewConstMetric(dk.dataRead.desc, dk.dataRead.tipo, float64(s.DataReadNumeric))
	ch <- prometheus.MustNewConstMetric(dk.dataWritten.desc, dk.dataWritten.tipo, float64(s.DataWrittenNumeric))
	ch <- prometheus.MustNewConstMetric(dk.bytesPerSecond.desc, dk.bytesPerSecond.tipo, float64(s.BytesPerSecondNumeric))
	ch <- prometheus.MustNewConstMetric(dk.iops.desc, dk.iops.tipo, float64(s.Iops))
	ch <- prometheus.MustNewConstMetric(dk.avgRspTime.desc, dk.avgRspTime.tipo, float64(s.AvgRspTime))
	ch <- prometheus.MustNewConstMetric(dk.avgReadRspTime.desc, dk.avgReadRspTime.tipo, float64(s.AvgReadRspTime))
	ch <- prometheus.MustNewConstMetric(dk.AvgWriteRspTime.desc, dk.AvgWriteRspTime.tipo, float64(s.AvgWriteRspTime))

	// Disk Group Statistics Paged
	for _, dg := range s.DiskGroupStatisticsPaged {
		ch <- prometheus.MustNewConstMetric(dk.pagesAllocPerMinute.desc, dk.pagesAllocPerMinute.tipo, float64(dg.PagesAllocPerMinute))
		ch <- prometheus.MustNewConstMetric(dk.pagesDeallocPerMinute.desc, dk.pagesDeallocPerMinute.tipo, float64(dg.PagesDeallocPerMinute))
		ch <- prometheus.MustNewConstMetric(dk.pagesReclaimed.desc, dk.pagesReclaimed.tipo, float64(dg.PagesReclaimed))
		ch <- prometheus.MustNewConstMetric(dk.numPagesUnmapPerMinute.desc, dk.numPagesUnmapPerMinute.tipo, float64(dg.NumPagesUnmapPerMinute))
	}

	return nil
}
