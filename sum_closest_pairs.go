package coding

import (
	"math"
	"sort"
)

/*
	Given two integer arrays and a target which is a number.
	Find pairs of number that you can make from these two arrays whose sum is the closest to the target.

Example 1:
    Input:  [-1, 3, 8, 2, 9, 5], [4, 1, 2, 10, 5, 20]   24
	Output: [{3, 20}, {5, 20}]
*/

func SumClosestPairs(arr0 []int, arr1 []int, target int) [][]int {
	if len(arr0) == 0 || len(arr1) == 0 {
		return [][]int{}
	}
	var result [][]int
	distance := math.MaxFloat64
	sort.Ints(arr0)
	sort.Ints(arr1)
	for i, j := 0, len(arr1)-1; i < len(arr0) && j >= 0; {
		curDistance := float64(arr0[i] + arr1[j] - target)
		absCurDistance := math.Abs(curDistance)
		if absCurDistance < distance {
			result = [][]int{{arr0[i], arr1[j]}}
			distance = absCurDistance
		} else if absCurDistance == distance {
			result = append(result, []int{arr0[i], arr1[j]})
		}
		if curDistance >= 0 {
			j--
		} else {
			i++
		}
	}
	return result
}
