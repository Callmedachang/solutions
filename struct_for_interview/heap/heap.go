package main

import "log"

/*
一般情况下，堆通常指的是二叉堆，二叉堆是一个近似完全二叉树的数据结构，但由于对二叉树平衡及插入/删除操作较为麻烦，二叉堆实际上使用数组来实现。即物理结构为数组，逻辑结构为完全二叉树。子结点的键值或索引总是小于（或者大于）它的父节点，且每个节点的左右子树又是一个二叉堆(大根堆或者小根堆)。根节点最大的堆叫做最大堆或大根堆，根节点最小的堆叫做最小堆或小根堆。常被用作实现优先队列。

特点
	以数组表示，但是以完全二叉树的方式理解。
	唯一能够同时最优地利用空间和时间的方法——最坏情况下也能保证使用 2N \log N2NlogN 次比较和恒定的额外空间。
	在索引从0开始的数组中：
	父节点 i 的左子节点在位置(2*i+1)
	父节点 i 的右子节点在位置(2*i+2)
	子节点 i 的父节点在位置floor((i-1)/2)
堆的基本操作(以大根堆为例，堆的常用操作如下。)
	最大堆调整：将堆的末端子节点作调整，使得子节点永远小于父节点
	创建最大堆：将堆所有数据重新排序
	堆排序：移除位于第一个数据的根节点，并做最大堆调整的递归运算
*/

func heapSort(source []int) []int {
	res, lenSource := make([]int, 0), len(source)
	for i := 0; i < lenSource; i++ {
		//log.Println(source)
		afterAdjust := adjustHeap(source)
		log.Println(afterAdjust)
		res = append(res, afterAdjust[0])
		source = afterAdjust[1:]
	}
	return res
}

func adjustHeap(source []int) []int {
	i, lenSource, heap, res, ok := 0, len(source), make([][]int, 0), make([]int, 0), false
	for {
		rootN := make([]int, 0)
		if 2<<uint(i)-1 > lenSource {
			if i == 0 {
				rootN = source[0:]
			} else {
				rootN = source[2<<(uint(i - 1))-1:]
			}
			heap = append(heap, rootN)
			break
		} else {
			if i == 0 {
				rootN = source[0:1]
			} else {
				rootN = source[2<<(uint(i - 1))-1 : 2<<(uint(i))-1]
			}
			heap = append(heap, rootN)
			i++
		}
	}
	for !ok {
		heap, ok = adjust(heap)
	}
	for _, c := range heap {
		for _, n := range c {
			res = append(res, n)
		}
	}
	return res
}

func adjust(source [][]int) ([][]int, bool) {
	lenSource := len(source)
	for i := 0; i < lenSource-1; i++ {
		roots := source[i]
		leaves := source[i+1]
		for j := 0; j < len(roots); j++ {
			if len(leaves) >= 2*(j+1) {
				if roots[j] < leaves[2*j] || roots[j] < leaves[2*j+1] {
					if leaves[2*j] > leaves[2*j+1] {
						source[i][j], source[i+1][2*j] = source[i+1][2*j], source[i][j]
						return source, false
					} else {
						source[i][j], source[i+1][2*j+1] = source[i+1][2*j+1], source[i][j]
						return source, false
					}
				}
			} else if len(leaves) > 2*j && len(leaves) < 2*(j+1) {
				if roots[j] < leaves[2*j] {
					source[i][j], source[i+1][2*j] = source[i+1][2*j], source[i][j]
					return source, false
				} else {
					return source, true
				}
			} else {
				return source, true
			}
		}
	}
	return source, true
}
