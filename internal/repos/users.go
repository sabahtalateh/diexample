package repos

import (
	"context"
	"fmt"

	"github.com/sabahtalateh/di"
	"github.com/sabahtalateh/diexample/internal/system"
)

type UsersRepo struct {
	db *system.DB
}

func NewUsersRepo(db *system.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func SetupUsersRepo(c *di.Container) error {
	return di.Setup[*UsersRepo](c,
		di.Name("UsersRepo"),
		di.Init(func(c *di.Container) *UsersRepo {
			return NewUsersRepo(di.Get[*system.DB](c))
		}),
	)
}

func (r *UsersRepo) GetUserLocale(ctx context.Context, user string) (string, error) {
	// imitation of db work
	var locale string
	if user == "Ivan" {
		locale = "ru"
	}
	if user == "John" {
		locale = "en"
	}
	if user == "Hans" {
		locale = "de"
	}

	var g string
	err := r.db.SqlX().QueryRowxContext(ctx, fmt.Sprintf("SELECT '%s'", locale)).Scan(&g)
	if err != nil {
		return "", err
	}
	return g, nil
}
