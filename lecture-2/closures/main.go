package main

import "fmt"

func incrementer(start int) func(int) int {
	s := start

	return func(x int) int {
		s += x
		return s
	}
}

type compileInc struct {
	s *int
}

func (c *compileInc) incrementer(x int) int {
	(*c.s) += x
	return *c.s
}

func main() {
	i5 := incrementer(5)

	i5(3)

	fmt.Println(i5(2))

	start := 5
	ci5 := compileInc{s: &start}
	fmt.Println(ci5.incrementer(4))

}
