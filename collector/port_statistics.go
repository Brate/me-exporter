package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type portStatistics struct {
	meSession *MeMetrics

	numberOfReads   descMétrica
	numberOfWrites  descMétrica
	dataRead        descMétrica
	dataWritten     descMétrica
	avgRspTime      descMétrica
	avgReadRspTime  descMétrica
	avgWriteRspTime descMétrica

	logger log.Logger
}

func init() {
	registerCollector("port_statistics", NewPortStatisticsCollector)
}

func NewPortStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &portStatistics{
		meSession: me,

		numberOfReads: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("port", "read_count"),
				"Number of reads on the port", []string{"id"}),
		},
		numberOfWrites: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("port", "write_count"),
				"Number of writes on the port", []string{"id"}),
		},
		dataRead: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("port", "data_read_bytes"),
				"Data read on the port", []string{"id"}),
		},
		dataWritten: descMétrica{prometheus.CounterValue,
			NewDescritor(
				NomeMetrica("port", "data_written_bytes"),
				"Data written on the port", []string{"id"}),
		},
		avgRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "avg_rsp_time_microseconds"),
				"Average response time on the port", []string{"id"}),
		},
		avgReadRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "avg_read_rsp_time_microseconds"),
				"Average read response time on the port", []string{"id"}),
		},
		avgWriteRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port", "avg_write_rsp_time_microseconds"),
				"Average write response time on the port", []string{"id"}),
		},
		logger: logger,
	}, nil
}

func (p portStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := p.meSession.PortStatistics(); err != nil {
		return err
	}

	for _, portStats := range p.meSession.portStatistics {
		ch <- p.numberOfReads.constMetric(float64(portStats.NumberOfReads), portStats.DurableID)
		ch <- p.numberOfWrites.constMetric(float64(portStats.NumberOfWrites), portStats.DurableID)
		ch <- p.dataRead.constMetric(float64(portStats.DataReadNumeric), portStats.DurableID)
		ch <- p.dataWritten.constMetric(float64(portStats.DataWrittenNumeric), portStats.DurableID)
		ch <- p.avgRspTime.constMetric(float64(portStats.AvgRspTime), portStats.DurableID)
		ch <- p.avgReadRspTime.constMetric(float64(portStats.AvgReadRspTime), portStats.DurableID)
		ch <- p.avgWriteRspTime.constMetric(float64(portStats.AvgWriteRspTime), portStats.DurableID)
	}

	return nil
}
