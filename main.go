package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"shiva":  39,
		"julie":  12,
		"juliet": 45}
	ages1 := map[string]int{
		"shiva":  39,
		"julie":  12,
		"juliet": 44}
	mapSort(ages)
	fmt.Println(mapEqual(ages, ages1))

}

/////////end of main
