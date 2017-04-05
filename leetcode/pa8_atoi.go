package main

import "fmt"

const (
	INT_MAX = 2147483647
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	positive := true
	first := true
	result := 0
	for _, r := range str {
		if first {
			if r == ' ' || r == '\r' {
				continue
			}
			if r == '-' {
				positive = false
			} else if r == '+' || (r >= '0' && r <= '9') {
				positive = true
				if r != '+' {
					result = int(r - '0')
				}
			} else {
				return 0
			}
			first = false
		} else {
			if r >= '0' && r <= '9' {
				result = result*10 + int(r-'0')
				// overflow
				if positive && result >= INT_MAX  {
					return INT_MAX
				}
				if !positive && result >= (INT_MAX+1) {
					return -(INT_MAX + 1)
				}
			} else {
				break
			}
		}

	}
	if !positive {
		result = -result
	}
	return result
}

func main() {
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("-1043dx4444222323dasdfadfasdf00"))
	fmt.Println(myAtoi("dd"))
	fmt.Println(myAtoi("+100"))
	fmt.Println(myAtoi("         +13c00"))
	fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("2147483648"))
}
