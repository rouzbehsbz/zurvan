package zurvan

func QueryMany1[A any](world *World, fn func([]Entity, []A)) {
	componentId := dataIdFor[A](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentId)
	for _, archetype := range archetypes {
		entities := archetype.entities
		column, _ := archetype.column(componentId)
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

		columnA, _ := archetype.column(componentAId)
		columnB, _ := archetype.column(componentBId)

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

		columnA, _ := archetype.column(componentAId)
		columnB, _ := archetype.column(componentBId)
		columnC, _ := archetype.column(componentCId)

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

		columnA, _ := archetype.column(componentAId)
		columnB, _ := archetype.column(componentBId)
		columnC, _ := archetype.column(componentCId)
		columnD, _ := archetype.column(componentDId)

		sliceA := columnA.asSlice().([]A)
		sliceB := columnB.asSlice().([]B)
		sliceC := columnC.asSlice().([]C)
		sliceD := columnD.asSlice().([]D)

		fn(entities, sliceA, sliceB, sliceC, sliceD)
	}
}

func QueryMany5[A, B, C, D, E any](world *World, fn func([]Entity, []A, []B, []C, []D, []E)) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)
	componentEId := dataIdFor[E](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentAId, componentBId, componentCId, componentDId)
	for _, archetype := range archetypes {
		entities := archetype.entities

		columnA, _ := archetype.column(componentAId)
		columnB, _ := archetype.column(componentBId)
		columnC, _ := archetype.column(componentCId)
		columnD, _ := archetype.column(componentDId)
		columnE, _ := archetype.column(componentEId)

		sliceA := columnA.asSlice().([]A)
		sliceB := columnB.asSlice().([]B)
		sliceC := columnC.asSlice().([]C)
		sliceD := columnD.asSlice().([]D)
		sliceE := columnE.asSlice().([]E)

		fn(entities, sliceA, sliceB, sliceC, sliceD, sliceE)
	}
}

func QueryMany6[A, B, C, D, E, F any](world *World, fn func([]Entity, []A, []B, []C, []D, []E, []F)) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)
	componentEId := dataIdFor[E](world.componentRegistry)
	componentFId := dataIdFor[F](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentAId, componentBId, componentCId, componentDId)
	for _, archetype := range archetypes {
		entities := archetype.entities

		columnA, _ := archetype.column(componentAId)
		columnB, _ := archetype.column(componentBId)
		columnC, _ := archetype.column(componentCId)
		columnD, _ := archetype.column(componentDId)
		columnE, _ := archetype.column(componentEId)
		columnF, _ := archetype.column(componentFId)

		sliceA := columnA.asSlice().([]A)
		sliceB := columnB.asSlice().([]B)
		sliceC := columnC.asSlice().([]C)
		sliceD := columnD.asSlice().([]D)
		sliceE := columnE.asSlice().([]E)
		sliceF := columnF.asSlice().([]F)

		fn(entities, sliceA, sliceB, sliceC, sliceD, sliceE, sliceF)
	}
}

func QueryOne1[A any](world *World, entity Entity) *A {
	componentId := dataIdFor[A](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil
	}

	column, ok := archetype.column(componentId)
	if !ok {
		return nil
	}

	slice := column.asSlice().([]A)

	return &slice[row]
}

func QueryOne2[A, B any](world *World, entity Entity) (*A, *B) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil
	}

	columnA, ok := archetype.column(componentAId)
	columnB, ok := archetype.column(componentBId)
	if !ok {
		return nil, nil
	}

	sliceA := columnA.asSlice().([]A)
	sliceB := columnB.asSlice().([]B)

	return &sliceA[row], &sliceB[row]
}

func QueryOne3[A, B, C any](world *World, entity Entity) (*A, *B, *C) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)

	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil
	}

	columnA, ok := archetype.column(componentAId)
	columnB, ok := archetype.column(componentBId)
	columnC, ok := archetype.column(componentCId)
	if !ok {
		return nil, nil, nil
	}

	sliceA := columnA.asSlice().([]A)
	sliceB := columnB.asSlice().([]B)
	sliceC := columnC.asSlice().([]C)

	return &sliceA[row], &sliceB[row], &sliceC[row]
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

	columnA, ok := archetype.column(componentAId)
	columnB, ok := archetype.column(componentBId)
	columnC, ok := archetype.column(componentCId)
	columnD, ok := archetype.column(componentDId)
	if !ok {
		return nil, nil, nil, nil
	}

	sliceA := columnA.asSlice().([]A)
	sliceB := columnB.asSlice().([]B)
	sliceC := columnC.asSlice().([]C)
	sliceD := columnD.asSlice().([]D)

	return &sliceA[row], &sliceB[row], &sliceC[row], &sliceD[row]
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

	columnA, ok := archetype.column(componentAId)
	columnB, ok := archetype.column(componentBId)
	columnC, ok := archetype.column(componentCId)
	columnD, ok := archetype.column(componentDId)
	columnE, ok := archetype.column(componentEId)
	if !ok {
		return nil, nil, nil, nil, nil
	}

	sliceA := columnA.asSlice().([]A)
	sliceB := columnB.asSlice().([]B)
	sliceC := columnC.asSlice().([]C)
	sliceD := columnD.asSlice().([]D)
	sliceE := columnE.asSlice().([]E)

	return &sliceA[row], &sliceB[row], &sliceC[row], &sliceD[row], &sliceE[row]
}
