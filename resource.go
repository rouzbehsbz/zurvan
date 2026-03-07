package zurvan

import "reflect"

type resources struct {
	storage  map[int]any
	registry *registry
}

func newResources() *resources {
	return &resources{
		storage:  make(map[int]any),
		registry: newRegistry(),
	}
}

func (r *resources) addResource(resource any) {
	resourceType := reflect.TypeOf(resource)
	resourceId := r.registry.dataId(resourceType)

	r.storage[resourceId] = resource
}

func Resource[T any](w *World) (T, bool) {
	resourceId := dataIdFor[T](w.resources.registry)

	resource, ok := w.resources.storage[resourceId]
	if !ok {
		var defaultVal T
		return defaultVal, false
	}

	return resource.(T), true
}
