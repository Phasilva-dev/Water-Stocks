package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	//"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"devtest"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	//app := NewApp()
	simulador := &devtest.Simulador{}

	// Inicializa o app Wails
	err := wails.Run(&options.App{
		Title:     "Simulação Estocástica",
		Width:     800,
		Height:    600,
		Assets:    assets,
		Bind:      []interface{}{simulador}, // Expõe o simulador para o frontend
	})

	if err != nil {
		//log.Fatal(err)
		println("Error:", err.Error())
	}

	/*// Create application with options
	err := wails.Run(&options.App{
		Title:  "app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}*/
}
