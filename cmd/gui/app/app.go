package app

import (
	"context"
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	"github.com/ksckaan1/apiredator/internal/core/port"
	"github.com/ksckaan1/apiredator/internal/core/service"
)

type App struct {
	appService *service.AppService
	logger     port.Logger
	assets     embed.FS
}

func New(lg port.Logger, assets embed.FS) *App {
	return &App{
		logger: lg,
		assets: assets,
	}
}

func (a *App) Init(ctx context.Context) error {
	a.appService = service.NewAppService(a.logger)
	return nil
}

func (a *App) Run(ctx context.Context) error {
	err := wails.Run(&options.App{
		Title:  "APIredator",
		Width:  1024,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: a.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.appService.Startup,
		MinWidth:         400,
		MinHeight:        400,
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
			OpenInspectorOnStartup: true,
		},
	})
	if err != nil {
		return fmt.Errorf("wails run: %w", err)
	}
	return nil
}
