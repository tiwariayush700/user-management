package repository

import (
	"context"
	"github.com/tiwariayush700/user-management/models"
)

type UserRepository interface {
	Repository
	GetUserByEmailAndPassword(ctx context.Context, email, password string) (*models.User, error)
	UpdateUserRole(ctx context.Context, userId uint, role string) error
}
