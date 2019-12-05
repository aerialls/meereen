package core

import(
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartMetricsEndpoint starts the HTTP endpoint for Prometheus metrics
func StartMetricsEndpoint(address string) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(address, nil)
}
