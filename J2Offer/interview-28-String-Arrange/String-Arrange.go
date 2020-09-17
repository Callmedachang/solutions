package interview_28_String_Arrange

import (
	"strings"
)

/*
题目：输入一个字符串，打印出该字符串中字符的所有排列。
例如输入字符串abc，则打印出由字符a、b、c所能排列出来的所有字符串abc、acb、bac、bca、cab和cba
*/

/*
递归解析：
终止条件： 当 x = len(c) - 1 时，代表所有位已固定（最后一位只有 1 种情况），
则将当前组合 c 转化为字符串并加入 res，并返回；
递推参数： 当前固定位 xx ；
递推工作： 初始化一个 Set，用于排除重复的字符；将第 x 位字符与 i∈[x,len(c)] 字符分别交换，并进入下层递归；
	1.剪枝：若 c[i] 在 Set​ 中，代表其是重复字符，因此“剪枝”；
	2.将 c[i] 加入 Set​ ，以便之后遇到重复字符时剪枝；
	3.固定字符： 将字符 c[i] 和 c[x] 交换，即固定 c[i] 为当前位字符；
	4.开启下层递归： 调用 dfs(x + 1) ，即开始固定第 x + 1 个字符；
	5.还原交换： 将字符 c[i] 和 c[x] 交换（还原之前的交换）；

*/
func StringArrange(str string) []string {
	return dfs(strings.Split(str, ""), 0, []string{})
}

func dfs(s []string, index int, res []string) []string {
	if index == len(s)-1 {
		return append(res, strings.Join(s, ""))
	}
	deDup := make(map[string]int)
	for i := index; i < len(s); i++ {
		if deDup[s[i]] == 1 {
			continue
		}
		deDup[s[i]] = 1
		s[i], s[index] = s[index], s[i]
		res = dfs(s, index+1, res)
		s[i], s[index] = s[index], s[i]
	}
	return res
}