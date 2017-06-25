package main

import "fmt"

func SelectionSort(arr []int) []int {
	var e int
	for f := range arr {
		e = f + 1
		for ;e < len(arr); e++ {
			if arr[f] > arr[e] {
				arr[f], arr[e] = arr[e], arr[f]
				continue
			}
		}
	}
	return arr
}

func main() {
	arr := []int{8,5,4,7,3,6,1,4,7}

	fmt.Println(SelectionSort(arr))
}