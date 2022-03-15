package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)

	//below is a map literal
	ages1 := map[string]int{
		"julie":  12,
		"juliet": 45}

	fmt.Println(ages)
	fmt.Println(ages1)
	fmt.Println(ages1["juliet"]) //map elements can be accessed as a subscript notation
	delete(ages1, "juliet")
	fmt.Println(ages1)
	ages["Babs"] += 1
	fmt.Println(ages["Babs"])

}
