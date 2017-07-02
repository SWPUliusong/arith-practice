package main

import "fmt"

func max(s []int) int {
	max := 0
	for _, val := range s {
		if val > max {
			max = val
		}
	}
	return max
}

func findWinMax(input []int, w int) []int {
	l := len(input)
	res := []int{}
	for i := range input {
		res = append(res, max(input[i:w+i]))
		if l == i+w {
			break
		}
	}
	return res
}

func main() {
	arr := []int{4, 3, 5, 4, 3, 3, 6, 7}
	fmt.Println(findWinMax(arr, 3))
}
