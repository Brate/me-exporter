package app

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"me_exporter/app/helpers"
	"me_exporter/collector"
	"net"
	"net/http"
)

var (
	mcMap = make(map[string]metricsCollector)
)

type metricsCollector struct {
	Instance  string
	meMetrics *collector.MeMetrics
	collector *collector.MeCollector

	logger log.Logger
}

type metricsController struct {
	logger log.Logger
}

func NewMetricsController(logger log.Logger) *metricsController {
	return &metricsController{
		logger: logger,
	}
}

func (m *metricsController) Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	instance := params.Get("instance")

	if err := validateIP(instance); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	col := getMetricsCollector(instance, m.logger)
	registry := prometheus.NewRegistry()
	err := registry.Register(col.collector)
	if err != nil {
		level.Error(m.logger).Log("msg", "Couldn't register collector:", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Couldn't register collector: %s", err)))
		return
	}

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		registry,
	}

	h := promhttp.InstrumentMetricHandler(
		registry,
		promhttp.HandlerFor(gatherers,
			promhttp.HandlerOpts{
				ErrorHandling: promhttp.ContinueOnError,
			}),
	)
	h.ServeHTTP(w, r)
}

func getMetricsCollector(ipStr string, logger log.Logger) metricsCollector {
	if mc, ok := mcMap[ipStr]; ok {
		return mc
	}
	metrics := collector.NewMeMetrics(ipStr, logger)
	col, _ := collector.NewMECollector(metrics, logger)
	mcMap[ipStr] = metricsCollector{
		Instance:  ipStr,
		collector: col,
		meMetrics: metrics,
		logger:    logger,
	}

	return mcMap[ipStr]
}
func validateIP(ipStr string) error {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return fmt.Errorf("Not a valid IP")
	}
	if ip.To4() != nil {
		return nil
	}
	if ip.To16() != nil {
		return nil
	}

	return fmt.Errorf("Not a valid IP")
}
