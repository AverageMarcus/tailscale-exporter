package metrics

import (
	"tailscale-exporter/pkg/tailscale"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

var defaultCollectors = []prometheus.Collector{
	collectors.NewGoCollector(),
	collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
}

func Collect(client *tailscale.Client, reg *prometheus.Registry) {
	reg.MustRegister(defaultCollectors...)
	reg.MustRegister(collectKeys(client)...)
	reg.MustRegister(collectDevices(client)...)
}
