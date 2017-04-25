package main
/*

[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

123 132 213 231 3

 */
func permute(nums []int) [][]int {

}

func sort(nums []int, left int, right int)  {

	if left >= right {
		return
	}
	pivot := nums[(left+right)/2]
	for left <= right {
		if left <= right && nums[left] < pivot {
			left++
		}
		if left <= right && nums[right] > pivot{
			right--
		}
		if left <= right {
			nums[left], nums[right]= nums[right], nums[left]
		}
	}

	sort(nums, 0, left)
}

/*
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
 */
func next(nums[]int) []int {
	i:=len(nums)
	for i>0 {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		} else {

		}
	}
}
