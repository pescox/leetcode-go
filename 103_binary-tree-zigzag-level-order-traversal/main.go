package main

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

// 解题思路：每层节点添加到队列中，添加每层节点后，将上层节点的值加入结果，如果是奇数层结果从左往右返回，如果是偶数层结果从右往左返回
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := []*TreeNode{root}
	var index int // 层级

	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			if index%2 == 0 {
				// 如果index是偶数，则从左边返回 追加到结果（第一层 index为0）
				tmp = append(tmp, queue[i].Val)
			} else {
				// index为奇数，则从右边返回，追加到结果
				tmp = append(tmp, queue[l-i-1].Val)
			}
		}

		index++
		queue = queue[l:]
		res = append(res, tmp)
	}
	return res
}
