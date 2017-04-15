package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r, c := 0, len(nums)-1, 0
	for l <= r {
		if nums[l] == val {
			c++
			for nums[r] == val && l < r {
				c++
				r--
			}
			nums[l], nums[r] = nums[r], nums[l]
			r--
			l++
		} else {
			l++
		}
	}
	return len(nums) - c
}

func main() {
	n1 := []int{3, 2, 2, 3}
	fmt.Println(removeElement(n1, 3), n1)
	n1 = []int{3, 3, 3, 3}
	fmt.Println(removeElement(n1, 3), n1)
	n1 = []int{}
	fmt.Println(removeElement(n1, 3), n1)
	n1 = []int{3, 2, 3, 3}
	fmt.Println(removeElement(n1, 3), n1)
	n1 = []int{4,5}
	fmt.Println(removeElement(n1, 5), n1)
	n1 = []int{1}
	fmt.Println(removeElement(n1, 5), n1)
	n1 = []int{1}
	fmt.Println(removeElement(n1, 1), n1)
}
