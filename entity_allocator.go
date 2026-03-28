package zurvan

import (
	"sync/atomic"
)

type entityAllocator struct {
	counter atomic.Uint32
}

func newEntityAllocator() *entityAllocator {
	return &entityAllocator{
		counter: atomic.Uint32{},
	}
}

func (e *entityAllocator) create() Entity {
	entityId := e.counter.Add(1)

	return newEntity(entityId)
}
