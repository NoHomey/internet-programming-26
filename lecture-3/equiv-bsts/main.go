package main

import "fmt"

type BinaryTree struct {
	left  *BinaryTree
	right *BinaryTree
	value int
}

func Walk(root *BinaryTree, nodes chan int) {
	if root == nil {
		close(nodes)
		return
	}

	nodesLeft, nodesRight := make(chan int), make(chan int)
	go Walk(root.left, nodesLeft) // (*root).Left
	go Walk(root.right, nodesRight)

	for left := range nodesLeft {
		nodes <- left
	}
	nodes <- root.value
	for right := range nodesRight {
		nodes <- right
	}
	close(nodes)
}

func main() {
	root := &BinaryTree{
		left: &BinaryTree{
			left:  &BinaryTree{value: 2},
			right: nil,
			value: 4,
		},
		right: &BinaryTree{
			left:  nil,
			right: &BinaryTree{value: 32},
			value: 16,
		},
		value: 8,
	}

	nodes := make(chan int)
	go Walk(root, nodes)
	for value := range nodes {
		fmt.Println(value)
	}
}
