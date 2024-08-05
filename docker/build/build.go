package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/risersh/builder/docker/client"
	"github.com/risersh/builder/docker/images"
	"github.com/risersh/util/archiving"
)

type BuildArgs struct {
	GetClientArgs client.GetClientArgs
	Context       string
	Dockerfile    string
	Tags          []string
	NoCache       bool
	BuildArgs     map[string]*string
}

func Build(args BuildArgs) (image.Summary, error) {
	buildCtx, err := archiving.Tar(args.Context)
	if err != nil {
		return image.Summary{}, err
	}

	c, err := client.GetClient(args.GetClientArgs)
	if err != nil {
		return image.Summary{}, err
	}

	contextReader, err := os.Open(args.Context)
	if err != nil {
		return image.Summary{}, err
	}
	defer contextReader.Close()

	options := types.ImageBuildOptions{
		Dockerfile: args.Dockerfile,
		Tags:       args.Tags,
		NoCache:    args.NoCache,
		BuildArgs:  args.BuildArgs,
		Platform:   "linux/amd64",
		Memory:     1024 * 1024 * 1024, // 1GB
		CPUShares:  100,                // 100%
		CPUQuota:   10000,
		CPUPeriod:  10000,
	}

	response, err := c.Client.ImageBuild(context.Background(), buildCtx, options)
	if err != nil {
		return image.Summary{}, err
	}
	defer response.Body.Close()

	// Progress reporting.
	decoder := json.NewDecoder(response.Body)
	for {
		var msg jsonmessage.JSONMessage
		if err := decoder.Decode(&msg); err != nil {
			if err == io.EOF {
				break
			}
			return image.Summary{}, err
		}
		if msg.Error != nil {
			return image.Summary{}, fmt.Errorf(msg.Error.Message)
		}
		if msg.Stream != "" {
			fmt.Print(msg.Stream)
		}
	}

	images, err := images.ListImages(images.ListImagesArgs{
		GetClientArgs: args.GetClientArgs,
		Filters: []filters.KeyValuePair{
			{
				Key:   "reference",
				Value: args.Tags[0],
			},
		},
	})
	if err != nil {
		return image.Summary{}, err
	}

	return images[0], nil
}
