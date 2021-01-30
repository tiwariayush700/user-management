package repositoryImpl

import "github.com/tiwariayush700/user-management/repository"

type userRepositoryImpl struct {
	repositoryImpl //overrides basic CRUD repo
}

func NewUserRepositoryImpl(repositoryImpl repositoryImpl) repository.Repository {
	return &userRepositoryImpl{repositoryImpl: repositoryImpl}
}
