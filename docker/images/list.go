package images

import (
	"context"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/risersh/builder/docker/client"
)

// ListImagesArgs is the arguments for the ListImages function.
type ListImagesArgs struct {
	// GetClientArgs is the arguments for the GetClient function.
	client.GetClientArgs
	// All is whether to list all images.
	All bool
	// Filters is the filters for the images.
	Filters []filters.KeyValuePair
}

// ListImages lists the current images.
//
// Arguments:
//   - ListImagesArgs: The arguments for the ListImages function.
//
// Returns:
//   - []image.Summary: The list of matching images.
//   - error: The error that occurred while listing the images.
func ListImages(args ListImagesArgs) ([]image.Summary, error) {
	c, err := client.GetClient(args.GetClientArgs)
	if err != nil {
		return nil, err
	}
	return c.Client.ImageList(context.Background(), image.ListOptions{
		All:     args.All,
		Filters: filters.NewArgs(args.Filters...),
	})
}
