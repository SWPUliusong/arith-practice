package main

import "fmt"

// 声明队列
type Queue []string

// 入队
func (q *Queue) push(strs ...string) {
	*q = append(*q, strs...)
}

// 出队
func (q *Queue) pop() string {
	if len(*q) > 0 {
		temp := (*q)[0]
		*q = (*q)[1:]
		return temp
	}
	return ""
}

// 选中条件
func is_seller(str string) bool {
	if(str[0:1] == "t") {
		return true
	}
	return false
}

// 广度优先
func bfs(graph map[string][]string, name string) string {
	search_queue := Queue(graph[name])
	searched := make(map[string]bool)
	for len(search_queue) > 0 {
		person := search_queue.pop()
		if _, exist := searched[person]; exist {
			continue
		}
		if is_seller(person) {
			return person
		} else {
			search_queue.push(graph[person]...)
			searched[person] = true
		}
	}
	return ""
}

func main() {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	fmt.Println(bfs(graph, "you"))
}