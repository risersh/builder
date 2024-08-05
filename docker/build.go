package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/risersh/util/archiving"
)

type BuildArgs struct {
	GetClientArgs
	Context    string
	Dockerfile string
	Tags       []string
	NoCache    bool
	BuildArgs  map[string]*string
}

func Build(args BuildArgs) error {

	buildCtx, err := archiving.Tar(args.Context)
	if err != nil {
		return err
	}

	c, err := GetClient(args.GetClientArgs)
	if err != nil {
		return err
	}

	ctx := context.Background()

	contextReader, err := os.Open(args.Context)
	if err != nil {
		return err
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

	response, err := c.Client.ImageBuild(ctx, buildCtx, options)
	if err != nil {
		return err
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
			return err
		}
		if msg.Error != nil {
			return fmt.Errorf(msg.Error.Message)
		}
		if msg.Stream != "" {
			fmt.Print(msg.Stream)
		}
	}

	return nil
}
