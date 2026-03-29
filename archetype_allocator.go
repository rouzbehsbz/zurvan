package zurvan

import (
	"reflect"
	"slices"
	"sort"
)

type entityLocation struct {
	mask mask
	row  int
}

func newEntityLocation(mask mask, row int) entityLocation {
	return entityLocation{
		mask: mask,
		row:  row,
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
	archetypes map[mask]*archetype
	locations  map[Entity]entityLocation

	registry *registry
}

func newArchetypeAllocator(registry *registry) *archetypeAllocator {
	return &archetypeAllocator{
		archetypes: make(map[mask]*archetype),
		locations:  make(map[Entity]entityLocation),
		registry:   registry,
	}
}

func (a *archetypeAllocator) addComponents(entity Entity, components ...any) {
	location, exists := a.locations[entity]
	if exists {
		target := a.archetypes[location.mask]

		var mask mask

		for _, c := range components {
			id := a.registry.dataIdOf(c)
			mask |= maskBit(id)
		}

		if maskHasComponents(location.mask, mask) {
			a.setComponents(target, location.row, components)
			return
		}

		newMask := location.mask | mask

		source := a.archetypes[location.mask]
		target, ok := a.archetypes[newMask]
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
			a.archetypes[newMask] = target
		}

		newRow := target.addEntity(entity)

		a.setComponents(target, newRow, components)
		source.moveComponents(location.row, newRow, target)
		swappedEntity, swappedRow := source.removeEntity(location.row)

		if swappedRow != -1 {
			swLocation := a.locations[swappedEntity]
			swLocation.row = swappedRow
			a.locations[swappedEntity] = swLocation
		}

		a.locations[entity] = newEntityLocation(newMask, newRow)
		return
	}

	var mask mask

	entries := make([]componentEntry, 0, len(components))

	for _, c := range components {
		id := a.registry.dataIdOf(c)
		mask |= maskBit(id)
		entries = append(entries, newComponentEntry(id, reflect.TypeOf(c)))
	}

	sortComponentEntries(entries)

	targetArchetype, ok := a.archetypes[mask]
	if !ok {
		targetArchetype = newArchetype(entries)
		a.archetypes[mask] = targetArchetype
	}

	row := targetArchetype.addEntity(entity)

	a.setComponents(targetArchetype, row, components)
	a.locations[entity] = newEntityLocation(mask, row)
}

func (a *archetypeAllocator) deleteComponents(entity Entity, components ...any) {
	location, exists := a.locations[entity]
	if !exists {
		return
	}

	mask := location.mask
	excludeCompIds := []int{}

	for _, c := range components {
		id := a.registry.dataIdOf(c)
		mask &= ^maskBit(id)
		excludeCompIds = append(excludeCompIds, id)
	}

	source := a.archetypes[location.mask]
	target, ok := a.archetypes[mask]
	if !ok {
		entries := make([]componentEntry, 0, len(source.columns)-len(components))

		for _, column := range source.columns {
			if slices.Contains(excludeCompIds, column.componentId) {
				continue
			}

			entries = append(entries, newComponentEntry(column.componentId, column.elemType))
		}

		sortComponentEntries(entries)

		target = newArchetype(entries)
		a.archetypes[mask] = target
	}

	row := target.addEntity(entity)
	source.moveComponents(location.row, row, target, excludeCompIds...)
	swappedEntity, swappedRow := source.removeEntity(location.row)

	if swappedRow != -1 {
		swLocation := a.locations[swappedEntity]
		swLocation.row = swappedRow
		a.locations[swappedEntity] = swLocation
	}

	a.locations[entity] = newEntityLocation(mask, row)
}

func (a *archetypeAllocator) removeEntity(entity Entity) {
	location, ok := a.locations[entity]
	if !ok {
		return
	}

	archetype := a.archetypes[location.mask]

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

	for mask, archetype := range a.archetypes {
		queryMask := maskBit(componentIds...)

		if maskHasComponents(mask, queryMask) {
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

	archetype := a.archetypes[location.mask]
	return archetype, location.row
}

func (a *archetypeAllocator) setComponents(archetype *archetype, row int, components []any) {
	for _, c := range components {
		id := a.registry.dataIdOf(c)
		archetype.addComponent(row, id, c)
	}
}
