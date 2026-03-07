package zurvan

func QueryMany1[A any](world *World, fn func([]Entity, []A)) {
	componentId := dataIdFor[A](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentId)
	for _, archetype := range archetypes {
		entities := archetype.entities
		column := archetype.column(componentId)
		slice := column.asSlice().([]A)

		fn(entities, slice)
	}
}

func QueryMany2[A, B any](world *World, fn func([]Entity, []A, []B)) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentAId, componentBId)
	for _, archetype := range archetypes {
		entities := archetype.entities

		columnA := archetype.column(componentAId)
		columnB := archetype.column(componentBId)

		sliceA := columnA.asSlice().([]A)
		sliceB := columnB.asSlice().([]B)

		fn(entities, sliceA, sliceB)
	}
}

func QueryMany3[A, B, C any](world *World, fn func([]Entity, []A, []B, []C)) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentAId, componentBId, componentCId)
	for _, archetype := range archetypes {
		entities := archetype.entities

		columnA := archetype.column(componentAId)
		columnB := archetype.column(componentBId)
		columnC := archetype.column(componentCId)

		sliceA := columnA.asSlice().([]A)
		sliceB := columnB.asSlice().([]B)
		sliceC := columnC.asSlice().([]C)

		fn(entities, sliceA, sliceB, sliceC)
	}
}

func QueryMany4[A, B, C, D any](world *World, fn func([]Entity, []A, []B, []C, []D)) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentAId, componentBId, componentCId, componentDId)
	for _, archetype := range archetypes {
		entities := archetype.entities

		columnA := archetype.column(componentAId)
		columnB := archetype.column(componentBId)
		columnC := archetype.column(componentCId)
		columnD := archetype.column(componentDId)

		sliceA := columnA.asSlice().([]A)
		sliceB := columnB.asSlice().([]B)
		sliceC := columnC.asSlice().([]C)
		sliceD := columnD.asSlice().([]D)

		fn(entities, sliceA, sliceB, sliceC, sliceD)
	}
}

func QueryOne1[A any](world *World, entity Entity) *A {
	componentId := dataIdFor[A](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil
	}

	column := archetype.column(componentId)
	data := column.get(row).(A)

	return &data
}

func QueryOne2[A, B any](world *World, entity Entity) (*A, *B) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil
	}

	columnA := archetype.column(componentAId)
	columnB := archetype.column(componentBId)

	dataA := columnA.get(row).(A)
	dataB := columnB.get(row).(B)

	return &dataA, &dataB
}

func QueryOne3[A, B, C any](world *World, entity Entity) (*A, *B, *C) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil
	}

	columnA := archetype.column(componentAId)
	columnB := archetype.column(componentBId)
	columnC := archetype.column(componentCId)

	dataA := columnA.get(row).(A)
	dataB := columnB.get(row).(B)
	dataC := columnC.get(row).(C)

	return &dataA, &dataB, &dataC
}

func QueryOne4[A, B, C, D any](world *World, entity Entity) (*A, *B, *C, *D) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil
	}

	columnA := archetype.column(componentAId)
	columnB := archetype.column(componentBId)
	columnC := archetype.column(componentCId)
	columnD := archetype.column(componentDId)

	dataA := columnA.get(row).(A)
	dataB := columnB.get(row).(B)
	dataC := columnC.get(row).(C)
	dataD := columnD.get(row).(D)

	return &dataA, &dataB, &dataC, &dataD
}

func QueryOne5[A, B, C, D, E any](world *World, entity Entity) (*A, *B, *C, *D, *E) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)
	componentEId := dataIdFor[E](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil, nil
	}

	columnA := archetype.column(componentAId)
	columnB := archetype.column(componentBId)
	columnC := archetype.column(componentCId)
	columnD := archetype.column(componentDId)
	columnE := archetype.column(componentEId)

	dataA := columnA.get(row).(A)
	dataB := columnB.get(row).(B)
	dataC := columnC.get(row).(C)
	dataD := columnD.get(row).(D)
	dataE := columnE.get(row).(E)

	return &dataA, &dataB, &dataC, &dataD, &dataE
}
