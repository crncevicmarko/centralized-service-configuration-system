package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Initial count.
	//currentCount = 0

	httpHits = createCounter("my_app_http_hit_total", "Total number of http hits.")

	httpHitsGetConfig        = createCounter("my_app_http_hit_config_get", "Total number of http hits for GET Config.")
	httpHitsGetConfigByLabel = createCounter("my_app_http_hit_config_get_label", "Total number of http hits for GET Config by label.")
	httpHitsCreateConfig     = createCounter("my_app_http_hit_config_all", "Total number of http hits for POST Config.")
	httpHitsDeleteConfig     = createCounter("my_app_http_hit_config_delete", "Total number of http hits for DELETE Config.")

	httpHitsGetGroup    = createCounter("my_app_http_hit_group_get", "Total number of http hits for GET group.")
	httpHitsCreateGroup = createCounter("my_app_http_hit_group_post", "Total number of http hits for POST group.")
	httpHitsUpdateGroup = createCounter("my_app_http_hit_group_update", "Total number of http hits for PUT group")
	httpHitsDeleteGroup = createCounter("my_app_http_hit_group_delete", "Total number of http hits for DELETE group.")

	httpHitsGetAll = createCounter("my_app_http_hit_config_get_all", "Total number of http hits for GET all.")

	// The Prometheus metric that will be exposed.
	//httpHits = prometheus.NewCounter(
	//	prometheus.CounterOpts{
	//		Name: "my_app_http_hit_total",
	//		Help: "Total number of http hits.",
	//	},
	//)

	// Add all metrics that will be resisted
	metricsList = []prometheus.Collector{
		httpHits,
		httpHitsGetConfig,
		httpHitsGetConfigByLabel,
		httpHitsCreateConfig,
		httpHitsDeleteConfig,
		httpHitsGetGroup,
		httpHitsCreateGroup,
		httpHitsUpdateGroup,
		httpHitsDeleteGroup,
		httpHitsGetAll,
	}

	// Prometheus Registry to register metrics.
	prometheusRegistry = prometheus.NewRegistry()
)

func createCounter(name string, help string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
	)
}

func init() {
	// Register metrics that will be exposed.

	prometheusRegistry.MustRegister(metricsList...)
}

func metricsHandler() http.Handler {
	return promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})
}

func count(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		httpHits.Inc()
		f(w, r) // original function call
	}
}
