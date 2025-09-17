type Number interface {
	int | int64 | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}
