package main

import (
	"fmt"
	"math"
)

type M map[string]int
// 已检测的node
var processed = []string{}

func main() {
	// 图形结构
	graph := map[string]M{
		"start": M{
			"a": 6,
			"b": 2,
		},
		"a": M{
			"fin": 1,
		},
		"b": M{
			"a":   3,
			"fin": 5,
		},
	}

	// 花费
	costs := M{
		"a": 6,
		"b": 2,
	}

	// 父节点
	parents := map[string]string{
		"a": "start",
		"b": "start",
	}

	node := find_lowest_cost_node(costs)
	for !has_elem(processed, node) && node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n, c := range neighbors {
			new_cost := c + cost
			if old_cost, flag := costs[n]; !flag || new_cost < old_cost {
				costs[n] = new_cost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = find_lowest_cost_node(costs)
	}

	fmt.Println(parents)
	fmt.Println(costs["fin"])
}

// 找出costs中最小花费的node
func find_lowest_cost_node(costs map[string]int) string {
	lowest_cost := int(math.Pow(2, 31) - 1)
	lowest_cost_node := ""
	for node, cost := range costs {
		if cost < lowest_cost && !has_elem(processed, node) {
			lowest_cost = cost
			lowest_cost_node = node
		}
	}
	return lowest_cost_node
}

// 判断元素是否存在
func has_elem(s []string, elem string) bool {
	for _, val := range s {
		if elem == val {
			return true
		}
	}
	return false
}
