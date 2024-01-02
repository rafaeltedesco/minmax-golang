package minmax

import "errors"

type minMaxEvaluator func(int32, int32, int, int) (int32, int)

var invalidLength = errors.New("invalid length")

func minEvaluator(lowestEl, nextEl int32, lowestPos int, nextPos int) (int32, int) {
	if nextEl < lowestEl {
		lowestEl = nextEl
		lowestPos = nextPos
	}
	return lowestEl, lowestPos
}

func maxEvaluator(lowestEl, nextEl int32, lowestPos int, nextPos int) (int32, int) {
	if nextEl > lowestEl {
		lowestEl = nextEl
		lowestPos = nextPos
	}
	return lowestEl, lowestPos
}

func GetEl(arr []int32, fn minMaxEvaluator) (int32, int) {
	currentElement := arr[0]
	position := 0
	for idx, el := range arr {
		currentElement, position = fn(currentElement, el, position, idx)
	}
	return currentElement, position
}

func getLowest(arr []int32) (int32, int) {
	return GetEl(arr, minEvaluator)
}

func getHighest(arr []int32) (int32, int) {
	return GetEl(arr, maxEvaluator)
}

type sort func(int32, int32) (int32, int32)

func Sort(arr []int32, fn sort) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			arr[i], arr[j] = ascending(arr[i], arr[j])
		}
	}
	return arr
}

func ascending(el1, el2 int32) (int32, int32) {
	if el1-el2 > 0 {
		return el2, el1
	} else if el1-el2 < 0 {
		return el1, el2
	} else {
		return el1, el2
	}
}

func GetNLowestElementsFromArray(n int, arr []int32) []int32 {
	var newArr []int32
	tempArr := arr

	for len(newArr) < n {
		el, pos := getLowest(tempArr)
		newArr = append(newArr, el)
		tempArr = append(tempArr[:pos], tempArr[pos+1:]...)
	}
	return newArr
}

func GetNHighestElementsFromArray(n int, arr []int32) []int32 {
	var newArr []int32
	tempArr := arr

	for len(newArr) < n {
		el, pos := getHighest(tempArr)
		newArr = append(newArr, el)
		tempArr = append(tempArr[:pos], tempArr[pos+1:]...)
	}
	return Sort(newArr, ascending)
}

func SumMinNElements(n int, arr []int32) (int32, error) {
	if n >= len(arr) {
		return -1, invalidLength
	}
	elements := GetNLowestElementsFromArray(n, arr)
	var total int32 = 0
	for _, el := range elements {
		total += el
	}
	return total, nil
}

func SumMaxNElements(n int, arr []int32) (int32, error) {
	if n >= len(arr) {
		return -1, invalidLength
	}
	elements := GetNHighestElementsFromArray(n, arr)
	var total int32 = 0
	for _, el := range elements {
		total += el
	}
	return total, nil
}
