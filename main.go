package main

import (
	"fmt"

	"github.com/manugeo4/gd"
	"github.com/manugeo4/gf"
)

func main() {

	gf.MapSort(gd.Ages)
	fmt.Println(gf.MapEqual(gd.Ages, gd.Ages1))

}
