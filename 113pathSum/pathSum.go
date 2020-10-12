package _13pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res := Dfs(root, []int{},[][]int{})
	result := make([][]int, 0)
	for _, v := range res {
		cur := 0
		for _, iv := range v {
			cur += iv
		}
		if cur == sum {
			result = append(result, v)
		}
	}
	return result
}
func copySlice(src []int) []int {
	res := make([]int, len(src))
	for i := range src {
		res[i] = src[i]
	}
	return res
}

func Dfs(root *TreeNode, path []int, resSet [][]int) [][]int {
	newPath := make([]int, len(path))
	if len(path) == 0 {
		newPath = []int{root.Val}
	} else {
		newPath = append(path, root.Val)
	}
	if root.Left == nil && root.Right == nil {
		resSet = append(resSet, newPath)
		return resSet
	}
	var leftPath, rightPath [][]int
	if root.Left != nil {
		leftPath = Dfs(root.Left, copySlice(newPath), leftPath)
	}
	if root.Right != nil {
		rightPath = Dfs(root.Right, copySlice(newPath), rightPath)
	}
	return append(leftPath, rightPath...)
}
