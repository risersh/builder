package images

import (
	"testing"

	"github.com/risersh/builder/docker/client"
	"github.com/risersh/builder/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RemoveSuite struct {
	suite.Suite
}

func (s *RemoveSuite) SetupTest() {
	test.Setup()
}

func TestRemoveSuite(t *testing.T) {
	suite.Run(t, new(RemoveSuite))
}

func (s *RemoveSuite) TestRemove() {
	res, err := Remove(RemoveArgs{
		GetClientArgs: client.GetClientArgs{
			Host: test.GetDockerHost(),
		},
		ImageID: "test",
	})
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(res), 0)
}
