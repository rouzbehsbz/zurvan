package zurvan

import "reflect"

type vector struct {
	data reflect.Value
}

func newVector(elemType reflect.Type) *vector {
	return &vector{
		data: reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0),
	}
}

func (v *vector) resize(length int) {
	if length <= v.data.Len() {
		v.data = v.data.Slice(0, length)
		return
	}

	newCap := length
	newLen := length

	newSlice := reflect.MakeSlice(v.data.Type(), newLen, newCap)

	reflect.Copy(newSlice, v.data)

	v.data = newSlice
}

func (v *vector) len() int {
	return v.data.Len()
}

func (v *vector) remove(index int) {
	length := v.len()
	if index >= length {
		return
	}

	lastIndex := length - 1

	if index != lastIndex {
		lastValue := v.data.Index(lastIndex)
		v.data.Index(index).Set(lastValue)
	}

	v.data = v.data.Slice(0, lastIndex)
}

func (v *vector) set(index int, value any) {
	if index >= v.len() {
		return
	}

	val := reflect.ValueOf(value)
	v.data.Index(index).Set(val)
}

func (v *vector) get(index int) any {
	if index >= v.len() {
		return nil
	}

	return v.data.Index(index).Interface()
}

func (v *vector) asSlice() any {
	return v.data.Interface()
}

func (v *vector) push(value any) {
	val := reflect.ValueOf(value)

	v.data = reflect.Append(v.data, val)
}
