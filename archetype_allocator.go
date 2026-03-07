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

type archetypeAllocator struct {
	archetypes map[mask]*archetype
	locations  map[Entity]entityLocation

	registry *registry
}

func NewArchetypeAllocator(registry *registry) *archetypeAllocator {
	return &archetypeAllocator{
		archetypes: make(map[mask]*archetype),
		locations:  make(map[Entity]entityLocation),
		registry:   registry,
	}
}

func (a *archetypeAllocator) addComponents(entity Entity, components ...any) {
	var mask mask

	entries := make([]componentEntry, 0, len(components))

	for _, c := range components {
		id := a.registry.dataIdOf(c)
		mask |= maskBit(id)
		entries = append(entries, newComponentEntry(id, reflect.TypeOf(c)))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].id < entries[j].id
	})

	location, exists := a.locations[entity]
	if exists {
		target := a.archetypes[location.mask]

		if maskHasComponents(location.mask, mask) {
			a.setComponents(target, location.row, components)
			return
		}

		target, ok := a.archetypes[mask]
		if !ok {
			target = newArchetype(entries)
			a.archetypes[mask] = target
		}

		source := a.archetypes[location.mask]
		newRow := target.addEntity(entity)

		source.moveComponents(location.row, newRow, target)
		source.removeEntity(location.row)

		a.locations[entity] = newEntityLocation(mask, newRow)
		return
	}

	targetArchetype, ok := a.archetypes[mask]
	if !ok {
		targetArchetype = newArchetype(entries)
		a.archetypes[mask] = targetArchetype
	}

	row := targetArchetype.addEntity(entity)

	a.setComponents(targetArchetype, row, components)
	a.locations[entity] = newEntityLocation(mask, row)
}

func (a *archetypeAllocator) DeleteComponents(entity Entity, components ...any) {
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
		entries := make([]componentEntry, 0, len(components))

		for _, column := range source.columns {
			if slices.Contains(excludeCompIds, column.componentId) {
				continue
			}

			entries = append(entries, newComponentEntry(column.componentId, column.elemType))
		}

		sort.Slice(entries, func(i, j int) bool {
			return entries[i].id < entries[j].id
		})

		target = newArchetype(entries)
		a.archetypes[mask] = target
	}

	row := target.addEntity(entity)
	source.moveComponents(location.row, row, target, excludeCompIds...)
	source.removeEntity(location.row)
}

func (a *archetypeAllocator) RemoveEntity(entity Entity) {
	location := a.locations[entity]
	archetype := a.archetypes[location.mask]

	archetype.removeEntity(location.row)
	delete(a.locations, entity)
}

func (a *archetypeAllocator) MatchingArchetypes(componentIds ...int) []*archetype {
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
