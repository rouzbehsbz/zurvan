package zurvan

import (
	"reflect"
	"slices"
	"sort"
)

type entityLocation struct {
	maskKey string
	row     int
}

func newEntityLocation(maskKey string, row int) entityLocation {
	return entityLocation{
		maskKey: maskKey,
		row:     row,
	}
}

type componentEntry struct {
	id       int
	elemType reflect.Type
}

func newComponentEntry(id int, elemType reflect.Type) componentEntry {
	return componentEntry{
		id:       id,
		elemType: elemType,
	}
}

func sortComponentEntries(entries []componentEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].id < entries[j].id
	})
}

type archetypeAllocator struct {
	archetypes map[string]*archetype
	locations  map[Entity]entityLocation

	registry *registry
}

func newArchetypeAllocator(registry *registry) *archetypeAllocator {
	return &archetypeAllocator{
		archetypes: make(map[string]*archetype),
		locations:  make(map[Entity]entityLocation),
		registry:   registry,
	}
}

func (a *archetypeAllocator) addComponents(entity Entity, components ...any) {
	if !entity.isAlive() {
		return
	}

	location, exists := a.locations[entity]
	if exists {
		source := a.archetypes[location.maskKey]

		var mask mask

		for _, c := range components {
			id := a.registry.dataIdOf(c)
			mask.set(id)
		}

		if maskHasComponents(source.mask, mask) {
			a.setComponents(source, location.row, components)
			return
		}

		newMask := maskUnion(source.mask, mask)
		newMaskKey := newMask.key()
		target, ok := a.archetypes[newMaskKey]
		if !ok {
			entries := make([]componentEntry, 0, len(components)+len(source.columns))

			for _, c := range source.columns {
				entries = append(entries, newComponentEntry(c.componentId, c.elemType))
			}
			for _, c := range components {
				id := a.registry.dataIdOf(c)
				entries = append(entries, newComponentEntry(id, reflect.TypeOf(c)))
			}

			sortComponentEntries(entries)

			target = newArchetype(entries)
			a.archetypes[newMaskKey] = target
		}

		newRow := target.addEntity(entity)

		source.moveComponents(location.row, newRow, target)
		a.setComponents(target, newRow, components)
		swappedEntity, swappedRow := source.removeEntity(location.row)

		if swappedRow != -1 {
			swLocation := a.locations[swappedEntity]
			swLocation.row = swappedRow
			a.locations[swappedEntity] = swLocation
		}

		a.locations[entity] = newEntityLocation(newMaskKey, newRow)
		return
	}

	var mask mask

	entries := make([]componentEntry, 0, len(components))

	for _, c := range components {
		id := a.registry.dataIdOf(c)
		mask.set(id)
		entries = append(entries, newComponentEntry(id, reflect.TypeOf(c)))
	}

	sortComponentEntries(entries)

	maskKey := mask.key()
	targetArchetype, ok := a.archetypes[maskKey]
	if !ok {
		targetArchetype = newArchetype(entries)
		a.archetypes[maskKey] = targetArchetype
	}

	row := targetArchetype.addEntity(entity)

	a.setComponents(targetArchetype, row, components)
	a.locations[entity] = newEntityLocation(maskKey, row)
}

func (a *archetypeAllocator) deleteComponents(entity Entity, components ...any) {
	if !entity.isAlive() {
		return
	}

	location, exists := a.locations[entity]
	if !exists {
		return
	}

	source := a.archetypes[location.maskKey]
	mask := slices.Clone(source.mask)
	excludeCompIds := []int{}

	for _, c := range components {
		id := a.registry.dataIdOf(c)

		if mask.unset(id) {
			excludeCompIds = append(excludeCompIds, id)
		}
	}

	if len(excludeCompIds) == 0 {
		return
	}

	maskKey := mask.key()
	target, ok := a.archetypes[maskKey]
	if !ok {
		entries := make([]componentEntry, 0, len(source.columns)-len(excludeCompIds))

		for _, column := range source.columns {
			if slices.Contains(excludeCompIds, column.componentId) {
				continue
			}

			entries = append(entries, newComponentEntry(column.componentId, column.elemType))
		}

		sortComponentEntries(entries)

		target = newArchetype(entries)
		a.archetypes[maskKey] = target
	}

	row := target.addEntity(entity)
	source.moveComponents(location.row, row, target, excludeCompIds...)
	swappedEntity, swappedRow := source.removeEntity(location.row)

	if swappedRow != -1 {
		swLocation := a.locations[swappedEntity]
		swLocation.row = swappedRow
		a.locations[swappedEntity] = swLocation
	}

	a.locations[entity] = newEntityLocation(maskKey, row)
}

func (a *archetypeAllocator) removeEntity(entity Entity) {
	if !entity.destroy() {
		return
	}

	location, ok := a.locations[entity]
	if !ok {
		return
	}

	archetype := a.archetypes[location.maskKey]

	swappedEntity, swappedRow := archetype.removeEntity(location.row)
	delete(a.locations, entity)

	if swappedRow != -1 {
		swLocation := a.locations[swappedEntity]
		swLocation.row = swappedRow
		a.locations[swappedEntity] = swLocation
	}
}

func (a *archetypeAllocator) matchingArchetypes(componentIds ...int) []*archetype {
	archetypes := []*archetype{}
	queryMask := maskBit(componentIds...)

	for _, archetype := range a.archetypes {
		if maskHasComponents(archetype.mask, queryMask) {
			archetypes = append(archetypes, archetype)
		}
	}

	return archetypes
}

func (a *archetypeAllocator) matchingArchetype(entity Entity) (*archetype, int) {
	location, ok := a.locations[entity]
	if !ok {
		return nil, -1
	}

	archetype := a.archetypes[location.maskKey]
	return archetype, location.row
}

func (a *archetypeAllocator) setComponents(archetype *archetype, row int, components []any) {
	for _, c := range components {
		id := a.registry.dataIdOf(c)
		archetype.addComponent(row, id, c)
	}
}
