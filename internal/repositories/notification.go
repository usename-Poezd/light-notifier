package repositories

type NotificationRepo struct {
	status bool
}

func NewNotificationRepo() *NotificationRepo {
	return &NotificationRepo{
		status: true,
	}
}

func (r *NotificationRepo) On() bool {
	r.status = true
	return r.status
}
func (r *NotificationRepo) Off() bool {
	r.status = false
	return r.status
}
func (r *NotificationRepo) Status() bool {
	return r.status
}
