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

func AreEquiv(left *BinaryTree, right *BinaryTree) bool {
	leftNodes, rightNodes := make(chan int), make(chan int)
	go Walk(left, leftNodes)
	go Walk(right, rightNodes)

	for {
		leftNode, leftOk := <-leftNodes
		rightNode, rightOk := <-rightNodes

		if leftOk != rightOk {
			return false
		}

		if leftOk == false {
			return true
		}

		if leftNode != rightNode {
			return false
		}
	}
}

func main() {
	root1 := &BinaryTree{
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

	root2 := &BinaryTree{
		left: &BinaryTree{
			left: &BinaryTree{
				left:  nil,
				right: nil,
				value: 2,
			},
			right: &BinaryTree{
				left:  nil,
				right: nil,
				value: 8,
			},

			value: 4,
		},
		right: &BinaryTree{
			left:  nil,
			right: nil,
			value: 31,
		},

		value: 16,
	}

	fmt.Println(AreEquiv(root1, root2))
}
