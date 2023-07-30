package metrics

import (
	"fmt"
	"tailscale-exporter/pkg/tailscale"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func collectDevices(client *tailscale.Client) []prometheus.Collector {
	deviceExpiry := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tailscale_devices_expiry_time",
			Help: "The expiry time of devices authentication",
			ConstLabels: prometheus.Labels{
				"tailnet": client.GetTailnet(),
			},
		},
		[]string{"id", "created", "name"},
	)

	deviceSecondsRemaining := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tailscale_devices_expiry_seconds_remaining",
			Help: "The number of seconds remaining until a device expires",
			ConstLabels: prometheus.Labels{
				"tailnet": client.GetTailnet(),
			},
		},
		[]string{"id", "created", "name"},
	)

	deviceUpdateAvailable := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tailscale_devices_update_available",
			Help: "If the device can be updated (1) or is running the latest version of Tailscale (0)",
			ConstLabels: prometheus.Labels{
				"tailnet": client.GetTailnet(),
			},
		},
		[]string{"id", "created", "name", "version"},
	)

	go func() {
		for {
			devices, err := client.GetDevices()
			if err != nil {
				fmt.Println("Failed to get devices: ", err)
			} else {
				// Reset gauges so we don't leave old devices around
				deviceExpiry.Reset()
				deviceSecondsRemaining.Reset()
				deviceUpdateAvailable.Reset()

				for _, device := range devices {
					if !device.KeyExpiryDisabled {
						remainingSeconds := time.Until(device.Expires.Time).Seconds()

						deviceExpiry.With(prometheus.Labels{"id": device.ID, "created": device.Created.String(), "name": device.Name}).Set(float64(device.Expires.Unix()))
						deviceSecondsRemaining.With(prometheus.Labels{"id": device.ID, "created": device.Created.String(), "name": device.Name}).Set(remainingSeconds)
					}

					updateAvailable := 0.0
					if device.UpdateAvailable {
						updateAvailable = 1.0
					}
					deviceUpdateAvailable.With(prometheus.Labels{"id": device.ID, "created": device.Created.String(), "name": device.Name, "version": device.ClientVersion}).Set(updateAvailable)
				}
			}

			time.Sleep(60 * time.Second)
		}
	}()

	return []prometheus.Collector{deviceExpiry, deviceSecondsRemaining, deviceUpdateAvailable}
}
