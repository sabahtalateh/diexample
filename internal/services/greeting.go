package services

import (
	"context"
	"fmt"

	"github.com/sabahtalateh/di"
	"github.com/sabahtalateh/diexample/internal/repos"
)

type UsersRepo interface {
	GetUserLocale(context.Context, string) (string, error)
}

type GreetingsRepo interface {
	GetGreetingForLocale(context.Context, string) (string, error)
}

type GreetingService struct {
	users     UsersRepo
	greetings GreetingsRepo
}

func NewGreetingService(ur UsersRepo, gr GreetingsRepo) *GreetingService {
	return &GreetingService{users: ur, greetings: gr}
}

func SetupGreetingService(c *di.Container) error {
	return di.Setup[*GreetingService](c,
		di.Init(func(c *di.Container) *GreetingService {
			return NewGreetingService(
				di.Get[*repos.UsersRepo](c, di.Name("UsersRepo")),
				di.Get[*repos.GreetingsRepo](c, di.Name("GreetingsRepo")),
			)
		}),
	)
}

func (g *GreetingService) GetGreeting(ctx context.Context, name string) (string, error) {
	loc, err := g.users.GetUserLocale(ctx, name)
	if err != nil {
		return "", err
	}

	greeting, err := g.greetings.GetGreetingForLocale(ctx, loc)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s!", greeting, name), nil
}
