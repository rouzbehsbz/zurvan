package zurvan

import (
	"reflect"
)

type events struct {
	columns map[int]column

	registry *registry
}

func newEvents(registry *registry) *events {
	return &events{
		columns:  make(map[int]column),
		registry: registry,
	}
}

func (e *events) emit(event any) {
	eventId := e.registry.dataIdOf(event)

	column, ok := e.columns[eventId]
	if !ok {
		column = newVector(reflect.TypeOf(event))
		e.columns[eventId] = column
	}

	column.push(event)
}

func (e *events) Clear() {
	for _, column := range e.columns {
		column.resize(0)
	}
}

func OnEvent[T any](w *World) []T {
	eventId := dataIdFor[T](w.events.registry)

	column, ok := w.events.columns[eventId]
	if !ok {
		return []T{}
	}

	slice := column.asSlice()

	return slice.([]T)
}
