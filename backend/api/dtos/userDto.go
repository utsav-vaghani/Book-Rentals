package dtos

import (
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
)

type UserDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func MapUserToUserDto(user *models.User) *UserDto {
	return &UserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}