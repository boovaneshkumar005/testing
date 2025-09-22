package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// UserServiceMock is a mock of repository functions
type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) GetUserByID(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func TestGetUserWithMock(t *testing.T) {
	mockRepo := new(UserServiceMock)
	mockUser := &User{ID: 1, Name: "Alice"}

	// Setup expectation
	mockRepo.On("GetUserByID", 1).Return(mockUser, nil)

	user, err := mockRepo.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", user.Name)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
