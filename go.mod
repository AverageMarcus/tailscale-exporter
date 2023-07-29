module tailscale-exporter

go 1.20

replace github.com/tailscale/tailscale-client-go => github.com/AverageMarcus/tailscale-client-go v0.0.0-20230729201523-e4a8f131596d

require (
	github.com/prometheus/client_golang v1.16.0
	github.com/tailscale/tailscale-client-go v1.9.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/tailscale/hujson v0.0.0-20220506213045-af5ed07155e5 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/oauth2 v0.8.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
