package labutils

type NumberString interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int | ~uint | ~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64 | ~string
}

func InArray[T NumberString](needle T, hayStack []T) bool {
	for _, t := range hayStack {
		if t == needle {
			return true
		}
	}
	return false
}

type arrayMapForTrimCallback func(s, cutset string) string

// ArrayMapForTrim 数组过滤
// Example:
// data = ArrayMapForTrim(data, "'", strings.Trim)
// data = ArrayMapForTrim(data, "'", strings.TrimLeft)
// data = ArrayMapForTrim(data, "'", strings.TrimRight)
// data = ArrayMapForTrim(data, "'", strings.TrimPrefix)
// data = ArrayMapForTrim(data, "'", strings.TrimSuffix)
func ArrayMapForTrim(hayStack []string, cutset string, callback arrayMapForTrimCallback) []string {
	newArray := make([]string, 0)
	for _, s := range hayStack {
		s = callback(s, cutset)
		newArray = append(newArray, s)
	}
	return newArray
}
