package main

import "fmt"

func mergeSort(a []int) {
	c := make(chan struct{})
	mergeSortHelper(a, c)
	<-c
}

func mergeSortHelper(a []int, c chan struct{}) {
	if len(a) < 2 {
		close(c)
		return
	}

	mid := len(a) / 2
	left := a[:mid]
	right := a[mid:]

	leftC := make(chan struct{})
	rightC := make(chan struct{})

	go mergeSortHelper(left, leftC)
	go mergeSortHelper(right, rightC)

	<-leftC
	<-rightC

	tmp := make([]int, 0, len(a))

	n := len(a)

	leftI, rightI := 0, 0
	leftN, rightN := len(left), len(right)

	for i := 0; i < n; i++ {
		hasLeft := leftI < leftN
		hasRight := rightI < rightN

		if hasLeft && hasRight {
			if left[leftI] < right[rightI] {
				tmp = append(tmp, left[leftI])
				leftI++
				continue
			}
			tmp = append(tmp, right[rightI])
			rightI++
			continue
		}

		if hasLeft {
			tmp = append(tmp, left[leftI])
			leftI++
			continue
		}

		tmp = append(tmp, right[rightI])
		rightI++
		continue
	}

	for i := 0; i < n; i++ {
		a[i] = tmp[i]
	}

	close(c)
}

func main() {
	a := []int{2, 1, 5, 3, 8, 6, 7}
	mergeSort(a)
	fmt.Println(a)
}
