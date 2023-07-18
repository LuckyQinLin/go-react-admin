package utils

// Types 类型集合
type Types interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func In[T Types](target T, arr []T) (bool, int) {
	for idx, item := range arr {
		if target == item {
			return true, idx
		}
	}
	return false, 0
}
