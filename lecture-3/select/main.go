package main

func worker(a []int, transform func(int) int, c chan int) {
	for x := range a {
		c <- transform(x)
	}
	close(c)
}

func sum() int {
	s := 0

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go worker([]int{2, 3, 1, 2}, func(x int) int {
		return x * x
	}, c1)
	go worker([]int{2, 3, 1, 2}, func(x int) int {
		return 2*x + 1
	}, c2)
	go worker([]int{2, 3, 1, 2}, func(x int) int {
		return x + 3
	}, c3)

	for {
		select {
		case x1 := <-c1:
			s += x1
		case x2 := <-c2:
			s += x2
		case x3 := <-c3:
			s += x3
		}
	}

	return s
}
