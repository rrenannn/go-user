package user

import (
	"context"
	"database/sql"

	db "github.com/rrenannn/go-user/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserById(ctx context.Context, id int64) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	ResetPassword(ctx context.Context, arg db.ResetPasswordParams) error
}

type Repository struct {
	Conn    *sql.DB
	Queries *db.Queries
}

func NewRepository(conn *sql.DB, queries *db.Queries) *Repository {
	return &Repository{
		Conn:    conn,
		Queries: queries,
	}
}

func (b *Repository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return b.Queries.CreateUser(ctx, arg)
}

func (b *Repository) GetUserById(ctx context.Context, id int64) (db.User, error) {
	return b.Queries.GetUserById(ctx, id)
}

func (b *Repository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return b.Queries.GetUserByEmail(ctx, email)
}

func (b *Repository) ResetPassword(ctx context.Context, arg db.ResetPasswordParams) error {
	return b.Queries.ResetPassword(ctx, arg)
}
