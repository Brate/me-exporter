package main

import (
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/promlog"
	"me_exporter/app"
	"me_exporter/collector"
	"me_exporter/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	sha1ver string
	Version = "0.0.0.dev"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGHUP:
				config.LoadConfig()
				collector.FlushMECollectors()
			}
		}
	}()

	logger := promlog.New(&promlog.Config{})
	controller := app.NewMetricsController(logger)
	//col, _ := collector.NewMECollector(logger)
	//prometheus.MustRegister(col)

	router := mux.NewRouter()
	router.HandleFunc("/metrics", controller.Handler).Methods(http.MethodGet)
	err := http.ListenAndServe(":10005", router)
	if err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		return
	}
}
