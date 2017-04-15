package main

import "fmt"

func firstMissingPositive(nums []int) int {
	i:=0
	for i<len(nums) {
		num:=nums[i]
		if num>0 && num<=len(nums) && num!=i+1 && nums[num-1]!=num{
			nums[num-1], nums[i]=nums[i],nums[num-1]
		} else {
			i++
		}
	}

	for i, num:=range nums {
		if i+1!=num {
			return i+1
		}
	}
	return len(nums)+1
}

func main()  {
	fmt.Println(firstMissingPositive([]int{1,2,3,4}))
	fmt.Println(firstMissingPositive([]int{1,2,0}))
	fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
	fmt.Println(firstMissingPositive([]int{2,1}))
}
