package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/usename-Poezd/light-notifier/internal/domain"
)

type Repositories struct {
	Users Users
	Notification Notification
}

type Users interface {
	GetAll() ([]domain.User, error)
	Create(s *domain.User) (*domain.User, error)
}

type Notification interface {
	On() bool
	Off() bool
	Status() bool
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
		Notification: NewNotificationRepo(),
	}
}