package domain

type User struct {
	Id   int    `json:"id" db:"id"`
	ChatId int64 `json:"chat_id" db:"chat_id"`
}