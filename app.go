package main

import (
	"byto/internal/builder"
	"byto/internal/command"
	"byto/internal/domain"
	"byto/internal/queue"
	"context"
	"fmt"
)

type App struct {
	ctx      context.Context
	queue    *queue.Queue
	settings *domain.Setting
}

func NewApp() *App {
	return &App{
		queue:    queue.NewQueue(),
		settings: domain.NewSetting(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AddToQueue(url string) {
	a.queue.Add(&domain.Media{
		URL:    url,
		Status: domain.Pending,
	})
}

func (a *App) RemoveFromQueue(id string) error {
	return a.queue.Remove(id)
}

func (a *App) GetQueue() []*domain.Media {
	return a.queue.GetAll()
}

func (a *App) StartDownloads() {
	if a.settings == nil {
		// Default settings if nil
		a.settings = &domain.Setting{
			Quality:           domain.Quality1080p,
			ParallelDownloads: 3,
			DownloadPath:      "./downloads",
		}
	}

	queueItems := a.queue.GetAll()
	semaphore := make(chan struct{}, a.settings.ParallelDownloads)

	for _, media := range queueItems {
		if media.Status == domain.Pending {
			go func(m *domain.Media) {
				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				m.SetStatus(domain.InProgress)

				// Initialize builder
				b := &builder.YTDLPBuilder{}
				b.NewYTDLPBuilder().
					URL(m.URL).
					Quality(a.settings.Quality).
					DownloadPath(a.settings.DownloadPath)

				cmd := &command.DownloadCommand{
					Builder: b,
				}

				if err := cmd.Execute(m); err != nil {
					m.SetStatus(domain.Failed)
					m.AppendLog(fmt.Sprintf("Download failed: %v", err))
				}
			}(media)
		}
	}
}
