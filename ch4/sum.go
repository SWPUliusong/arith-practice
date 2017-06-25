package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6}

	fmt.Println(sum(arr))
}

func sum(arr []int) int {
	if (len(arr) == 1) {
		return arr[0]
	}

	return arr[0] + sum(arr[1:])
}