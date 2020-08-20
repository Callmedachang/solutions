package interview_31_Couter_Of_1

import (
	"math"
	"strconv"
)

/*
题目：输入一个整数n，求从1到n这n个整数的十进制表示中1出现的次数。
例如输入12，从1到12这些整数中包含1的数字有1，10，11和12，1一共出现了5次。
*/
///思路
// if >1   =1 else 0   1位数
// if >20 n*1 if >2 +10
//        n*10 if >2 +100

func CounterOf1(num int) int {
	length := len(strconv.Itoa(num))
	p := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		temp := num / (int(math.Pow10(i)))
		p = append(p, temp)
		num = num - temp*int(math.Pow10(i))
	}
	res := 0

	for i, v := range p {
		if v == 1 {
			res += getNum(p[i+1:]) + 1
		} else {
			s1 := int(math.Pow10(len(p)-i-1))
			if v == 0 {
				v = 10
			}
			if len(p)-i-2<0{
				res += 1 + s1
			}else{
				res += (v-2)*int(math.Pow10(len(p)-i-2)) + s1
			}
		}
	}
	return res
}

func getNum(s []int) int {
	//123
	res := 0
	for i, v := range s {
		res += v * (int(math.Pow10(len(s) - i - 1)))
	}
	if res==0{
		res=1
	}
	return res
}
