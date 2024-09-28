package main

import (
	"context"
	"embed"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ksckaan1/apiredator/cmd/gui/app"
	"github.com/ksckaan1/apiredator/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	db, err := gorm.Open(sqlite.Open(getDBFilePath()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	a := app.New(lg, assets, db)

	err = a.Init(ctx)
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

// TODO: Add db file path for other operating systems
func getDBFilePath() string {
	var dbPath string
	switch runtime.GOOS {
	case "darwin":
		dbPath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "APIredator", "main.db")
	case "linux":
		dbPath = filepath.Join(os.Getenv("HOME"), ".config", "APIredator", "main.db")
	default:
		panic("unsupported operating system")
	}

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(dbPath), os.ModePerm)
		_, err = os.Create(dbPath)
		if err != nil {
			panic(err)
		}
	}

	return dbPath
}
