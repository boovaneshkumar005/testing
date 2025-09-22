package integration

import (
	"gotesting/config"
	"gotesting/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserIntegrationSuite struct {
	suite.Suite
}

func (s *UserIntegrationSuite) SetupSuite() {
	config.InitDB()
}

func (s *UserIntegrationSuite) TearDownSuite() {
	config.CloseDB()
}

func (s *UserIntegrationSuite) TestAddAndGetUser() {
	user := repository.User{ID: 100, Name: "IntegrationUser"}
	err := repository.AddUser(user)
	assert.NoError(s.T(), err)

	fetched, err := repository.GetUserByID(100)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "IntegrationUser", fetched.Name)
}

func TestUserIntegrationSuite(t *testing.T) {
	suite.Run(t, new(UserIntegrationSuite))
}
