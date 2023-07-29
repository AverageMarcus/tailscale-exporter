# tailscale-exporter

Exports Prometheus metrics for Tailscale devices and keys.

## Metrics

* `tailscale_devices_expiry_seconds_remaining` - The number of seconds remaining until the device authentication expires
* `tailscale_devices_expiry_time` - The timestamp (as Unix timestamp) that the device expires
* `tailscale_devices_update_available` - Whether the device can be updated to a newer version of Tailscale or not
* `tailscale_keys_expiry_seconds_remaining` - The number of seconds remaining until the key expires
* `tailscale_keys_expiry_time` - The timestamp (as Unix timestamp) that the key expires

## Configuration

The following environment variable can be used to configure the exporter:

* `TAILSCALE_API_KEY` - A valid Tailscale API key [Required]
* `TAILSCALE_TAILNET` - The Tailnet to export metrics for [Required]
* `PORT` - The port to run the exporter on (Defaults to `8080`)

## Running with Docker

```shell
export TAILSCALE_API_KEY="my-tailscale-api-key"
export TAILSCALE_TAILNET="my-tailnet.github"
docker run --rm -it -p 8080:8080 -e TAILSCALE_API_KEY -e TAILSCALE_TAILNET ghcr.io/averagemarcus/tailscale-exporter:latest
```

Then visit: [http://localhost:8080/metrics](http://localhost:8080/metrics)

## Deploying to Kubernetes

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: tailscale-exporter
  labels:
    app.kubernetes.io/name: tailscale-exporter
stringData:
  TAILSCALE_API_KEY: xxxx
  TAILSCALE_TAILNET: xxxx
---
apiVersion: v1
kind: Service
metadata:
  name: tailscale-exporter
  labels:
    app.kubernetes.io/name: tailscale-exporter
  annotations:
    prometheus.io/scrape: "true"
    prometheus.io/port: "8080"
spec:
  type: ClusterIP
  ports:
  - port: 8080
    targetPort: 8080
  selector:
    app: tailscale-exporter
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tailscale-exporter
  labels:
    app.kubernetes.io/name: tailscale-exporter
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: tailscale-exporter
  template:
    metadata:
      labels:
        app.kubernetes.io/name: tailscale-exporter
    spec:
      containers:
      - name: tailscale-exporter
        image: ghcr.io/averagemarcus/tailscale-exporter:latest
        imagePullPolicy: Always
        ports:
        - containerPort: 8080
          name: metrics
        envFrom:
        - secretRef:
            name: tailscale-exporter

```
