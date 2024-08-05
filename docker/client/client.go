package client

import (
	"context"
	"errors"
	"time"

	"github.com/docker/docker/client"
)

type GetClientArgs struct {
	Host string
}

type Client struct {
	Client *client.Client
}

func (c *Client) Ping() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := c.Client.Ping(ctx)
	return err == nil
}

func GetClient(args GetClientArgs) (*Client, error) {
	c := &Client{}

	if args.Host == "" {
		args.Host = "file:///var/lib/docker.sock"
	}

	var err error
	// Create a new Docker client with default configuration
	c.Client, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
		client.WithHost(args.Host),
		client.WithTimeout(time.Second*30),
	)
	if err != nil {
		return nil, err
	}

	if !c.Ping() {
		return nil, errors.New("docker daemon is unreachable")
	}

	return c, nil
}
