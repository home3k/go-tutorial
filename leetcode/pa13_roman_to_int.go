package main

import "fmt"

var romanMap = map[string]int{"M":1000,"CM":900,"D":500,"CD":400,"C":100,"XC":90,"L":50,"XL":40,"X":10,"IX":9,"V":5,"IV":4,"I":1}

func romanToInt(s string) int {
	result:=0
	l:=len(s)
	for i:=0;i<l;{
		r:=string(s[i])
		step:=1
		c:=romanMap[string(r)]
		if i+1 < l {
			r2:=string(s[i])+string(s[i+1])
			if c2, ok:=romanMap[r2];ok {
				c=c2
				step=2
			}
		}
		result+=c
		i+=step
	}
	return result
}

func main()  {
	fmt.Println(romanToInt("I"))
	fmt.Println(romanToInt("XL"))
	fmt.Println(romanToInt("CM"))
	fmt.Println(romanToInt("DCXXI"))
}


