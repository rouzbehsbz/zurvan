package zurvan

import (
	"reflect"
	"sync"
)

type registry struct {
	registry map[reflect.Type]int
	counter  int

	mu sync.RWMutex
}

func newRegistry() *registry {
	return &registry{
		registry: make(map[reflect.Type]int),
		counter:  0,
		mu:       sync.RWMutex{},
	}
}

func (r *registry) dataId(dataType reflect.Type) int {
	r.mu.RLock()
	dataId, ok := r.registry[dataType]
	r.mu.RUnlock()

	if !ok {
		r.mu.Lock()

		dataId, ok = r.registry[dataType]
		if !ok {
			dataId = r.counter
			r.counter += 1
			r.registry[dataType] = dataId
		}

		r.mu.Unlock()
	}

	return dataId
}

func (r *registry) dataIdOf(data any) int {
	dataType := reflect.TypeOf(data)

	return r.dataId(dataType)
}

func dataIdFor[T any](registry *registry) int {
	dataType := reflect.TypeFor[T]()

	return registry.dataId(dataType)
}
