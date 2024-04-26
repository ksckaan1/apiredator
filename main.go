package main

import (
	"context"
	"embed"
	"os"

	"github.com/ksckaan1/apiredator/cmd/gui/app"
	"github.com/ksckaan1/apiredator/pkg/logger"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	lg := logger.NewZerolog().
		WithWriter(os.Stdout).
		WithLevel(logger.TraceLevel).
		WithTime().
		WithLayer("apiredator")

	ctx := context.Background()

	a := app.New(lg, assets)

	err := a.Init(ctx)
	if err != nil {
		lg.Fatal("error when app init",
			"error", err,
		)
	}

	err = a.Run(ctx)
	if err != nil {
		lg.Fatal("error when app run",
			"error", err,
		)
	}
}
