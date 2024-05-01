package system

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/sabahtalateh/di"
	"github.com/sabahtalateh/diexample/internal/config"
	"github.com/sabahtalateh/diexample/internal/setup/stages"
)

type DB struct {
	dsn  string
	sqlx *sqlx.DB
}

func NewDB(dsn string) *DB {
	return &DB{dsn: dsn}
}

func (db *DB) Connect(ctx context.Context) error {
	var err error
	db.sqlx, err = sqlx.ConnectContext(ctx, "postgres", db.dsn)
	if err != nil {
		err = fmt.Errorf("starting DB: %w", err)
	}
	return err
}

func (db *DB) SqlX() *sqlx.DB {
	return db.sqlx
}

func SetupDB(c *di.Container) error {
	return di.Setup[*DB](c,
		di.Init(func(c *di.Container) *DB {
			return NewDB(di.Get[*config.Config](c).DB.DSN)
		}),
		di.Stage(stages.StartSystem, func(ctx context.Context, db *DB) error {
			return db.Connect(ctx)
		}),
		di.Stage(stages.StopSystem, func(_ context.Context, db *DB) error {
			return nil
		}),
	)
}
