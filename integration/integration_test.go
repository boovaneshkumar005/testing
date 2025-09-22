package integration

import (
	"gotesting/config"
	"gotesting/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryIntegration(t *testing.T) {
	config.InitDB()
	defer config.CloseDB()

	user := repository.User{ID: 500, Name: "IntegrationUser"}
	err := repository.AddUser(user)
	assert.Nil(t, err)

	fetchedUser, err := repository.GetUserByID(500)
	assert.Nil(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, "IntegrationUser", fetchedUser.Name)
}
