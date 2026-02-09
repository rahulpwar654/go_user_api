package service

import (
	"gin-user-api/internal/model"
	"gin-user-api/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUser(id int64) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) ListUsersPaged(limit, offset int) ([]model.User, int64, error) {
	return s.repo.ListPaged(limit, offset)
}

func (s *UserService) CreateUser(u *model.User) error {
	return s.repo.Create(u)
}

func (s *UserService) UpdateUser(u *model.User) error {
	return s.repo.Update(u)
}

func (s *UserService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}
