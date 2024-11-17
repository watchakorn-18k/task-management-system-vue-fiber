package services_test

import (
	"errors"
	"task_management_system/src/domain/entities"
	repositories_test "task_management_system/src/mocks/repositories"
	"task_management_system/src/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockUsersRepoUserSV = new(repositories_test.MockUserRepository)
	usersSV             = services.NewUsersService(mockUsersRepoUserSV)
)

func RefreshRepoUsersSV() {
	mockUsersRepoUserSV = new(repositories_test.MockUserRepository)
	usersSV = services.NewUsersService(mockUsersRepoUserSV)
}

func TestLogin(t *testing.T) {
	t.Run("Login username or password is required", func(t *testing.T) {
		RefreshRepoUsersSV()
		_, err := usersSV.Login(&entities.UserModel{})
		assert.Error(t, err)
	})
	t.Run("Login username username not found", func(t *testing.T) {
		RefreshRepoUsersSV()
		mockUsersRepoUserSV.On("GetUsername", mock.Anything).Return(nil, errors.New("username not found"))
		_, err := usersSV.Login(&entities.UserModel{Username: "test", Password: "test"})
		assert.Equal(t, err.Error(), "username not found")
	})
	t.Run("Login password not match", func(t *testing.T) {
		RefreshRepoUsersSV()
		mockUsersRepoUserSV.On("GetUsername", mock.Anything).Return(&entities.User{}, nil)
		_, err := usersSV.Login(&entities.UserModel{Username: "test", Password: "test"})
		assert.Error(t, err)
	})

}

func TestGetProfile(t *testing.T) {
	t.Run("Get profile Error GetUser", func(t *testing.T) {
		RefreshRepoUsersSV()
		mockUsersRepoUserSV.On("GetUser", mock.Anything).Return(entities.User{}, errors.New("user not found"))
		_, err := usersSV.GetProfile("test")
		assert.Equal(t, err.Error(), "user not found")
	})
	t.Run("Get profile success", func(t *testing.T) {
		RefreshRepoUsersSV()
		mockUsersRepoUserSV.On("GetUser", mock.Anything).Return(entities.User{Username: "test"}, nil)
		_, err := usersSV.GetProfile("test")
		assert.NoError(t, err)
	})
}
