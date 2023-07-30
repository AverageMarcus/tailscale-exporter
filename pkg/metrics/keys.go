package metrics

import (
	"fmt"
	"tailscale-exporter/pkg/tailscale"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func collectKeys(client *tailscale.Client) []prometheus.Collector {
	keyExpiry := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tailscale_keys_expiry_time",
			Help: "The expiry time of auth keys",
			ConstLabels: prometheus.Labels{
				"tailnet": client.GetTailnet(),
			},
		},
		[]string{"id", "created", "description", "type"},
	)

	keySecondsRemaining := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tailscale_keys_expiry_seconds_remaining",
			Help: "The number of seconds remaining until a key expires",
			ConstLabels: prometheus.Labels{
				"tailnet": client.GetTailnet(),
			},
		},
		[]string{"id", "created", "description", "type"},
	)

	go func() {
		for {
			keys, err := client.GetKeys()
			if err != nil {
				fmt.Println("Failed to get keys: ", err)
			} else {
				// Reset gauges so we don't leave old keys around
				keyExpiry.Reset()
				keySecondsRemaining.Reset()

				for _, key := range keys {
					remainingSeconds := time.Until(key.Expires).Seconds()
					keyType := "auth_key"
					if key.Capabilities == nil {
						keyType = "api_access_token"
					}
					keyExpiry.With(prometheus.Labels{"id": key.ID, "created": key.Created.String(), "description": key.Description, "type": keyType}).Set(float64(key.Expires.Unix()))
					keySecondsRemaining.With(prometheus.Labels{"id": key.ID, "created": key.Created.String(), "description": key.Description, "type": keyType}).Set(remainingSeconds)
				}
			}

			time.Sleep(60 * time.Second)
		}
	}()

	return []prometheus.Collector{keyExpiry, keySecondsRemaining}
}
