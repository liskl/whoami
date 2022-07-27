package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			routeViewsRoot.Inc()
		}
	}()
}

var (
	routeViewsRoot = promauto.NewCounter(prometheus.CounterOpts{
		Name: "whoami_route_root_total",
		Help: "The total number of views on /",
	})
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
	recordMetrics()
	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
		fmt.Fprintf(w, "I'm %s\n", hostname)
	})
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
