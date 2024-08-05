package images

import (
	"context"

	"github.com/docker/docker/api/types/image"
	"github.com/risersh/builder/docker/client"
)

type RemoveArgs struct {
	client.GetClientArgs
	ImageID string
}

func Remove(args RemoveArgs) ([]image.DeleteResponse, error) {
	c, err := client.GetClient(args.GetClientArgs)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.ImageRemove(context.Background(), args.ImageID, image.RemoveOptions{
		Force:         true,
		PruneChildren: true,
	})

	return res, err
}
