package client

import (
	"testing"

	"github.com/risersh/builder/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client *Client
}

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (s *ClientSuite) SetupTest() {
	test.Setup()
	var err error
	s.client, err = GetClient(GetClientArgs{
		Host: test.GetDockerHost(),
	})
	assert.NoError(s.T(), err)
}

func (s *ClientSuite) TestPing() {
	assert.True(s.T(), s.client.Ping())
}
