package main

import (
	"log"
	"net/http"
	"os"
	"tailscale-exporter/pkg/metrics"
	"tailscale-exporter/pkg/tailscale"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr string

func init() {
	port := os.Getenv("PORT")
	if port != "" {
		addr = ":" + port
	} else {
		addr = ":8080"
	}
}

func main() {
	client, err := tailscale.New()
	if err != nil {
		log.Fatal(err)
	}

	reg := prometheus.NewRegistry()
	metrics.Collect(client, reg)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Printf("tailscale-exporter")
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
