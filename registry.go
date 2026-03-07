package zurvan

import "reflect"

type registry struct {
	registry map[reflect.Type]int
	counter  int
}

func newRegistry() *registry {
	return &registry{
		registry: make(map[reflect.Type]int),
		counter:  0,
	}
}

func (r *registry) dataId(dataType reflect.Type) int {
	dataId, ok := r.registry[dataType]
	if !ok {
		dataId = r.counter
		r.counter += 1
		r.registry[dataType] = dataId
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
