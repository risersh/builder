package docker

import (
	"context"

	"github.com/docker/docker/api/types/image"
)

func PullImage(imageName string) error {
	c, err := GetClient(GetClientArgs{})
	if err != nil {
		return err
	}
	ctx := context.Background()
	options := image.PullOptions{
		Platform: "linux/amd64",
	}

	reader, err := c.Client.ImagePull(ctx, imageName, options)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Read the pull progress and handle it as needed
	// ...

	return nil
}
