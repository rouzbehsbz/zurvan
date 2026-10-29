package zurvan

import (
	"math"
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
	for {
		current := e.counter.Load()

		if current == math.MaxUint32 {
			panic("zurvan: entity IDs exhausted")
		}

		if e.counter.CompareAndSwap(current, current+1) {
			return newEntity(current + 1)
		}
	}
}
