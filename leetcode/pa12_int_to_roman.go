package main

import "fmt"

/*
I = 1
V = 5
X = 10
L = 50
C = 100
D = 500
M = 1000
 */

var it = []int {1000,900,500,400,100,90,50,40,10,9,5,4,1}
var roman = []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
func intToRoman(num int) string {
	result:=""
	for i,item:=range it {
		for num >= item {
			result+=roman[i]
			num-=item
		}
	}
	return result
}

func main()  {
	fmt.Println(intToRoman(3999))
	fmt.Println(intToRoman(1))
	fmt.Println(intToRoman(1000))
	fmt.Println(intToRoman(9))
}
