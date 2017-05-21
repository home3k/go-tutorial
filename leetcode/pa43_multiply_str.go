package main

import "fmt"

/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2.

Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
213232323 
*  123
--
213232323 * 3 + 213232323 * 2 * 10 + 213232323 * 1 * 100
 */

var cache = make(map[int]string)

var tail = map[int]string {
	0:"",
	1:"0",
	2:"00",
	3:"000",
	4:"0000",
	5:"00000",
	6:"000000",
	7:"0000000",
	8:"00000000",
	9:"000000000",
}

func multiply(num1 string, num2 string) string {
	cache = make(map[int]string)
	l1, l2, l, r:=len(num1), len(num2), num1, num2
	if l1<l2 {
		l, r = num2, num1
	}
	result:=""
	base:=0
	for i:=len(r)-1;i>=0;i--{
		result=sum(result, _multi(l, int(r[i]-'0'), base))
		base++
	}
	return result

}

func getBase(x int) string {
	if _, ok := tail[x]; ok {
		return tail[x]
	} else {
		tail[x] = getBase(x-1) + "0"
		return tail[x]
	}
}

func _multi(num1 string, val int, base int) string {
	if mval,ok:=cache[val];ok {
		return mval + getBase(base)
	}
	if val == 0 {
		return "0"
	}
	if val == '1'{
		return num1 + getBase(base)
	}
	num:=val
	result:=""
	next:=0
	for i:=len(num1)-1;i>=0;i-- {
		v:=int(num1[i]-'0')
		x:=num*v+next
		next=x/10
		result=string(x%10+'0')+result
	}
	if next > 0 {
		result=string(next+'0')+result
	}
	cache[val]=result
	return result + getBase(base)
}

func sum(num1 string, num2 string) string {
	i1:=len(num1)-1
	i2:=len(num2)-1
	result:=""
	next:=0
	for i1>=0 && i2>=0 {
		v1:=int(num1[i1]-'0')
		v2:=int(num2[i2]-'0')
		x:=v1+v2+next
		next=x/10
		result=string(x%10+'0')+result
		i1--;i2--
	}
	if i1>=0 {
		for i1>=0 {
			v:=int(num1[i1]-'0')
			x:=v+next
			next=x/10
			result=string(x%10+'0')+result
			i1--
		}
	}
	if i2>=0 {
		for i2>=0 {
			v:=int(num2[i2]-'0')
			x:=v+next
			next=x/10
			result=string(x%10+'0')+result
			i2--
		}
	}
	if next > 0 {
		result=string(next+'0')+result
	}
	return result
}

func main()  {
	fmt.Println(multiply("0","0"))
	fmt.Println(multiply("1","23"))
	// 23234324234234234234
	fmt.Println(multiply("34234234234","4523456767768995678657867"))
	fmt.Println(multiply("9","9"))
	fmt.Println(multiply("401716832807512840963","167141802233061013023557397451289113296441069"))
}