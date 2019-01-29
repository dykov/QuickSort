package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var array = []int{10,20,2,101,5,-1}
	fmt.Println(QuickSort(array))
}

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left := 0
	right := len(array)-1
	pivot := rand.Int() % len(array) // random pivot element

	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}

	array[left], array[right] = array[right], array[left]

	QuickSort(array[:left])
	QuickSort(array[left+1:])

	return array
}