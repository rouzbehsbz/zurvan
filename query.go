package zurvan

// Querying entities with a single component
//
// Iterates over all entities that have a component(s) of the specified type(s)
// The provided slices are aligned by index, for any index i, entities[i] owns
// components stored at index i in each component slice.
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

// Querying entities with two components
//
// Iterates over all entities that have a component(s) of the specified type(s)
// The provided slices are aligned by index, for any index i, entities[i] owns
// components stored at index i in each component slice.
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

// Querying entities with three components
//
// Iterates over all entities that have a component(s) of the specified type(s)
// The provided slices are aligned by index, for any index i, entities[i] owns
// components stored at index i in each component slice.
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

// Querying entities with four components
//
// Iterates over all entities that have a component(s) of the specified type(s)
// The provided slices are aligned by index, for any index i, entities[i] owns
// components stored at index i in each component slice.
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

// Querying entities with five components
//
// Iterates over all entities that have a component(s) of the specified type(s)
// The provided slices are aligned by index, for any index i, entities[i] owns
// components stored at index i in each component slice.
func QueryMany5[A, B, C, D, E any](world *World, fn func([]Entity, []A, []B, []C, []D, []E)) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)
	componentEId := dataIdFor[E](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentAId, componentBId, componentCId, componentDId, componentEId)
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

// Querying entities with six components
//
// Iterates over all entities that have a component(s) of the specified type(s)
// The provided slices are aligned by index, for any index i, entities[i] owns
// components stored at index i in each component slice.
func QueryMany6[A, B, C, D, E, F any](world *World, fn func([]Entity, []A, []B, []C, []D, []E, []F)) {
	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)
	componentEId := dataIdFor[E](world.componentRegistry)
	componentFId := dataIdFor[F](world.componentRegistry)

	archetypes := world.archetypeAllocator.matchingArchetypes(componentAId, componentBId, componentCId, componentDId, componentEId, componentFId)
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

// Querying a single entity with one component
//
// Returns a pointer to the component if the entity has it, otherwise returns nil
func QueryOne1[A any](world *World, entity Entity) *A {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil
	}

	resultA := componentPtr[A](world, archetype, row)

	return resultA
}

// Querying a single entity with two components
//
// Returns pointers to the components if the entity has them, otherwise returns nil
func QueryOne2[A, B any](world *World, entity Entity) (*A, *B) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil
	}

	resultA := componentPtr[A](world, archetype, row)
	resultB := componentPtr[B](world, archetype, row)

	return resultA, resultB
}

// Querying a single entity with three components
//
// Returns pointers to the components if the entity has them, otherwise returns nil
func QueryOne3[A, B, C any](world *World, entity Entity) (*A, *B, *C) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil
	}

	resultA := componentPtr[A](world, archetype, row)
	resultB := componentPtr[B](world, archetype, row)
	resultC := componentPtr[C](world, archetype, row)

	return resultA, resultB, resultC
}

// Querying a single entity with four components
//
// Returns pointers to the components if the entity has them, otherwise returns nil
func QueryOne4[A, B, C, D any](world *World, entity Entity) (*A, *B, *C, *D) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil
	}

	resultA := componentPtr[A](world, archetype, row)
	resultB := componentPtr[B](world, archetype, row)
	resultC := componentPtr[C](world, archetype, row)
	resultD := componentPtr[D](world, archetype, row)

	return resultA, resultB, resultC, resultD
}

// Querying a single entity with five components
//
// Returns pointers to the components if the entity has them, otherwise returns nil
func QueryOne5[A, B, C, D, E any](world *World, entity Entity) (*A, *B, *C, *D, *E) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil, nil
	}

	resultA := componentPtr[A](world, archetype, row)
	resultB := componentPtr[B](world, archetype, row)
	resultC := componentPtr[C](world, archetype, row)
	resultD := componentPtr[D](world, archetype, row)
	resultE := componentPtr[E](world, archetype, row)

	return resultA, resultB, resultC, resultD, resultE
}

// Querying a single entity with six components
//
// Returns pointers to the components if the entity has them, otherwise returns nil
func QueryOne6[A, B, C, D, E, F any](world *World, entity Entity) (*A, *B, *C, *D, *E, *F) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil, nil, nil
	}

	resultA := componentPtr[A](world, archetype, row)
	resultB := componentPtr[B](world, archetype, row)
	resultC := componentPtr[C](world, archetype, row)
	resultD := componentPtr[D](world, archetype, row)
	resultE := componentPtr[E](world, archetype, row)
	resultF := componentPtr[F](world, archetype, row)

	return resultA, resultB, resultC, resultD, resultE, resultF
}

func componentPtr[T any](world *World, archetype *archetype, row int) *T {
	componentId := dataIdFor[T](world.componentRegistry)

	column, ok := archetype.column(componentId)
	if !ok {
		return nil
	}

	slice := column.asSlice().([]T)
	return &slice[row]
}
