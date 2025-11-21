package main

import (
	"fmt"
	"sort"
)

func SortAndMerge(left, right []int) []int {
	sort.Ints(left)
	sort.Ints(right)

	res := make([]int, 0, len(left)+len(right))
	i := 0
	j := 0
	for i < len(left) {
		for j < len(right) {
			if left[i] < right[j] {
				res = append(res, left[i])
				i++
				break
			}
			if left[i] > right[j] {
				res = append(res, right[j])
				j++
				continue
			}
			if left[i] == right[j] {
				res = append(res, left[i])
				res = append(res, right[j])
				i++
				j++
				break
			}
		}
		if j == len(right) {
			res = append(res, left[i])
			i++
		}

	}

	for j < len(right) {
		res = append(res, right[j])
		j++
	}

	return res
}

func main() {
	left := []int{4, 1, 5, 0}
	right := []int{-1, 4, 5, 10}

	result := SortAndMerge(left, right)

	fmt.Println(result)

}
