package zurvan

// Entity represents a unique object in the world.
// Index is the unique identifier assigned to the entity.
//
// Limitation:
//   - The maximum number of entities is 2^32 - 1 (approximately 4.29 billion)
//   - Deleted entity indices will not be reused, so craeting and deleting large number of
//     entities may eventually exhaust the available indices
type Entity struct {
	Index uint32
}

func newEntity(index uint32) Entity {
	return Entity{
		Index: index,
	}
}
