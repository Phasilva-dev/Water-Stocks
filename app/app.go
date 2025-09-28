package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"simulation/cmd/simulation"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RunSimulation(size, day, toiletType, showerType int, filename string) {
	simulation.RunSimulation(size, day, toiletType, showerType, filename)
}

func (a *App) SelectFile() (string, error) {
    selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
        Title: "Selecione um arquivo de análise",
        Filters: []runtime.FileFilter{
            {
                DisplayName: "Arquivos CSV (*.csv)",
                Pattern:     "*.csv",
            },
        },
    })
    if err != nil {
        return "", err
    }
    // Se o usuário cancelar, a string será vazia, o que é o comportamento desejado.
    return selection, nil
}
