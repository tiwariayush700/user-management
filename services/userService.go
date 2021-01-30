package services

import (
	"context"
	"github.com/tiwariayush700/user-management/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
}
