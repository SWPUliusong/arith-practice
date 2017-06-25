package main

import "fmt"

func main() {
	arr := []int{8,4,3,6,7,9,2,1,5,7,4}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	l := len(arr)

	// 基线条件
	if l < 2 {
		return arr
	}

	// 基准值
	pivot := arr[0]

	left := make([]int, 0, l)
	right := make([]int, 0, l)

	for _, val := range arr[1:] {
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	return merge(quickSort(left), []int{pivot}, quickSort(right))
}

// 合并slice
func merge(arrs ...[]int) []int {
	res := []int{}
	for _, arr := range arrs {
		for _, val := range arr {
			res = append(res, val)
		}
	}
	return res
}