package services

import (
	"github.com/usename-Poezd/light-notifier/internal/repositories"
)

type NotificationService struct {
	repo repositories.Notification
}


func NewNotificationService(repo repositories.Notification) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}

func (s *NotificationService) Status() bool {
	return s.repo.Status()
}

func (s *NotificationService) On() bool {
	return s.repo.On()
}

func (s *NotificationService) Off() bool {
	return s.repo.Off()
}