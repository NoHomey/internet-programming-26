# Task 1

Implement a generalized Fibonacci sequence funcion generator that given the initial starting values can generate the next value on each call.

```go
func nextValue(x0 int, x1 int) func() int {
	// implement
}

func main() {
	fib := nextValue(0, 1)
	fmt.Println(fib()) // 1
	fmt.Println(fib()) // 2
	fmt.Println(fib()) // 3
}
```

# Task 2

Generalize even further the code by addding an argument that provides the logic for the next sequence value.

```go
func nextValue(x0 int, x1 int, next func(xi1 int, xi2 int) int) func() int {
	// implement
}
```

# Task 3

For this task generalize the above code by allowing N starting values and a function that returns the next value based on all **previous** values.

```go
func nextValue(start int[], next func(prev []int) int) func() int {
	// implement
}
```

Example:

```go
func main() {
	multiplyAll := nextValue([]int{2, 3, 5}, func(prev []int) int {
		x := 1
		for v := range prev {
			x *= v
		}
		return x
	})

	fmt.Println(multiplyAll()) // 30
	fmt.Println(multiplyAll()) // 900
}
```
