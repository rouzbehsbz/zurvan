package zurvan

import "sync"

type entityAllocator struct {
	generations []int
	availables  []int
	mu          sync.Mutex
}

func newEntityAllocator() *entityAllocator {
	return &entityAllocator{
		generations: []int{},
		availables:  []int{},
		mu:          sync.Mutex{},
	}
}

func (e *entityAllocator) create() Entity {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(e.availables) == 0 {
		index := len(e.generations)
		e.generations = append(e.generations, 0)

		return newEntity(index, 0)
	}

	lastIndex := len(e.availables) - 1
	index := e.availables[lastIndex]

	e.availables = e.availables[:lastIndex]
	generation := e.generations[index]

	return newEntity(index, generation)
}

func (e *entityAllocator) delete(entity Entity) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.isAliveUnsafe(entity) {
		e.availables = append(e.availables, entity.Index)
		e.generations[entity.Index] += 1
	}
}

func (e *entityAllocator) isAliveUnsafe(entity Entity) bool {
	if entity.Index >= len(e.generations) {
		return false
	}

	generation := e.generations[entity.Index]

	return generation == entity.Generation
}
