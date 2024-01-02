package main

import (
	"algo/min-max/minmax"
	"fmt"
)

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	result, _ := minmax.SumMinNElements(len(arr)-1, arr)
	fmt.Println(result)

	result, _ = minmax.SumMaxNElements(len(arr)-1, arr)
	fmt.Println(result)

}
