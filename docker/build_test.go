package docker

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/risersh/builder/test"
	"github.com/risersh/util/variables"
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

	contextPath := filepath.Join(dir, "..", "test")
	dockerfilePath := filepath.Join(contextPath, "dockerfile")

	log.Printf("Context path: %s", contextPath)
	log.Printf("Dockerfile path: %s", dockerfilePath)

	// Check if the Dockerfile exists
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		log.Printf("Dockerfile does not exist at: %s", dockerfilePath)
	} else {
		log.Printf("Dockerfile exists at: %s", dockerfilePath)
	}

	err = Build(BuildArgs{
		GetClientArgs: GetClientArgs{
			Host: test.GetDockerHost(),
		},
		Context:    contextPath,
		Dockerfile: "dockerfile",
		Tags:       []string{"test"},
		BuildArgs: map[string]*string{
			"FOO": variables.ToPtr("bar"),
		},
		NoCache: true,
	})
	assert.NoError(s.T(), err)
}
