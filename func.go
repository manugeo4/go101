package main

import (
	"fmt"
	"sort"
)

func mapSort(m map[string]int) {
	names := make([]string, 0, len(m))

	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%v\n", name, m[name])
	}

}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true

}
