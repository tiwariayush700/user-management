package serviceImpl

import (
	"context"
	"github.com/tiwariayush700/user-management/models"
	"github.com/tiwariayush700/user-management/repository"
	"github.com/tiwariayush700/user-management/services"
	"github.com/tiwariayush700/user-management/utils"
)

type userServiceImpl struct {
	repository repository.UserRepository
}

func (u *userServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	user.Password = utils.GetMd5(user.Password)

	err := u.repository.Create(ctx, user)
	return err
}

func NewUserServiceImpl(repository repository.UserRepository) services.UserService {
	return &userServiceImpl{repository: repository}
}




