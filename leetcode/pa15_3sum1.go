package main

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	for i, num:= range nums  {
		if i+3== len(nums) {
			break
		}
		two:=twoSum(nums[i+1:], -num)
		for _, t:=range two {
			result = append(result, t)
		}

	}
}

func twoSum(nums []int, target int) [][]int {
	var result [][]int
	var m = make(map[int]int)
	for index, num := range nums {
		// match
		if val, ok := m[num]; ok {
			result = []int{val, index}
			break
		}
		m[target-num] = index
	}
	return result
}
