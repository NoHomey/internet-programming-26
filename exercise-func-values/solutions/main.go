package main

import "fmt"

/*
func nextValue(x0 int, x1 int) func() int {
	prev, curr := x0, x1
	return func() int {
		prev, curr = curr, prev+curr
		return curr
	}
}
*/

/*
func nextValue(x0 int, x1 int, next func(xi1 int, xi2 int) int) func() int {
	prev, curr := x0, x1
	return func() int {
		prev, curr = curr, next(prev, curr)
		return curr
	}
}

func main() {
	fib := nextValue(0, 1, func(xi1, xi2 int) int {
		return xi1 + xi2
	})
	fmt.Println(fib()) // 1
	fmt.Println(fib()) // 2
	fmt.Println(fib()) // 3
	fmt.Println(fib()) // 5
}
*/

func nextValue(start []int, next func(prev []int) int) func() int {
	values := make([]int, len(start), len(start))
	n := copy(values, start)
	if n < len(start) {
		panic(fmt.Sprintf("failed to copy start, copied %d out of %d elements", n, len(start)))
	}
	
	return func() int {
		curr := next(values)
		values = append(values, curr)
		return curr
	}

}

func main() {
	multiplyAll := nextValue([]int{2, 3, 5}, func(prev []int) int {
		x := 1
		for _, v := range prev {
			x *= v
		}
		return x
	})

	fmt.Println(multiplyAll()) // 30
	fmt.Println(multiplyAll()) // 900
}
