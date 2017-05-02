package main

import "fmt"

/*
 kmp
*/
func strStr(haystack string, needle string) int {
	i, j, hl, nl := 0, 0, len(haystack), len(needle)
	if nl == 0 {
		return 0
	} else if hl == 0 {
		return -1
	}
	next := make([]int, nl, nl)

	getNext(needle, next)
	for i<hl && j<nl {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if next[j] == -1 {
				i++
				j=0
			} else {
				j = next[j]
			}
		}
		if j == nl {
			return i-j
		}
	}

	return -1

}

func getNext(needle string, next []int) {
	j, k, l := 0, -1, len(needle)
	next[0] = -1
	for j < l -1  {
		if k == -1 || needle[k] == needle[j] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
}

func main() {
	fmt.Println(strStr("asdferwfasdfasd","dfas"))
}
