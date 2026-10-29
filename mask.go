package zurvan

import "encoding/binary"

type mask []uint64

func maskBit(componentIds ...int) mask {
	var result mask

	for _, id := range componentIds {
		result.set(id)
	}

	return result
}

func (m *mask) set(id int) {
	if id < 0 {
		panic("zurvan: component ID must be non-negative")
	}

	word := id / 64
	if word >= len(*m) {
		*m = append(*m, make([]uint64, word+1-len(*m))...)
	}

	(*m)[word] |= uint64(1) << (id % 64)
}

func (m mask) has(id int) bool {
	return id >= 0 && id/64 < len(m) && m[id/64]&(uint64(1)<<(id%64)) != 0
}

func (m mask) unset(id int) bool {
	if !m.has(id) {
		return false
	}

	m[id/64] &^= uint64(1) << (id % 64)
	return true
}

func maskUnion(a, b mask) mask {
	result := make(mask, max(len(a), len(b)))
	copy(result, a)
	for i, word := range b {
		result[i] |= word
	}
	return result
}

func maskHasComponents(m, query mask) bool {
	for i, word := range query {
		if word != 0 && (i >= len(m) || m[i]&word != word) {
			return false
		}
	}
	return true
}

func (m mask) key() string {
	end := len(m)

	for end > 0 && m[end-1] == 0 {
		end--
	}

	if end == 0 {
		return ""
	}

	data := make([]byte, end*8)
	for i, word := range m[:end] {
		binary.LittleEndian.PutUint64(data[i*8:], word)
	}

	return string(data)
}
