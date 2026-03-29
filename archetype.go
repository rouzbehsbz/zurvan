package zurvan

import (
	"reflect"
	"slices"
)

type columnEntry struct {
	componentId int
	elemType    reflect.Type
	column      column
}

func newColumnEntry(componentId int, elemType reflect.Type) columnEntry {
	return columnEntry{
		componentId: componentId,
		elemType:    elemType,
		column:      newVector(elemType),
	}
}

type archetype struct {
	entities []Entity
	columns  []columnEntry

	componentIndex map[int]int
}

func newArchetype(entries []componentEntry) *archetype {
	columns := []columnEntry{}
	componentIndex := make(map[int]int, len(entries))

	for _, entry := range entries {
		index := len(columns)

		columns = append(columns, newColumnEntry(entry.id, entry.elemType))

		componentIndex[entry.id] = index
	}

	return &archetype{
		entities:       []Entity{},
		columns:        columns,
		componentIndex: componentIndex,
	}
}

func (a *archetype) addEntity(entity Entity) int {
	row := len(a.entities)
	a.entities = append(a.entities, entity)

	for _, entry := range a.columns {
		entry.column.resize(len(a.entities))
	}

	return row
}

func (a *archetype) removeEntity(row int) (Entity, int) {
	var swapped Entity
	var swappedRow int = -1

	length := len(a.entities)
	if row >= length {
		return swapped, swappedRow
	}

	lastIndex := length - 1

	if row != lastIndex {
		swapped = a.entities[lastIndex]
		a.entities[row] = swapped
		swappedRow = row
	}

	a.entities = a.entities[:lastIndex]

	for _, entry := range a.columns {
		entry.column.remove(row)
	}

	return swapped, swappedRow
}

func (a *archetype) addComponent(row int, componentId int, component any) {
	columnIndex := a.componentIndex[componentId]
	entry := a.columns[columnIndex]
	entry.column.set(row, component)
}

func (a *archetype) moveComponents(row int, dstRow int, dstArchetype *archetype, excludeCompIds ...int) {
	for _, entry := range a.columns {
		component := entry.column.get(row)

		if slices.Contains(excludeCompIds, entry.componentId) {
			continue
		}

		dstArchetype.addComponent(dstRow, entry.componentId, component)
	}
}

func (a *archetype) column(componentId int) (column, bool) {
	index, ok := a.componentIndex[componentId]
	if !ok {
		return nil, false
	}

	entry := a.columns[index]

	return entry.column, true
}
