package services

import (
	"github.com/usename-Poezd/light-notifier/internal/domain"
	"github.com/usename-Poezd/light-notifier/internal/repositories"
)

type UsersService struct {
	repo repositories.Users
}


func NewUsersService(repo repositories.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) GetAll() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *UsersService) Create(u *domain.User) (*domain.User, error) {
	return s.repo.Create(u)
}