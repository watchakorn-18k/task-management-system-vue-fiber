package services

import (
	"errors"
	"task_management_system/src/domain/entities"
	"task_management_system/src/domain/repositories"
	"task_management_system/src/infrastructure/generals"
	"task_management_system/src/middlewares"

	"github.com/google/uuid"
)

type usersService struct {
	UsersRepo repositories.IUsersRepository
}

type IUsersService interface {
	RegisterUsers(data *entities.UserModel) error
	Login(data *entities.UserModel) (*string, error)
	GetProfile(userID string) (*entities.User, error)
}

func NewUsersService(usersRepo repositories.IUsersRepository) IUsersService {
	return &usersService{
		UsersRepo: usersRepo,
	}
}

func (s *usersService) RegisterUsers(data *entities.UserModel) error {
	if data.Username == "" {
		return errors.New("username is required")
	}
	if data.Password == "" || data.PasswordConfirm == "" {
		return errors.New("password is required")
	}
	if data.Password != data.PasswordConfirm {
		return errors.New("password not match")
	}

	userData, err := s.UsersRepo.GetUser(data.Username)
	if err != nil {
		return err
	}
	if userData != nil {
		return errors.New("username already exist")
	}
	hashPassword, err := generals.HashPassword(data.Password)
	if err != nil {
		return err
	}
	userID := uuid.NewString()
	userNew := entities.User{
		UserID:   userID,
		Username: data.Username,
		Password: hashPassword,
	}

	if err := s.UsersRepo.RegisterUsers(&userNew); err != nil {
		return err
	}
	return nil
}

func (s *usersService) Login(data *entities.UserModel) (*string, error) {
	if data.Username == "" || data.Password == "" {
		return nil, errors.New("username or password is required")
	}
	userData, err := s.UsersRepo.GetUser(data.Username)
	if err != nil {
		return nil, err
	}
	if userData == nil {
		return nil, errors.New("username not found")
	}
	if status := generals.CheckPassword(userData.Password, data.Password); !status {
		return nil, errors.New("password not match")
	}
	jwt, err := middlewares.GenerateJWTToken(userData.UserID)
	if err != nil {
		return nil, err
	}
	userUpdate := &entities.User{
		JWT: *jwt.Token,
	}
	if err := s.UsersRepo.UpdateUsers(userData.UserID, userUpdate); err != nil {
		return nil, err
	}
	return jwt.Token, nil
}

func (s *usersService) GetProfile(userID string) (*entities.User, error) {
	return s.UsersRepo.GetUser(userID)
}
