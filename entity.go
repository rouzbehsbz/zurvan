package zurvan

import "sync/atomic"

// Entity is a handle returned by World.Spawn for use within a single world.
// Its index is read-only. The zero value is invalid.
//
// Limitation:
//   - The maximum number of entities is 2^32 - 1 (approximately 4.29 billion)
//   - Deleted entity indices are not reused. Spawn panics when IDs are exhausted.
type Entity struct {
	index uint32
	alive *atomic.Bool
}

func newEntity(index uint32) Entity {
	alive := new(atomic.Bool)
	alive.Store(true)

	return Entity{
		index: index,
		alive: alive,
	}
}

func (e Entity) Index() uint32 {
	return e.index
}

func (e Entity) isAlive() bool {
	return e.alive != nil && e.alive.Load()
}

func (e Entity) destroy() bool {
	return e.alive != nil && e.alive.CompareAndSwap(true, false)
}
