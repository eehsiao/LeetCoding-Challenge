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


### Day3 : Prison Cells After N Days

###  [Prison Cells After N Days](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day3-prisonCellsAfterNDays.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day3-prisonCellsAfterNDays_test.go)
```
There are 8 prison cells in a row, and each cell is either occupied or vacant.
Each day, whether the cell is occupied or vacant changes according to the following rules:
If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell becomes occupied.
Otherwise, it becomes vacant.
(Note that because the prison is a row, the first and the last cells in the row can't have two adjacent neighbors.)
We describe the current state of the prison in the following way: cells[i] == 1 if the i-th cell is occupied, else cells[i] == 0.
Given the initial state of the prison, return the state of the prison after N days (and N such changes described above.)
Example 1:
Input: cells = [0,1,0,1,1,0,0,1], N = 7
Output: [0,0,1,1,0,0,0,0]
Explanation:
The following table summarizes the state of the prison on each day:
Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
Day 7: [0, 0, 1, 1, 0, 0, 0, 0]
Example 2:
Input: cells = [1,0,0,1,0,0,1,0], N = 1000000000
Output: [0,0,1,1,1,1,1,0]
Note:
cells.length == 8
cells[i] is in {0, 1}
1 <= N <= 10^9
```


### Day4 : Ugly Number II

###  [Ugly Number II](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day4-uglyNumberII.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day4-uglyNumberII_test.go)
```
Write a program to find the n-th ugly number.
Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
Example:
Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
Note:
1 is typically treated as an ugly number.
n does not exceed 1690.
   Hide Hint #1
The naive approach is to call isUgly for every number until you reach the nth one. Most numbers are not ugly. Try to focus your effort on generating only the ugly ones.
   Hide Hint #2
An ugly number must be multiplied by either 2, 3, or 5 from a smaller ugly number.
   Hide Hint #3
The key is how to maintain the order of the ugly numbers. Try a similar approach of merging from three sorted lists: L1, L2, and L3.
   Hide Hint #4
Assume you have Uk, the kth ugly number. Then Uk+1 must be Min(L1 * 2, L2 * 3, L3 * 5).
```


### Day5 : Hamming Distance

###  [Hamming Distance](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day5-hammingDistance.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day5-hammingDistance_test.go)
```
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, calculate the Hamming distance.
Note:
0 ≤ x, y < 231.
Example:
Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
```


### Day6 : Plus One

###  [Plus One](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day6-plusOne.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day6-plusOne_test.go)
```
Given a non-empty array of digits representing a non-negative integer, increment one to the integer.
The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.
You may assume the integer does not contain any leading zero, except the number 0 itself.
Example 1:
Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:
Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
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


### Day10 : Flatten a Multilevel Doubly Linked List

###  [Flatten a Multilevel Doubly Linked List](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day10-flattenaMultilevelDoublyLinkedList.go)
```
ou are given a doubly linked list which in addition to the next and previous pointers, it could have a child pointer, which may or may not point to a separate doubly linked list. These child lists may have one or more children of their own, and so on, to produce a multilevel data structure, as shown in the example below.
Flatten the list so that all the nodes appear in a single-level, doubly linked list. You are given the head of the first level of the list.
Example 1:
Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
Output: [1,2,3,7,8,11,12,9,10,4,5,6]
Explanation:
The multilevel linked list in the input is as follows:
https://assets.leetcode.com/uploads/2018/10/12/multilevellinkedlist.png
After flattening the multilevel linked list it becomes:
https://assets.leetcode.com/uploads/2018/10/12/multilevellinkedlistflattened.png
Example 2:
Input: head = [1,2,null,3]
Output: [1,3,2]
Explanation:
The input multilevel linked list is as follows:
  1---2---NULL
  |
  3---NULL
Example 3:
Input: head = []
Output: []
How multilevel linked list is represented in test case:
We use the multilevel linked list from Example 1 above:
 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL
The serialization of each level is as follows:
[1,2,3,4,5,6,null]
[7,8,9,10,null]
[11,12,null]
To serialize all levels together we will add nulls in each level to signify no node connects to the upper node of the previous level. The serialization becomes:
[1,2,3,4,5,6,null]
[null,null,7,8,9,10,null]
[null,11,12,null]
Merging the serialization of each level and removing trailing nulls we obtain:
[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
Constraints:
Number of Nodes will not exceed 1000.
1 <= Node.val <= 10^5
```


### Day11 : Subsets

###  [Subsets](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day11-subsets.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/July2020/day11-subsets_test.go)
```
Given a set of distinct integers, nums, return all possible subsets (the power set).
Note: The solution set must not contain duplicate subsets.
Example:
Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```
