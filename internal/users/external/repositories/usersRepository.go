package users

import (
	"context"
	"database/sql"
	"time"

	users "github.com/read-manager/read-manager-main-api/internal/users/domain"
)

type IUsersRepository interface {
	Insert(user users.User) error
}

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *usersRepository {
	return &usersRepository{
		db,
	}
}

func (ur *usersRepository) Insert(user users.User) error {
	query := `
        INSERT INTO users (
            id,
            name,
            email,
            password,
            nickname,
            read_points
        )
        VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6
        )
        RETURNING
            created_at
    `
	args := []any{
		user.Id,
		user.Name,
		user.Email,
		user.Password.GetHash(),
		user.Nickname,
		0,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := ur.db.QueryRowContext(ctx, query, args...).Scan(&user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
