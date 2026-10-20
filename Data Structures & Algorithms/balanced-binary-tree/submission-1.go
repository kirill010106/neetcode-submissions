/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	balanced := true
    if root == nil {
		return true
	}

	var height func(node *TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		right := height(node.Right)

		if math.Abs(float64(right-left)) > 1 {
			balanced = false
		}

		return max(left,right)+1
	}

	height(root)
	return balanced
}
