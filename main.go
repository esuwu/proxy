package main

import (
	"fmt"
	"log"
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
	server := &http.Server{
		Addr: ":8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handleHTTP(w, r)
		}),
	}

	log.Println("server started on 8080 port")
	log.Fatal(server.ListenAndServe())
}