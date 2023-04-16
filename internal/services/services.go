package services

import (
	"github.com/usename-Poezd/light-notifier/internal/domain"
	"github.com/usename-Poezd/light-notifier/internal/repositories"
)

type Keenetic interface {
	Check() error
}

type Users interface {
	GetAll() ([]domain.User, error)
	Create(*domain.User) (*domain.User, error)
}

type Notification interface {
	On() bool
	Off() bool
	Status() bool
}

type Services struct {
	Keenetic Keenetic
	Users Users
	Notification Notification
}

type Deps struct {
	KeeneticDnsDomain string
}

func NewServices(repos *repositories.Repositories, deps *Deps) *Services {
	return &Services{
		Keenetic: NewKeeneticService(deps.KeeneticDnsDomain),
		Users: NewUsersService(repos.Users),
		Notification: NewNotificationService(repos.Notification),
	}
}