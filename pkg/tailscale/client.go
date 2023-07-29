package tailscale

import (
	"context"
	"fmt"
	"os"

	ts "github.com/tailscale/tailscale-client-go/tailscale"
)

type Client struct {
	tsClient *ts.Client
	tailnet  string
	ctx      context.Context
}

func New() (*Client, error) {
	apiKey := os.Getenv("TAILSCALE_API_KEY")
	tailnet := os.Getenv("TAILSCALE_TAILNET")

	if apiKey == "" {
		return nil, fmt.Errorf("TAILSCALE_API_KEY must be set")
	}
	if tailnet == "" {
		return nil, fmt.Errorf("TAILSCALE_TAILNET must be set")
	}

	client, err := ts.NewClient(apiKey, tailnet)
	if err != nil {
		return nil, err
	}

	return &Client{
		tsClient: client,
		tailnet:  tailnet,
		ctx:      context.Background(),
	}, nil
}

func (c *Client) GetTailnet() string {
	return c.tailnet
}

func (c *Client) GetKeys() ([]ts.Key, error) {
	allKeys := []ts.Key{}

	keys, err := c.tsClient.Keys(c.ctx)
	if err != nil {
		return nil, err
	} else {
		for _, k := range keys {
			key, err := c.tsClient.GetKey(c.ctx, k.ID)
			if err != nil {
				return nil, err
			}
			allKeys = append(allKeys, key)
		}
	}
	return allKeys, nil
}

func (c *Client) GetDevices() ([]ts.Device, error) {
	return c.tsClient.Devices(c.ctx)
}
