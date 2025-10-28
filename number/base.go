package number

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

func NumLen[T Number](number T) T {
	if number == 0 {
		return 1
	}

	var count T = 0
	if number < 0 {
		for number < 0 {
			number /= 10
			count++
		}
	} else {
		for number > 0 {
			number /= 10
			count++
		}
	}
	return count
}
