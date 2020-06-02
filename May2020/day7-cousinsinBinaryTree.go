// In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
// Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
// We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
// Return true if and only if the nodes corresponding to the values x and y are cousins.
// Example 1:
//     1
//    / \
//   2   3
//  /
// 4
// Input: root = [1,2,3,4], x = 4, y = 3
// Output: false
// Example 2:
//     1
//    / \
//   2   3
//    \   \
//     4   5
// Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
// Output: true
// Example 3:
//     1
//    / \
//   2   3
//    \
//     4
// Input: root = [1,2,3,null,4], x = 2, y = 3
// Output: false
// Note:
// The number of nodes in the tree will be between 2 and 100.
// Each node has a unique integer value from 1 to 100.

package main

func findNode(parent *TreeNode, node *TreeNode, v int, depth int) (p *TreeNode, d int) {
	if node.Val == v {
		return parent, depth
	}

	if node.Left != nil {
		if p1, d1 := findNode(node, node.Left, v, depth+1); p1 != nil {
			return p1, d1
		}
	}
	if node.Right != nil {
		if p1, d1 := findNode(node, node.Right, v, depth+1); p1 != nil {
			return p1, d1
		}
	}

	return nil, 0
}

func isCousins(root *TreeNode, x int, y int) bool {

	p1, d1 := findNode(nil, root, x, 0)
	p2, d2 := findNode(nil, root, y, 0)

	if p1 != p2 && d1 == d2 {
		return true
	}

	return false
}
