package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows<=1 {
		return s
	}
	m := make(map[int]string)
	positive :=true
	index := 0
	for _, r := range s {
		if positive {
			index++
		} else {
			index--
		}
		if index == numRows {
			positive = false
		}
		if index == 1 {
			positive = true
		}
		if _, ok:=m[index]; !ok {
			m[index]=""
		}
		m[index]+=string(r)
	}
	result :=""
	for i:=1; i<=numRows;i++ {
		result += m[i]
	}
	return result
}

func main()  {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 1))
	fmt.Println(convert("PAY", 4))
	fmt.Println(convert("", 1))
}