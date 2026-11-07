package zurvan

import (
	"reflect"
)

type events struct {
	events  *queue[any]
	columns map[int]column

	registry *registry
}

func newEvents(registry *registry) *events {
	return &events{
		events:   newQueue[any](),
		columns:  make(map[int]column),
		registry: registry,
	}
}

func (e *events) emit(events []any) {
	e.events.push(events)
}

func (e *events) apply() {
	events := e.events.drain()

	for _, event := range events {
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

// OnEvent returns the current frame's snapshot of events of type T.
// Reading events does not consume them. Events emitted during the frame become
// readable next frame. Treat the slice as read-only and do not retain it beyond
// the frame. EndStage also sees events published during shutdown.
func OnEvent[T any](w *World) []T {
	eventId := dataIdFor[T](w.events.registry)

	column, ok := w.events.columns[eventId]
	if !ok {
		return []T{}
	}

	slice := column.asSlice()

	return slice.([]T)
}
