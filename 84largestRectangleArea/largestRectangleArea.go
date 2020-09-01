package _4largestRectangleArea

import "log"

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

 



以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。

 



图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

 

示例:

输入: [2,1,5,6,2,3]
输出: 10
*/
func LargestRectangleArea(heights []int) int {
	res := 0
	for i := range heights {
		temp := FindMaxBefore(heights, i)
		if temp > res {
			res = temp
		}
	}
	return res
}

func FindMaxBefore(heights []int, index int) int {
	min, temp := heights[index], heights[index]
	for i := index; i > 0; i-- {
		if heights[i-1] < heights[i] {
			if heights[i-1] < min {
				min = heights[i-1]
			}
			if min*(index-i+2) == 0 {
				return temp
			} else {
				ss := min * (index - i + 2)
				if ss > temp {
					temp = ss
				}
			}
		} else {
			ss := min * (index - i + 2)
			if ss > temp {
				temp = ss
			}
		}
	}
	return temp
}

type Animal interface {
	Bark()
}
type Dog struct {
}

func (d *Dog) Bark() {

}

type _Human struct {
}
type HumanAndDog struct {
	Dog
	_Human
}

func (d *HumanAndDog) Bark() {
	log.Println(11)
}

func (d *HumanAndDog) cs() {
	log.Println(11)
}

func ss() {
	s := HumanAndDog{}
	s.Bark()
}
