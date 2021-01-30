package repositoryImpl

import (
	"github.com/tiwariayush700/user-management/repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	repositoryImpl //overrides basic CRUD repo
}

func NewUserRepositoryImpl(db *gorm.DB) repository.Repository {
	repoImpl := repositoryImpl{
		DB: db,
	}
	return &userRepositoryImpl{repoImpl}
}
