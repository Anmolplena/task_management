package user

import (
	"GoWithMongo/datstore/entity"
	"GoWithMongo/datstore/repository"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	CreateUser(c *gin.Context, user entity.User) error
	DeleteUser(c *gin.Context, username string) error
	GetUserByUsername(c *gin.Context, username string) (*UserDTO, error)
}

type userService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		UserRepo: userRepo,
	}
}

func (s *userService) CreateUser(c *gin.Context, user entity.User) error {
	return s.UserRepo.CreateUser(user)

}

func (s *userService) DeleteUser(c *gin.Context, username string) error {
	return s.UserRepo.DeleteUser(username)

}

func (s *userService) GetUserByUsername(c *gin.Context, username string) (*UserDTO, error) {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	userDTO := UserDTO{
		Username: user.Username,
		Email:    user.Email,
		// Set other fields as needed
	}

	return &userDTO, nil
}
