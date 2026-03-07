package zurvan

type mask = uint64

func maskBit(componentIds ...int) mask {
	var mask mask

	for _, componentId := range componentIds {
		mask |= 1 << componentId
	}

	return mask
}

func maskHasComponents(mask, query mask) bool {
	return mask&query == query
}
