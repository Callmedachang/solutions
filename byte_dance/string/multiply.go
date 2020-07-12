package string
import (
	"log"
	"strings"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。*/
var siMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}
var isMap = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

func multiply(num1 string, num2 string) string {
	if num1=="0"||num2=="0"{
		return "0"
	}
	l2 := strings.Split(num2, "")
	res := "0"
	for i, v := range l2 {
		res=strAdd(res,strMul(num1, v, len(num2)-i-1))
		log.Println(strMul(num1, v, len(num2)-i-1))
		log.Println(res)
	}
	return res
}

func strAdd(num1 string, num2 string) string {
	s1 := strings.Split(num1, "")
	s2 := strings.Split(num2, "")
	res := ""
	l := 0
	if len(s1) > len(s2) {
		l = len(s2)
	} else {
		l = len(s1)
	}
	nextAdd := 0
	for i := 0; i < l; i++ {
		res = getIntStr((getStrInt(s1[len(s1)-1-i])+getStrInt(s2[len(s2)-1-i])+nextAdd)%10) + res
		nextAdd = (getStrInt(s1[len(s1)-1-i]) + getStrInt(s2[len(s2)-1-i])+nextAdd) / 10
	}
	if len(s1) > len(s2) {
		for i := len(s1) - len(s2) - 1; i >= 0; i-- {
			res = getIntStr((getStrInt(s1[i])+nextAdd)%10) + res
			nextAdd = (getStrInt(s1[i]) + nextAdd) / 10
		}
	} else if len(s1) < len(s2) {
		for i := len(s2) - len(s1) - 1; i >= 0; i-- {
			res = getIntStr((getStrInt(s2[i])+nextAdd)%10) + res
			nextAdd = (getStrInt(s2[i]) + nextAdd) / 10
		}
	} else {
		if nextAdd > 0 {
			res = getIntStr(nextAdd) + res
			nextAdd=0
		}
	}
	if nextAdd > 0 {
		res = getIntStr(nextAdd) + res
	}
	return res

}

func strMul(num1 string, num2 string, zeroNum int) string {
	bs := strings.Split(num1, "")
	ss := getStrInt(num2)
	res := ""
	nextAdd := 0
	for i := len(bs) - 1; i >= 0; i-- {
		res = getIntStr((getStrInt(bs[i])*ss+nextAdd)%10) + res
		nextAdd = (getStrInt(bs[i])*ss + nextAdd) / 10
	}
	if nextAdd > 0 {
		res = getIntStr(nextAdd) + res
	}
	for i := 0; i < zeroNum; i++ {
		res = res + "0"
	}
	return res
}

func getStrInt(s string) int {
	return siMap[s]
}

func getIntStr(s int) string {
	return isMap[s]
}
