package labutils

type NumberString interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int | ~uint | ~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64 | ~string
}

func InSlice[T NumberString](needle T, hayStack []T) bool {
	for _, t := range hayStack {
		if t == needle {
			return true
		}
	}
	return false
}
