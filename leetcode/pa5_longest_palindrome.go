package main

import "fmt"

// abc
func longestPalindrome(s string) string {
	if len(s) <=1 {
		return s
	}
	ns := "#"
	var rl []int = []int{1}
	for _, r := range s {
		ns += string(r) + "#"
		rl = append(rl,1,1)
	}
	var max, maxRight, pos, maxPos int
	for i := range ns {
		if i < maxRight {
			min := maxRight - i
			if rl[pos-(i-pos)] < min {
				min = rl[pos-(i-pos)]
			}
			rl[i] = min
		} else {
			rl[i] = 1
		}
		// check
		for i-rl[i] >= 0 && i+rl[i] < len(ns) && ns[i-rl[i]] == ns[i+rl[i]] {
			rl[i]++
		}
		if rl[i]+i-1 > maxRight {
			maxRight = rl[i] + i - 1
			pos = i
		}
		if max < rl[i]  {
			max = rl[i]
			maxPos = i
		}
	}
	ns = ns[maxPos - max + 1: maxPos + max -1]
	result:=""
	for _, r := range ns {
		if string(r)!="#" {
			result += string(r)
		}
	}
	return result
}

func main()  {
	fmt.Println(longestPalindrome("abcdc"))
	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("abcdsdcbsdeddf"))
}