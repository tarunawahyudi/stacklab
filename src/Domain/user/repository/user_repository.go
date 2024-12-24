package repository

import "stacklab/src/Domain/user/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindById(id string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(id string) error
}
