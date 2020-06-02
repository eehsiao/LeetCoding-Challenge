// Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
// Note:
// You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
// Example 1:
// Input: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// Output: 1
// Example 2:
// Input: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1
// Output: 3
// Follow up:
// What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?
//    Hide Hint #1
// Try to utilize the property of a BST.
//    Hide Hint #2
// Try in-order traversal. (Credits to @chan13)
//    Hide Hint #3
// What if you could modify the BST node's structure?
//    Hide Hint #4
// The optimal runtime complexity is O(height of BST).

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k--; k == 0 {
			return root.Val
		}
		root = root.Right
	}

}
