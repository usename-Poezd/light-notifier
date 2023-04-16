package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/usename-Poezd/light-notifier/internal/domain"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		db,
	}
}

func (r *UsersRepo) GetAll() ([]domain.User, error) {

	s := []domain.User{}
	err := r.db.Select(&s, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return s, nil
}


func (r *UsersRepo) Create(u *domain.User) (*domain.User, error) {
	tx := r.db.MustBegin()
	rows, err := tx.Exec(`
		INSERT OR IGNORE INTO users (
			chat_id
		) VALUES (
			?
		) RETURNING id`,
		u.ChatId,
	)

	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.Id = int(id)

	log.Println(u)
	return u, nil
}
