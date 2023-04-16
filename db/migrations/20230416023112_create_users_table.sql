-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
	"id" integer PRIMARY KEY AUTOINCREMENT,
	"chat_id" bigint NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users"
-- +goose StatementEnd
