# 2020-July-LeetCoding-Challenge
[Challenge site](https://leetcode.com/explore/challenge/card/july-leetcoding-challenge/)

[My leetcode](https://leetcode.com/eehsiao/)


### Day1 : Arranging Coins

###  [Arranging Coins](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day1-arrangingCoins.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day1-arrangingCoins_test.go)
```
You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.
Given n, find the total number of full staircase rows that can be formed.
n is a non-negative integer and fits within the range of a 32-bit signed integer.
Example 1:
n = 5
The coins can form the following rows:
¤
¤ ¤
¤ ¤
Because the 3rd row is incomplete, we return 2.
Example 2:
n = 8
The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤
Because the 4th row is incomplete, we return 3.
```


### Day2 : Arranging Coins

###  [Arranging Coins](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day2-binaryTreeLevelOrderTraversalII.go)
```
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
```


### Day7 : Island Perimeter

###  [Island Perimeter](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day7-islandPerimeter.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day7-islandPerimeter_test.go)
```
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.
Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.
Example:
Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]
Output: 16
```


### Day8 : 3Sum

###  [3Sum](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day8-3Sum.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day8-3Sum_test.go)
```
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
Note:
The solution set must not contain duplicate triplets.
Example:
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
   Hide Hint #1
So, we essentially need to find three numbers x, y, and z such that they add up to the given value. If we fix one of the numbers say x, we are left with the two-sum problem at hand!
   Hide Hint #2
For the two-sum problem, if we fix one of the numbers, say
x
, we have to scan the entire array to find the next number
y
which is
value - x
where value is the input parameter. Can we change our array somehow so that this search becomes faster?
   Hide Hint #3
The second train of thought for two-sum is, without changing the array, can we use additional space somehow? Like maybe a hash map to speed up the search?

```


### Day9 : Maximum Width of Binary Tree

###  [Maximum Width of Binary Tree](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day9-maximumWidthofBinaryTree.go)
```
Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a full binary tree, but some nodes are null.
The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the null nodes between the end-nodes are also counted into the length calculation.
Example 1:
Input:
           1
         /   \
        3     2
       / \     \
      5   3     9
Output: 4
Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
Example 2:
Input:
          1
         /
        3
       / \
      5   3
Output: 2
Explanation: The maximum width existing in the third level with the length 2 (5,3).
Example 3:
Input:
          1
         / \
        3   2
       /
      5
Output: 2
Explanation: The maximum width existing in the second level with the length 2 (3,2).
Example 4:
Input:
          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
Output: 8
Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).
Note: Answer will in the range of 32-bit signed integer.

```