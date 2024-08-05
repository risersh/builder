package docker

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/risersh/builder/docker/client"
	"github.com/risersh/builder/test"
	"github.com/risersh/util/variables"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BuildSuite struct {
	suite.Suite
}

func (s *BuildSuite) SetupTest() {
	test.Setup()
}

func TestBuildSuite(t *testing.T) {
	suite.Run(t, new(BuildSuite))
}

func (s *BuildSuite) TestBuild() {
	dir, err := os.Getwd()
	assert.NoError(s.T(), err)
	log.Printf("Current working directory: %s", dir)

	image, err := Build(BuildArgs{
		GetClientArgs: client.GetClientArgs{
			Host: test.GetDockerHost(),
		},
		Context:    "../../test",
		Dockerfile: "dockerfile",
		Tags:       []string{"test"},
		BuildArgs: map[string]*string{
			"FOO": variables.ToPtr("bar"),
		},
		NoCache: true,
	})
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), image.Size, int64(0))

	multilog.Debug("test.build", "built image", map[string]interface{}{
		"id":   image.ID,
		"size": fmt.Sprintf("%d MB", image.Size/1024/1024),
	})
}
