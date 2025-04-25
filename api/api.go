package api

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var pingCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "ping_total",
	Help: "Total number of pings",
})

func ping(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong\n", pingCounter)
}
func RunServer(port string) {
	prometheus.MustRegister(pingCounter)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", ping)
	address := ":" + port
	fmt.Println("listening on", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Server failed to start: %v", err)
	}
}
