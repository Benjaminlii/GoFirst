package sort

import (
	"math/rand"
	"testing"
	"time"
)

func createSortReq(length, start, end int) []int {
	res := make([]int, 0, length)

	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		random := rand.Int()
		random %= end - start
		random += start
		res = append(res, random)
	}

	return res
}
func TestQuickSort(t *testing.T) {
	req := createSortReq(500000, 0, 500000)
	// fmt.Println(req)
	QuickSort(req)
	// fmt.Println(req)
}
