package queue

import (
	"byto/internal/domain"
	"testing"
)

func TestQueue_Add(t *testing.T) {
	q := NewQueue()
	media := &domain.Media{URL: "http://example.com/video"}

	q.Add(media)

	items := q.GetAll()
	if len(items) != 1 {
		t.Errorf("Expected 1 item in queue, got %d", len(items))
	}
	if items[0].URL != "http://example.com/video" {
		t.Errorf("Expected URL to be 'http://example.com/video', got '%s'", items[0].URL)
	}
}

func TestQueue_Remove(t *testing.T) {
	q := NewQueue()
	media1 := &domain.Media{URL: "http://example.com/video1"}
	media2 := &domain.Media{URL: "http://example.com/video2"}

	q.Add(media1)
	q.Add(media2)

	err := q.Remove(0)
	if err != nil {
		t.Errorf("Unexpected error calling Remove: %v", err)
	}

	items := q.GetAll()
	if len(items) != 1 {
		t.Errorf("Expected 1 item in queue, got %d", len(items))
	}
	if items[0].URL != "http://example.com/video2" {
		t.Errorf("Expected remaining item to be 'http://example.com/video2', got '%s'", items[0].URL)
	}

	err = q.Remove(10) // Out of bounds
	if err == nil {
		t.Error("Expected error when removing out of bounds index, got nil")
	}
}
