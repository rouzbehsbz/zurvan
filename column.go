package zurvan

type column interface {
	resize(length int)
	len() int
	remove(index int)
	set(index int, value any)
	get(index int) any
	asSlice() any
	push(value any)
}
