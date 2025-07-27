package main

import (
	"embed"
	"log"
	"plainslate/backend/lib"
	"plainslate/backend/usecase"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	lib.InitLogger()
	defer lib.CloseLogFile()

	// Create an instance of the app structure
	c, err := lib.NewConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	u := usecase.NewUsecase(c)
	app := NewApp(u)

	// Create application with options
	err = wails.Run(&options.App{
		Title: "plainslate",
		// -1 set the window size to the current screen resolution, not exactly the entire screen
		Width:  -1,
		Height: -1,
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
		log.Fatal(err)
	} else {
		log.Println("[INFO] plainslate running")
	}
}
