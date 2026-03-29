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

func QueryOne1[A any](world *World, entity Entity) *A {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil
	}

	componentId := dataIdFor[A](world.componentRegistry)

	column, ok := archetype.column(componentId)

	var result *A
	if ok {
		slice := column.asSlice().([]A)
		result = &slice[row]
	}

	return result
}

func QueryOne2[A, B any](world *World, entity Entity) (*A, *B) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil
	}

	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)

	columnA, okA := archetype.column(componentAId)
	columnB, okB := archetype.column(componentBId)

	var resultA *A
	var resultB *B

	if okA {
		sliceA := columnA.asSlice().([]A)
		resultA = &sliceA[row]
	}
	if okB {
		sliceB := columnB.asSlice().([]B)
		resultB = &sliceB[row]
	}

	return resultA, resultB
}

func QueryOne3[A, B, C any](world *World, entity Entity) (*A, *B, *C) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil
	}

	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)

	columnA, okA := archetype.column(componentAId)
	columnB, okB := archetype.column(componentBId)
	columnC, okC := archetype.column(componentCId)

	var resultA *A
	var resultB *B
	var resultC *C

	if okA {
		sliceA := columnA.asSlice().([]A)
		resultA = &sliceA[row]
	}
	if okB {
		sliceB := columnB.asSlice().([]B)
		resultB = &sliceB[row]
	}
	if okC {
		sliceC := columnC.asSlice().([]C)
		resultC = &sliceC[row]
	}

	return resultA, resultB, resultC
}

func QueryOne4[A, B, C, D any](world *World, entity Entity) (*A, *B, *C, *D) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil
	}

	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)

	columnA, okA := archetype.column(componentAId)
	columnB, okB := archetype.column(componentBId)
	columnC, okC := archetype.column(componentCId)
	columnD, okD := archetype.column(componentDId)

	var resultA *A
	var resultB *B
	var resultC *C
	var resultD *D

	if okA {
		sliceA := columnA.asSlice().([]A)
		resultA = &sliceA[row]
	}
	if okB {
		sliceB := columnB.asSlice().([]B)
		resultB = &sliceB[row]
	}
	if okC {
		sliceC := columnC.asSlice().([]C)
		resultC = &sliceC[row]
	}
	if okD {
		sliceD := columnD.asSlice().([]D)
		resultD = &sliceD[row]
	}

	return resultA, resultB, resultC, resultD
}

func QueryOne5[A, B, C, D, E any](world *World, entity Entity) (*A, *B, *C, *D, *E) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil, nil
	}

	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)
	componentEId := dataIdFor[E](world.componentRegistry)

	columnA, okA := archetype.column(componentAId)
	columnB, okB := archetype.column(componentBId)
	columnC, okC := archetype.column(componentCId)
	columnD, okD := archetype.column(componentDId)
	columnE, okE := archetype.column(componentEId)

	var resultA *A
	var resultB *B
	var resultC *C
	var resultD *D
	var resultE *E

	if okA {
		sliceA := columnA.asSlice().([]A)
		resultA = &sliceA[row]
	}
	if okB {
		sliceB := columnB.asSlice().([]B)
		resultB = &sliceB[row]
	}
	if okC {
		sliceC := columnC.asSlice().([]C)
		resultC = &sliceC[row]
	}
	if okD {
		sliceD := columnD.asSlice().([]D)
		resultD = &sliceD[row]
	}
	if okE {
		sliceE := columnE.asSlice().([]E)
		resultE = &sliceE[row]
	}

	return resultA, resultB, resultC, resultD, resultE
}

func QueryOne6[A, B, C, D, E, F any](world *World, entity Entity) (*A, *B, *C, *D, *E, *F) {
	archetype, row := world.archetypeAllocator.matchingArchetype(entity)
	if archetype == nil || row == -1 {
		return nil, nil, nil, nil, nil, nil
	}

	componentAId := dataIdFor[A](world.componentRegistry)
	componentBId := dataIdFor[B](world.componentRegistry)
	componentCId := dataIdFor[C](world.componentRegistry)
	componentDId := dataIdFor[D](world.componentRegistry)
	componentEId := dataIdFor[E](world.componentRegistry)
	componentFId := dataIdFor[F](world.componentRegistry)

	columnA, okA := archetype.column(componentAId)
	columnB, okB := archetype.column(componentBId)
	columnC, okC := archetype.column(componentCId)
	columnD, okD := archetype.column(componentDId)
	columnE, okE := archetype.column(componentEId)
	columnF, okF := archetype.column(componentFId)

	var resultA *A
	var resultB *B
	var resultC *C
	var resultD *D
	var resultE *E
	var resultF *F

	if okA {
		sliceA := columnA.asSlice().([]A)
		resultA = &sliceA[row]
	}
	if okB {
		sliceB := columnB.asSlice().([]B)
		resultB = &sliceB[row]
	}
	if okC {
		sliceC := columnC.asSlice().([]C)
		resultC = &sliceC[row]
	}
	if okD {
		sliceD := columnD.asSlice().([]D)
		resultD = &sliceD[row]
	}
	if okE {
		sliceE := columnE.asSlice().([]E)
		resultE = &sliceE[row]
	}
	if okF {
		sliceF := columnF.asSlice().([]F)
		resultF = &sliceF[row]
	}

	return resultA, resultB, resultC, resultD, resultE, resultF
}
