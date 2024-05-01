package setup

import (
	"github.com/sabahtalateh/di"
	"github.com/sabahtalateh/diexample/internal/config"
	"github.com/sabahtalateh/diexample/internal/repos"
	"github.com/sabahtalateh/diexample/internal/services"
	"github.com/sabahtalateh/diexample/internal/system"
)

func Container() (*di.Container, error) {
	c := di.NewContainer()

	setupFns := []func(*di.Container) error{
		config.SetupConfig,
		system.SetupDB,
		repos.SetupGreetingsRepo,
		repos.SetupUsersRepo,
		services.SetupGreetingService,
	}

	for _, fn := range setupFns {
		if err := fn(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}
