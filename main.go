package main

import (
	"context"
	"embed"
	"plainslate/backend/lib"
	"plainslate/backend/usecase"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"go.uber.org/zap"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	lib.InitLogger()
	defer lib.CloseLogger()

	// Create an instance of the app structure
	ctx := context.Background()
	c, err := lib.NewConfig()
	if err != nil {
		lib.LogFatal(ctx, "failed to load config",
			zap.Error(err),
		)
	}
	s := lib.NewSearcher()
	u := usecase.NewUsecase(c, s)
	app := NewApp(u)

	// Create application with options
	err = wails.Run(&options.App{
		Title: "PlainSlate",
		// -1 set the window size to the current screen resolution, not exactly the entire screen
		Width:            -1,
		Height:           -1,
		WindowStartState: options.Maximised,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			app.Usecase,
		},
	})

	if err != nil {
		lib.LogFatal(ctx, "failed to run app",
			zap.Error(err),
		)
	} else {
		lib.LogInfo(ctx, "app running")
	}
}
