package zurvan

import "sync"

type queue[T any] struct {
	pending []T
	mu      sync.Mutex
}

func newQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) push(values []T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.pending = append(q.pending, values...)
}

func (q *queue[T]) drain() []T {
	q.mu.Lock()
	defer q.mu.Unlock()

	buffer := q.pending
	q.pending = nil

	return buffer
}
