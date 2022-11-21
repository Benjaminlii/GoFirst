package sort

import (
	"context"

	"github.com/Benjaminlii/go_some_learning/common/logger"
)

// 快速排序
func QuickSort(array []int) {
	defer logger.TimeCost(context.Background())()
	quickSort(array, 0, len(array)-1)

}

func quickSort(array []int, startIndex, endIndex int) {
	if startIndex < endIndex {
		index := partition(array, startIndex, endIndex)
		go quickSort(array, startIndex, index-1)
		go quickSort(array, index+1, endIndex)
	}
}

func partition(array []int, startIndex, endIndex int) int {
	mark := array[startIndex]
	for startIndex < endIndex {
		for array[endIndex] >= mark && startIndex < endIndex {
			endIndex--
		}
		array[startIndex] = array[endIndex]
		for array[startIndex] <= mark  && startIndex < endIndex {
			startIndex++
		}
		array[endIndex] = array[startIndex]
	}
	array[startIndex] = mark
	return startIndex
}
