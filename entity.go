package zurvan

type Entity struct {
	Index uint32
}

func newEntity(index uint32) Entity {
	return Entity{
		Index: index,
	}
}
