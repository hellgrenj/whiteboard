package fibonacci

import "fmt"

var slowFibCounter int = 0

func SlowFib(fibbNum int) {
	for i := 0; i < fibbNum; i++ {
		fmt.Printf("%d ", calculateFib(i))
	}
	fmt.Printf("recursive SlowFib() called %d times to print out %d fibonacci numbers\n", slowFibCounter, fibbNum)
}
func calculateFib(n int) int {
	slowFibCounter++
	if n < 2 {
		return n
	} else {
		return calculateFib(n-1) + calculateFib(n-2)
	}
}
func FasterFib(x int) {
	iterationCounter := 0
	fibs := []int{0, 1}
	for _, v := range fibs {
		fmt.Printf("%d ", v)
		iterationCounter++
	}
	for i := 2; i < x; i++ {
		iterationCounter++
		next := fibs[i-1] + fibs[i-2]
		fmt.Printf("%d ", next)
		fibs = append(fibs, next)
	}
	fmt.Printf("FasterFib() using memoization calculated %d fibonacci numbers with %d iterations\n", x, iterationCounter)
}
