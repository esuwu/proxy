package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"main/config"
	customHttp "main/http"
	"main/monitoring"
	"net/http"
	"strconv"
)

func CreatePrometheusMetricsMiddleware(config *config.Config) func(next http.Handler) http.Handler {
	if config.BackendID == 0 {
		log.Println("WARNING: using default BackendID value")
	}

	backendID := fmt.Sprintf("%d", config.BackendID)

	return func(next http.Handler) http.Handler {
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

				monitoring.Hits.With(
					prometheus.Labels{"path": template, "status": strconv.Itoa(statusCode), "backend": backendID},
				).Inc()
			}()

			timer := prometheus.NewTimer(monitoring.RequestDuration.With(
				prometheus.Labels{"path": template, "method": r.Method, "backend": backendID},
			))
			defer timer.ObserveDuration()

			next.ServeHTTP(newResponseWriter, r)
		})
	}
}
