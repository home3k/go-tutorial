package main

import "fmt"

/*

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.
[-2,1,-3,4,-1,2,1,-5,4]
 -1 1

 */


const (
	INT_MIN = -2147483648
)


func maxSubArray(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	max, sum, current, next:=INT_MIN, 0, 0, -1
	for i:=0;i<len(nums); {
		v:=nums[i]
		sum=v;current=i;next=-1
		if sum < max {
			continue
		} else {
			max = sum
		}
		j:=i+1
		for j<len(nums) {
			vj:=nums[j]
			if vj>=0 {
				sum += vj
				max = sum
				if vj > nums[current] {
					if (next >=0 && nums[next]<vj) || next<0{
						next = j
					}
				}
				j++
			} else {
				if next>=0 {
					i=next
				} else {
					i=j+1
				}
				break
			}
		}

	}
	return max
}

func main()  {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
