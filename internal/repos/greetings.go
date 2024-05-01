package repos

import (
	"context"
	"fmt"

	"github.com/sabahtalateh/di"
	"github.com/sabahtalateh/diexample/internal/system"
)

type GreetingsRepo struct {
	db *system.DB
}

func NewGreetingsRepo(db *system.DB) *GreetingsRepo {
	return &GreetingsRepo{db: db}
}

func SetupGreetingsRepo(c *di.Container) error {
	return di.Setup[*GreetingsRepo](c,
		di.Name("GreetingsRepo"),
		di.Init(func(c *di.Container) *GreetingsRepo {
			return NewGreetingsRepo(di.Get[*system.DB](c))
		}),
	)
}

func (r *GreetingsRepo) GetGreetingForLocale(ctx context.Context, locale string) (string, error) {
	// imitation of db work
	var greeting string
	if locale == "ru" {
		greeting = "Privet"
	}
	if locale == "en" {
		greeting = "Hello"
	}
	if locale == "de" {
		greeting = "Guten Tag"
	}

	var g string
	err := r.db.SqlX().QueryRowxContext(ctx, fmt.Sprintf("SELECT '%s'", greeting)).Scan(&g)
	if err != nil {
		return "", err
	}
	return g, nil
}
