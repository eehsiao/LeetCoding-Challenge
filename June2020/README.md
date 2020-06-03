# 2020-June-LeetCoding-Challenge
[Challenge site](https://leetcode.com/explore/challenge/card/june-leetcoding-challenge/)

[My leetcode](https://leetcode.com/eehsiao/)


### Day1 : Invert Binary Tree

###  [Invert Binary Tree](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day1-invertBinaryTree.go)
```
Invert a binary tree.
Example:
Input:
     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:
     4
   /   \
  7     2
 / \   / \
9   6 3   1
Trivia:
This problem was inspired by this original tweet by Max Howell:
Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.
```

### Day2 : Delete Node in a Linked List

###  [Delete Node in a Linked List](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day2-deleteNodeinaLinkedList.go)
```
Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.
Given linked list -- head = [4,5,1,9], which looks like following:
Example 1:
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.
Example 2:
Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.
Note:
The linked list will have at least two elements.
All of the nodes' values will be unique.
The given node will not be the tail and it will always be a valid node of the linked list.
Do not return anything from your function.
```