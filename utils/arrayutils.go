package utils

func InitArray[V int](inCount int, inDefault V) []V {
	items := make([]V, inCount)
	for i := range inCount {
		items[i] = inDefault
	}
	return items
}
