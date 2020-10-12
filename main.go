package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"main/config"
	"main/middleware"
	"main/monitoring"
	"math/rand"
	"net/http"
	"time"
)

var replyPhrase string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func normalInverse(mu float32, sigma float32) float32 {
	return float32(rand.NormFloat64()*float64(sigma) + float64(mu))
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	timeToSleep := int64(normalInverse(500, 200))
	if timeToSleep <= 0 {
		timeToSleep = 200
	}
	dur := time.Duration(timeToSleep) * time.Millisecond

	log.Println("time duration:", dur)

	time.Sleep(dur)
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(replyPhrase)); err != nil {
		log.Println(err)
	}
}

func main() {
	prometheus.MustRegister(monitoring.Hits, monitoring.RequestDuration)

	conf := config.New()
	mainRouter := mux.NewRouter()

	mainRouter.Use(middleware.CreatePrometheusMetricsMiddleware(&conf))

	replyPhrase = fmt.Sprintf("Hello World!\n\rFrom backend #%d\n\r", conf.BackendID)

	mainRouter.HandleFunc("/", handleHTTP).Methods(http.MethodGet)
	mainRouter.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	log.Printf("server started on %s", conf.ServerEndpoint)
	log.Fatal(http.ListenAndServe(conf.ServerEndpoint, mainRouter))
}
