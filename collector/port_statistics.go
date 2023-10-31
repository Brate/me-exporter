package collector

import (
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type portStatistics struct {
	//All metrics have durable id in your metrics
	meSession       *MeMetrics
	bytesPerSecond  descMétrica
	numberOfReads   descMétrica
	numberOfWrites  descMétrica
	dataRead        descMétrica
	dataWritten     descMétrica
	avgRspTime      descMétrica
	avgReadRspTime  descMétrica
	avgWriteRspTime descMétrica
	resetTime       descMétrica
	startSampleTime descMétrica
	stopSampleTime  descMétrica
	logger          log.Logger
}

func init() {
	registerCollector("port_statistics", NewPortStatisticsCollector)
}

func NewPortStatisticsCollector(me *MeMetrics, logger log.Logger) (Coletor, error) {
	return &portStatistics{
		meSession: me,
		bytesPerSecond: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "bytes_per_second"),
				"Bytes per second on the port", []string{"durableID"}),
		},
		numberOfReads: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "number_of_reads"),
				"Number of reads on the port", []string{"durableID"}),
		},
		numberOfWrites: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "number_of_writes"),
				"Number of writes on the port", []string{"durableID"}),
		},
		dataRead: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "data_read"),
				"Data read on the port", []string{"durableID"}),
		},
		dataWritten: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "data_written"),
				"Data written on the port", []string{"durableID"}),
		},
		avgRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "avg_rsp_time"),
				"Average response time on the port", []string{"durableID"}),
		},
		avgReadRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "avg_read_rsp_time"),
				"Average read response time on the port", []string{"durableID"}),
		},
		avgWriteRspTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "avg_write_rsp_time"),
				"Average write response time on the port", []string{"durableID"}),
		},
		resetTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "reset_time"),
				"Reset time on the port", []string{"durableID"}),
		},
		startSampleTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "start_sample_time"),
				"Start sample time on the port", []string{"port", "durableID"}),
		},
		stopSampleTime: descMétrica{prometheus.GaugeValue,
			NewDescritor(
				NomeMetrica("port_statistics", "stop_sample_time"),
				"Stop sample time on the port", []string{"durableID"}),
		},
		logger: logger,
	}, nil
}

func (p portStatistics) Update(ch chan<- prometheus.Metric) error {
	if err := p.meSession.PortStatistics(); err != nil {
		return err
	}

	for _, portStats := range p.meSession.portStatistics {
		ch <- p.bytesPerSecond.constMetric(float64(portStats.BytesPerSecondNumeric), portStats.DurableID)
		ch <- p.numberOfReads.constMetric(float64(portStats.NumberOfReads), portStats.DurableID)
		ch <- p.numberOfWrites.constMetric(float64(portStats.NumberOfWrites), portStats.DurableID)
		ch <- p.dataRead.constMetric(float64(portStats.DataReadNumeric), portStats.DurableID)
		ch <- p.dataWritten.constMetric(float64(portStats.DataWrittenNumeric), portStats.DurableID)
		ch <- p.avgRspTime.constMetric(float64(portStats.AvgRspTime), portStats.DurableID)
		ch <- p.avgReadRspTime.constMetric(float64(portStats.AvgReadRspTime), portStats.DurableID)
		ch <- p.avgWriteRspTime.constMetric(float64(portStats.AvgWriteRspTime), portStats.DurableID)
		ch <- p.resetTime.constMetric(float64(portStats.ResetTimeNumeric), portStats.DurableID)
		ch <- p.startSampleTime.constMetric(float64(portStats.StartSampleTimeNumeric), portStats.DurableID)
		ch <- p.stopSampleTime.constMetric(float64(portStats.StopSampleTimeNumeric), portStats.DurableID)
	}

	return nil
}
