package middleware

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	customHttp "main/http"
	"main/monitoring"
	"net/http"
	"strconv"
)

func PrometheusMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newResponseWriter := &customHttp.ResponseWriter{ResponseWriter: w}
		template, err := mux.CurrentRoute(r).GetPathTemplate()
		if err != nil {
			template = r.URL.Path
		}

		defer func() {
			statusCode := newResponseWriter.GetStatusCode()
			if statusCode == 0 {
				statusCode = http.StatusOK
			}
			monitoring.Hits.WithLabelValues(strconv.Itoa(statusCode), template).Inc()
		}()

		timer := prometheus.NewTimer(monitoring.RequestDuration.With(
			prometheus.Labels{"path": template, "method": r.Method},
		))
		defer timer.ObserveDuration()

		next.ServeHTTP(newResponseWriter, r)
	})
}
