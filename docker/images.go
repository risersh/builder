package docker

import (
	"context"

	"github.com/docker/docker/api/types/image"
)

func ListImages() ([]image.Summary, error) {
	c, err := GetClient(GetClientArgs{})
	if err != nil {
		return nil, err
	}
	return c.Client.ImageList(context.Background(), image.ListOptions{})
}
