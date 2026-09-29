package main

import "fmt"

func main() {
	var c chan int = make(chan int, 1)

	c <- 9
	x := <-c
	fmt.Println(x)

	y, ok := <-c
	if ok {
		fmt.Println(y)
	} else {
		fmt.Println("No value", y)
	}
}
