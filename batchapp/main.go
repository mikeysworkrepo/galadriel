package main

import (
	"batchapp/backend"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := backend.NewApp()

	err := wails.Run(&options.App{
		Title:     "BatchApp",
		Width:     1024,
		Height:    768,
		MinWidth:  800,
		MinHeight: 600,
		Windows:   &windows.Options{},
		Bind: []interface{}{
			app,
		},
		Assets: assets,
	})

	if err != nil {
		panic(err)
	}
}
