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
	u := usecase.NewUsecase()
	app := NewApp(u)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "plainslate",
		Width:  1024,
		Height: 768,
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
