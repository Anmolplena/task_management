package user

import "GoWithMongo/datstore/entity"

type UserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Type     string `json:"type"`
}

func NewUserDTO(user *entity.User) *UserDTO {
	return &UserDTO{
		Username: user.Username,
		Email:    user.Email,
		Type:     user.Type,
	}
}
