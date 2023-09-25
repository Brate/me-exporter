package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type diskGroupStatistics struct {
	meSession       *MeMetrics
	timeSinceReset  descMétrica
	numberOfReads   descMétrica
	numberOfWrites  descMétrica
	dataRead        descMétrica
	dataWritten     descMétrica
	avgRspTime      descMétrica
	avgReadRspTime  descMétrica
	AvgWriteRspTime descMétrica

	// Disk Group Statistics Paged
	pagesAllocPerMinute   descMétrica
	pagesDeallocPerMinute descMétrica
	pagesReclaimed        descMétrica
	pagesUnmapPerMinute   descMétrica
	logger                log.Logger
}

func init() {
	registerCollector("disk_group_statistics", NewDiskGroupStatisticsCollector)
}

func NewDiskGroupStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &diskGroupStatistics{
		meSession: me,
		timeSinceReset: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group", "time_since_reset_seconds"),
				"Time since reset", []string{"disk_group"}),
		},
		numberOfReads: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_group", "read_count"),
				"Number of reads", []string{"disk_group"}),
		},
		numberOfWrites: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_group", "write_count"),
				"Number of writes", []string{"disk_group"}),
		},
		dataRead: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_group", "data_read_bytes"),
				"Data read numeric", []string{"disk_group"}),
		},
		dataWritten: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_group", "data_written_bytes"),
				"Data written numeric", []string{"disk_group"}),
		},
		avgRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group", "avg_response_microseconds"),
				"Avg rsp time", []string{"disk_group"}),
		},
		avgReadRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group", "avg_read_response_microseconds"),
				"Avg read rsp time", []string{"disk_group"}),
		},
		AvgWriteRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group", "avg_write_response_microseconds"),
				"Avg write rsp time", []string{"disk_group"}),
		},

		// Disk Group Statistics Paged
		pagesAllocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group", "pages_alloc_per_minute"),
				"Pages allocations per minute", []string{"disk_group"}),
		},
		pagesDeallocPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group", "pages_dealloc_per_minute"),
				"Pages deallocations per minute", []string{"disk_group"}),
		},
		pagesReclaimed: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("disk_group", "pages_reclaimed"),
				"Pages reclaimed", []string{"disk_group"}),
		},
		pagesUnmapPerMinute: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("disk_group", "pages_unmap_per_minute"),
				"Num pages unmapped per minute", []string{"disk_group"}),
		},
		logger: logger,
	}, nil
}

func (dk *diskGroupStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := dk.meSession.DiskGroupStatistics(); err != nil {
		return err
	}

	for _, dgs := range dk.meSession.diskGroupsStatistics {
		ch <- prometheus.MustNewConstMetric(dk.timeSinceReset.desc, dk.timeSinceReset.tipo, float64(dgs.TimeSinceReset), dgs.Name)
		ch <- prometheus.MustNewConstMetric(dk.numberOfReads.desc, dk.numberOfReads.tipo, float64(dgs.NumberOfReads), dgs.Name)
		ch <- prometheus.MustNewConstMetric(dk.numberOfWrites.desc, dk.numberOfWrites.tipo, float64(dgs.NumberOfWrites), dgs.Name)
		ch <- prometheus.MustNewConstMetric(dk.dataRead.desc, dk.dataRead.tipo, float64(dgs.DataReadNumeric), dgs.Name)
		ch <- prometheus.MustNewConstMetric(dk.dataWritten.desc, dk.dataWritten.tipo, float64(dgs.DataWrittenNumeric), dgs.Name)
		ch <- prometheus.MustNewConstMetric(dk.avgRspTime.desc, dk.avgRspTime.tipo, float64(dgs.AvgRspTime), dgs.Name)
		ch <- prometheus.MustNewConstMetric(dk.avgReadRspTime.desc, dk.avgReadRspTime.tipo, float64(dgs.AvgReadRspTime), dgs.Name)
		ch <- prometheus.MustNewConstMetric(dk.AvgWriteRspTime.desc, dk.AvgWriteRspTime.tipo, float64(dgs.AvgWriteRspTime), dgs.Name)

		// Disk Group Statistics Paged
		for _, dg := range dgs.DiskGroupStatisticsPaged {
			ch <- prometheus.MustNewConstMetric(dk.pagesAllocPerMinute.desc, dk.pagesAllocPerMinute.tipo, float64(dg.PagesAllocPerMinute), dgs.Name)
			ch <- prometheus.MustNewConstMetric(dk.pagesDeallocPerMinute.desc, dk.pagesDeallocPerMinute.tipo, float64(dg.PagesDeallocPerMinute), dgs.Name)
			ch <- prometheus.MustNewConstMetric(dk.pagesReclaimed.desc, dk.pagesReclaimed.tipo, float64(dg.PagesReclaimed), dgs.Name)
			ch <- prometheus.MustNewConstMetric(dk.pagesUnmapPerMinute.desc, dk.pagesUnmapPerMinute.tipo, float64(dg.NumPagesUnmapPerMinute), dgs.Name)
		}
	}
	return nil
}
