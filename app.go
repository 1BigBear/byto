package main

import (
	"byto/internal/domain"
	"byto/internal/queue"
	"context"
	"fmt"
)

// App struct
// App struct
type App struct {
	ctx   context.Context
	queue *queue.Queue
}

// NewApp creates a new App application struct
// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		queue: queue.NewQueue(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) AddToQueue(url string) {
	a.queue.Add(&domain.Media{
		URL:    url,
		Status: domain.Pending,
	})
}

func (a *App) RemoveFromQueue(index int) error {
	return a.queue.Remove(index)
}

func (a *App) GetQueue() []*domain.Media {
	return a.queue.GetAll()
}
