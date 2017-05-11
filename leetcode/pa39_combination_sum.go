package main

/*
Given a set of candidate numbers (C) (without duplicates) and a target number (T),
find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7,

A solution set is:
[
  [7],
  [2, 2, 3]
]
 */
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	result:=[][]int{}
	sort(candidates, 0, len(candidates)-1)
	if search(candidates, target, 0, len(candidates)) {
		result = append(result, []int{target})
	}
	start, end := 0, len(candidates)
	for start < end {

	}
}

func search(nums[]int, target int, start int, end int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0]==target
	}
	half:=len(nums)/2
	if nums[half] == target {
		return true
	} else if nums[half] > target {
		return search(nums, target, start, half)
	} else {
		return search(nums, target, half+1, end)
	}
}

func sort(nums[]int, start int, end int) {
	if start >= end {
		return
	}
	pivot := nums[(start+end)/2]
	left, right:=start, end
	for left <= right {
		for left <= right && nums[left] <= pivot  {
			left++
		}
		for left <= right && nums[right] >= pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++;right--
		}
	}
	sort(nums, start, right)
	sort(nums, left, end)
}
