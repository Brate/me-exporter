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

type MetricsController struct {
	logger log.Logger
}

func NewMetricsController(logger log.Logger) *MetricsController {
	return &MetricsController{
		logger: logger,
	}
}

func (m *MetricsController) Handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	target := params.Get("target")

	if err := validateIP(target); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	col := getMetricsCollector(target, m.logger)
	registry := prometheus.NewRegistry()
	err := registry.Register(col.collector)
	if err != nil {
		_ = level.Error(m.logger).Log("msg", "Couldn't register collector:", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("Couldn't register collector: %s", err)))
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
	mc, ok := mcMap[ipStr]
	if ok && len(mc.collector.Coletores) > 0 {
		return mc
	}
	metrics := collector.NewMeMetrics(ipStr, logger)
	col, _ := collector.NewMECollectors(ipStr, metrics, logger)
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
		return fmt.Errorf("not a valid IP")
	}
	if ip.To4() != nil {
		return nil
	}
	if ip.To16() != nil {
		return nil
	}

	return fmt.Errorf("not a valid IP")
}
