package app

import (
	"context"
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"gorm.io/gorm"

	"github.com/ksckaan1/apiredator/internal/domain/core/port"
	"github.com/ksckaan1/apiredator/internal/domain/core/service"
	"github.com/ksckaan1/apiredator/internal/infrastructure/repository"
)

type App struct {
	appService *service.AppService
	logger     port.Logger
	assets     embed.FS
	db         *gorm.DB
}

func New(lg port.Logger, assets embed.FS, db *gorm.DB) *App {
	return &App{
		logger: lg,
		assets: assets,
		db:     db,
	}
}

func (a *App) Init(ctx context.Context) error {
	repo, err := repository.NewRepository(a.db)
	if err != nil {
		return fmt.Errorf("init app: %w", err)
	}

	err = repo.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("repo: migrate: %w", err)
	}

	a.appService = service.NewAppService(a.logger, repo)

	return nil
}

func (a *App) Run(ctx context.Context) error {
	err := wails.Run(&options.App{
		Title:  "APIredator",
		Width:  1600,
		Height: 1000,
		AssetServer: &assetserver.Options{
			Assets: a.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.appService.Startup,
		MinWidth:         800,
		MinHeight:        800,
		Bind: []any{
			a.appService,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            true,
				UseToolbar:                 true,
				HideToolbarSeparator:       true,
			},
			About: &mac.AboutInfo{
				Title:   "APIredator",
				Message: "API load testing tool",
			},
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})
	if err != nil {
		return fmt.Errorf("wails run: %w", err)
	}
	return nil
}
