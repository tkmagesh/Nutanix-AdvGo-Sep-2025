type Adder[T any] interface {
	Add(a, b T) T
}

type IntAdder struct{}

func (IntAdder) Add(a, b int) int { return a + b }
