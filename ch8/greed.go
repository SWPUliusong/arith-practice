package main

import "fmt"

type info map[string]float64

func main() {
	courses := map[string]info {
		"英语": info {
			"start": 9.30,
			"end": 10.30,
		},
		"美术": info {
			"start": 9.00,
			"end": 10.00,
		},
		"计算机": info {
			"start": 10.30,
			"end": 11.30,
		},
		"数学": info {
			"start": 10.00,
			"end": 11.00,
		},
		"音乐": info {
			"start": 11.00,
			"end": 12.00,
		},
	}

	fmt.Println(greed(courses))
}

func greed(courses map[string]info) []string {
	res := []string{}
	startTime := 0.0
	earliestEndTime := 24.0
	earliestEndCourse := ""

	for !isEmpty(courses) {
		for name, info := range courses {
			if info["start"] < startTime {
				delete(courses, name)
				continue
			}
			if info["end"] < earliestEndTime {
				earliestEndTime = info["end"]
				earliestEndCourse = name
			}

		}
		startTime, earliestEndTime = earliestEndTime, 24.0
		res = append(res, earliestEndCourse)
		delete(courses, earliestEndCourse)
	}

	return res
}

func isEmpty(m map[string]info) bool {
	for _ = range m {
		return false
	}
	return true
}