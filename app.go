package main

import (
	"context"
	"plainslate/backend/usecase"
)

// App struct
type App struct {
	ctx     context.Context
	Usecase *usecase.Usecase
}

// NewApp creates a new App application struct
func NewApp(u *usecase.Usecase) *App {
	return &App{
		Usecase: u,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
