package main

import "fmt"

// 最长公共子序列
func lcs(s1 string, s2 string) string {
	arr := make([][]int, len(s1))
	res := make([][]string, len(s1))
	for i := range arr {
		arr[i] = make([]int, len(s2))
		res[i] = make([]string, len(s2))
	}
	
	for i, t1 := range s1[:] {
		for j, t2 := range s2[:] {
			if t1 == t2 {
				if i==0 || j==0 {
					arr[i][j] = 1
					res[i][j] = string(t1)
					continue
				}
					
				arr[i][j] = arr[i-1][j-1] + 1
				res[i][j] = res[i-1][j-1] + string(t1)
			} else {
				if i==0 || j==0 {
					continue
				}

				if arr[i][j-1] > arr[i-1][j] {
					arr[i][j] = arr[i][j-1]
					res[i][j] += res[i][j-1]
				} else {
					arr[i][j] = arr[i-1][j]
					res[i][j] += res[i-1][j]
				}
			}
		}
	}
	return res[len(s1) - 1][len(s2) - 1]
}

func main() {
	fmt.Println(lcs("blue", "clues"))
}