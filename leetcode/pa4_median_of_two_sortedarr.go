package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	if l == 0 {
		return 0.0
	}
	var merge []int
	even := false
	median := l/2 + 1
	if l%2 == 0 {
		even = true
	}
	var i1, i2, count, v1, v2 int
	f1 := false
	f2 := false
	for index := 0; index < l; index++ {
		if count >= median {
			break
		}
		if i1 < l1 {
			v1 = nums1[i1]
		} else {
			f1 = true
		}
		if i2 < l2 {
			v2 = nums2[i2]
		} else {
			f2 = true
		}
		if !f1 && !f2 {
			if v1 <= v2 {
				merge = append(merge, v1)
				i1++
			} else {
				merge = append(merge, v2)
				i2++
			}
		} else if !f1 {
			merge = append(merge, v1)
			i1++
		} else {
			merge = append(merge, v2)
			i2++
		}
		count++
	}
	if even {
		return float64(merge[len(merge)-1]+merge[len(merge)-2]) / 2.0
	} else {
		return float64(merge[len(merge)-1])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{}))
	fmt.Println(findMedianSortedArrays([]int{1}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{}))
}
