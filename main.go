package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"main/middleware"
	"math/rand"
	"net/http"
	"time"
)



func normalInverse(mu float32, sigma float32) float32 {
	return float32(rand.NormFloat64() * float64(sigma) + float64(mu))
}


func handleHTTP(w http.ResponseWriter, req *http.Request) {

	timeToSleep := int64(normalInverse(500, 200))
	if timeToSleep <= 0 {
		timeToSleep = 200
	}
	dur := time.Duration(timeToSleep) * time.Millisecond
	fmt.Println(dur)
	time.Sleep(dur)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("kek"))
}


func main() {
	mainrouter := mux.NewRouter()
	mainrouter.HandleFunc("/", handleHTTP).Methods(http.MethodGet)
	mainrouter.Use(middleware.PrometheusMetricsMiddleware)


	metricsRouter := mainrouter.PathPrefix("/metrics").Subrouter()
	metricsRouter.Handle("", promhttp.Handler()).Methods(http.MethodGet)



	log.Println("server started on 8081 port")
	log.Fatal(http.ListenAndServe(":8081", mainrouter))
}