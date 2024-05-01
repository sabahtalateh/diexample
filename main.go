package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sabahtalateh/di"
	"github.com/sabahtalateh/diexample/internal/services"
	"github.com/sabahtalateh/diexample/internal/setup"
	"github.com/sabahtalateh/diexample/internal/setup/stages"
)

func main() {
	if len(os.Args) <= 1 {
		println("provide args")
		return
	}

	c, err := setup.Container()
	check(err)

	err = c.Init()
	check(err)

	err = c.ExecStage(context.Background(), stages.StartSystem)
	check(err)

	gs, err := di.GetE[*services.GreetingService](c)
	check(err)

	for _, n := range os.Args[1:] {
		g, err := gs.GetGreeting(context.Background(), n)
		check(err)
		fmt.Println(g)
	}

	err = c.ExecStage(context.Background(), stages.StopSystem)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
