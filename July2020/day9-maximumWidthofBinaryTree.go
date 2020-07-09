// Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a full binary tree, but some nodes are null.
// The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the null nodes between the end-nodes are also counted into the length calculation.
// Example 1:
// Input:
//            1
//          /   \
//         3     2
//        / \     \
//       5   3     9
// Output: 4
// Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
// Example 2:
// Input:
//           1
//          /
//         3
//        / \
//       5   3
// Output: 2
// Explanation: The maximum width existing in the third level with the length 2 (5,3).
// Example 3:
// Input:
//           1
//          / \
//         3   2
//        /
//       5
// Output: 2
// Explanation: The maximum width existing in the second level with the length 2 (3,2).
// Example 4:
// Input:
//           1
//          / \
//         3   2
//        /     \
//       5       9
//      /         \
//     6           7
// Output: 8
// Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).
// Note: Answer will in the range of 32-bit signed integer.

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	maxWidth       int
	firstColIdxMap map[int]int
)

func dfs(node *TreeNode, depth, colIndex int) {
	var (
		firstColIdx int
		ok          bool
	)
	if node == nil {
		return
	}
	if firstColIdx, ok = firstColIdxMap[depth]; !ok {
		firstColIdxMap[depth] = colIndex
		firstColIdx = colIndex
	}

	maxWidth = max(maxWidth, colIndex-firstColIdx+1)
	dfs(node.Left, depth+1, 2*colIndex)
	dfs(node.Right, depth+1, 2*colIndex+1)

}

func widthOfBinaryTree(root *TreeNode) int {
	maxWidth, firstColIdxMap = 0, make(map[int]int)
	dfs(root, 0, 0)

	return maxWidth
}
