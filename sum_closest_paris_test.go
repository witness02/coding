package coding

import (
	"fmt"
	"testing"
)

func TestSumClosestPairs(t *testing.T) {
	arr0 := []int{-1, 3, 8, 2, 9, 5, 21, 22}
	arr1 := []int{4, 1, 2, 10, 5, 20}
	fmt.Println(SumClosestPairs(arr0, arr1, 24))
}
