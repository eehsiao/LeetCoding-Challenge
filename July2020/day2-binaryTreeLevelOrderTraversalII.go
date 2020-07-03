// Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var levels [][]int

func levelOrder(node *TreeNode, level int) {
	if len(levels) == level {
		levels = append(levels, make([]int, 0))
	}
	levels[level] = append(levels[level], node.Val)

	if node.Left != nil {
		levelOrder(node.Left, level+1)
	}

	if node.Right != nil {
		levelOrder(node.Right, level+1)
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	levels = make([][]int, 0)
	results := make([][]int, 0)
	if root != nil {
		levelOrder(root, 0)
	}

	//reverse
	for _, l := range levels {
		results = append([][]int{l}, results...)
	}

	return results
}
