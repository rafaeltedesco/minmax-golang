package main

import (
	"algo/min-max/minmax"
	"fmt"
)

func main() {
	result, err := minmax.SumMinNElements(10, []int32{10, 20, 30})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
