func MapReduce[T any, R any](
	inputs []T,
	mapFn func(T) R,
	reduceFn func(R, R) R,
) R {
	out := make(chan R)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(val T) {
			defer wg.Done()
			out <- mapFn(val)
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	var result R
	first := true
	for r := range out {
		if first {
			result = r
			first = false
		} else {
			result = reduceFn(result, r)
		}
	}
	return result
}


sum := MapReduce[int, int](
	[]int{1, 2, 3, 4},
	func(i int) int { return i * i },
	func(a, b int) int { return a + b },
)
fmt.Println("Sum of squares:", sum)
