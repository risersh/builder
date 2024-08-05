package images

import (
	"testing"

	"github.com/docker/docker/api/types/filters"
	"github.com/risersh/builder/docker/client"
	"github.com/risersh/builder/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ImagesSuite struct {
	suite.Suite
}

func (s *ImagesSuite) SetupTest() {
	test.Setup()
}

func TestImagesSuite(t *testing.T) {
	suite.Run(t, new(ImagesSuite))
}

func (s *ImagesSuite) TestListImages() {
	images, err := ListImages(ListImagesArgs{
		GetClientArgs: client.GetClientArgs{
			Host: test.GetDockerHost(),
		},
		Filters: []filters.KeyValuePair{
			{
				Key:   "reference",
				Value: "test",
			},
		},
	})
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(images))
}
