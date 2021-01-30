package services

import (
	"context"
	"github.com/tiwariayush700/user-management/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID uint) (*models.UserResponse, error)
	LoginUser(ctx context.Context, email, password string) (*models.UserResponse, error)
}
