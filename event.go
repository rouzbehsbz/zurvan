package zurvan

import (
	"reflect"
)

const eventBufferSize int = 1024

type events struct {
	events  chan any
	columns map[int]column

	registry *registry
}

func newEvents(registry *registry) *events {
	return &events{
		events:   make(chan any, eventBufferSize),
		columns:  make(map[int]column),
		registry: registry,
	}
}

func (e *events) emit(event any) {
	e.events <- event
}

func (e *events) apply() {
	for len(e.events) > 0 {
		event := <-e.events

		eventId := e.registry.dataIdOf(event)
		column, ok := e.columns[eventId]
		if !ok {
			column = newVector(reflect.TypeOf(event))
			e.columns[eventId] = column
		}

		column.push(event)
	}
}

func (e *events) clear() {
	for _, column := range e.columns {
		column.resize(0)
	}
}

// Returns a slice of events of type T that were emitted during the current frame
func OnEvent[T any](w *World) []T {
	eventId := dataIdFor[T](w.events.registry)

	column, ok := w.events.columns[eventId]
	if !ok {
		return []T{}
	}

	slice := column.asSlice()

	return slice.([]T)
}
