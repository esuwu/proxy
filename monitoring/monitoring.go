package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "service_request_duration",
		Buckets: prometheus.LinearBuckets(0.01, 0.01, 10),
	}, []string{"path", "method", "backend"})

	Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "service_request_path_hits",
	}, []string{"path", "status", "backend"})
)
