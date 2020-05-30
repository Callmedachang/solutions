package radixSort

import (
	"strconv"
)

func RadixSort(arr []int) []int{
	if len(arr)<2{
		return arr
	}
	maxl:=MaxLen(arr)
	return RadixCore(arr,0,maxl)
}
func RadixCore(arr []int,digit,maxl int) []int{   //核心排序机制时间复杂度 O( d( r+n ) )
	if digit>=maxl{
		return arr                                               //排序稳定
	}
	radix:=10
	count:=make([]int,radix)
	bucket:=make([]int,len(arr))
	for i:=0;i<len(arr);i++{
		count[GetDigit(arr[i],digit)]++
	}
	for i:=1;i<radix;i++{
		count[i]+=count[i-1]
	}
	for i:=len(arr)-1;i>=0;i--{
		d:=GetDigit(arr[i],digit)
		bucket[count[d]-1]=arr[i]
		count[d]--
	}
	return RadixCore(bucket,digit+1,maxl)
}

func GetDigit(x,d int) int{                          //获取某位上的数字
	a:=[]int{1,10,100,1000,10000,100000,1000000}
	return (x/a[d])%10
}

func MaxLen(arr []int) int{                 //获取最大位数
	var maxl,curl int
	for i:=0;i<len(arr);i++{
		curl=len(strconv.Itoa(arr[i]))
		if curl>maxl{
			maxl=curl
		}
	}
	return maxl
}