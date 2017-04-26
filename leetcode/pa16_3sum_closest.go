package main

import (
	"fmt"
	"math"
)

/*
   For example, given array S = {-1 2 1 -4}, and target = 1.
   The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest(nums []int, target int) int {
	if nums == nil {
		return 0
	}
	sort(nums, 0, len(nums)-1)
	if len(nums) < 3 {
		sum := 0
		for _, n := range nums {
			sum += n
		}
		return sum
	}
	min := -1
	val := 0
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			gap := int(math.Abs(float64(sum - target)))
			if gap < min || min < 0 {
				min = gap
				val = sum
			}
			if sum == target {
				return target
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return val
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
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{-2, 2, 1, -4}, 1))
}
