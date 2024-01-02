package minmax

type minMaxEvaluator func(int32, int32, int, int) (int32, int)

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
