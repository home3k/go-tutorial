package main

import "fmt"

/*

Given a collection of candidate numbers (C) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

Each number in C may only be used once in the combination.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

 */
func combinationSum2(candidates []int, target int) [][]int {
	return _combSum(candidates, target, true)
}

func _combSum(candidates []int, target int, _sort bool) [][]int  {
	result:=[][]int{}
	if len(candidates) == 0 || (len(candidates) == 1 && candidates[0]!=target) {
		return result
	}
	if _sort {
		sort(candidates, 0, len(candidates)-1)
	}
	start, end := 0, len(candidates)
	for start < end {
		val:=candidates[start]
		if val>target {
			break
		}
		if start-1>=0 && candidates[start] == candidates[start-1] {
			start++
			continue
		}
		if val == target {
			result = append(result, []int{target})
			break
		} else {
			if start+1<end && candidates[start+1]<=target-val {
				next:=_combSum(candidates[start+1:], target - val, false)
				for _, ns:=range next {
					t:=[]int{val}
					t=append(t, ns...)
					result = append(result, t)
				}
			}
		}
		start++
	}
	return result
}

func sort(array []int, start int, end int) {
	if start >= end {
		return
	}
	pivot := array[(start+end)/2]
	left, right := start, end
	for left <= right {
		for left <= right && array[left] < pivot {
			left++
		}
		for left <= right && array[right] > pivot {
			right--
		}
		if left <= right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		}
	}
	sort(array, start, right)
	sort(array, left, end)
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

