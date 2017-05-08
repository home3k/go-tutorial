package main

import "fmt"

/*

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place, do not allocate extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
1,2,3 → 1,3,2  213 231 312 321 123
3,2,1 → 1,2,3
1,1,5 → 1,5,1
2 1 3 - 2 3 1 - 3 1 2 - 3 2 1

1 2 3 2 1   1 3 1 2 2

*/

func nextPermutation(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}
	i := len(nums) - 1
	for i >= 1 {
		if nums[i] <= nums[i-1] {
			i--
		} else {
			j := len(nums) - 1
			for ; j >= i; j-- {
				if nums[j] > nums[i-1] {
					break
				}
			}
			nums[i-1], nums[j] = nums[j], nums[i-1]
			// sort
			reverse(nums, i, len(nums)-1)
			nti:= len(nums) -1 + i - j
			for nti>i {
				if nums[nti-1] > nums[nti] {
					nums[nti-1], nums[nti]=nums[nti], nums[nti-1]
					nti--
				} else {
					break
				}
			}
			return
		}
	}
	//reverse
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, start int, end int) {
	if start >= end {
		return
	}
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++;end--
	}
}

func main()  {
	nums:= []int{1,2,3,4,5}
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)
	nums=[]int{5,4,3,2,1}
	nextPermutation(nums)
	fmt.Println(nums)
}
