package queue

import (
	"byto/internal/domain"
	"errors"
	"sync"
)

type Queue struct {
	items []*domain.Media
	mu    sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]*domain.Media, 0),
	}
}

func (q *Queue) Add(media *domain.Media) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, media)
}

func (q *Queue) Remove(index int) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if index < 0 || index >= len(q.items) {
		return errors.New("index out of bounds")
	}
	q.items = append(q.items[:index], q.items[index+1:]...)
	return nil
}

func (q *Queue) GetAll() []*domain.Media {
	q.mu.Lock()
	defer q.mu.Unlock()
	itemsCopy := make([]*domain.Media, len(q.items))
	copy(itemsCopy, q.items)
	return itemsCopy
}
