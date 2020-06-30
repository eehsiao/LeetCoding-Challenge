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


### Day3 : Two City Scheduling

###  [Two City Scheduling](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day3-twoCityScheduling.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day3-twoCityScheduling_test.go)
```
There are 2N people a company is planning to interview. The cost of flying the i-th person to city A is costs[i][0], and the cost of flying the i-th person to city B is costs[i][1].
Return the minimum cost to fly every person to a city such that exactly N people arrive in each city.
Example 1:
Input: [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.
The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.
Note:
1 <= costs.length <= 100
It is guaranteed that costs.length is even.
1 <= costs[i][0], costs[i][1] <= 1000
```


### Day4 : Reverse String

###  [Reverse String](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day4-reverseString.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day4-reverseString_test.go)
```
Write a function that reverses a string. The input string is given as an array of characters char[].
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
You may assume all the characters consist of printable ascii characters.
Example 1:
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
   Hide Hint #1
The entire logic for reversing a string is based on using the opposite directional two-pointer approach!
```


### Day5 : Random Pick with Weight

###  [Random Pick with Weight](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day5-randomPickwithWeight.go)
```
Given an array w of positive integers, where w[i] describes the weight of index i, write a function pickIndex which randomly picks an index in proportion to its weight.
For example, given an input list of values [1, 9], when we pick up a number out of it, the chance is that 9 times out of 10 we should pick the number 9 as the answer.
Example 1:
Input:
["Solution","pickIndex"]
[[[1]],[]]
Output: [null,0]
Example 2:
Input:
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
Output: [null,0,1,1,1,0]
Explanation of Input Syntax:
The input is two lists: the subroutines called and their arguments. Solution's constructor has one argument, the array w. pickIndex has no arguments. Arguments are always wrapped with a list, even if there aren't any.
Constraints:
1 <= w.length <= 10000
1 <= w[i] <= 10^5
pickIndex will be called at most 10000 times.
```


### Day6 : Queue Reconstruction by Height

###  [Queue Reconstruction by Height](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day6-queueReconstructionbyHeight.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day6-queueReconstructionbyHeight_test.go)
```
Suppose you have a random list of people standing in a queue. Each person is described by a pair of integers (h, k), where h is the height of the person and k is the number of people in front of this person who have a height greater than or equal to h. Write an algorithm to reconstruct the queue.
Note:
The number of people is less than 1,100.
Example
Input:
[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
Output:
[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
   Hide Hint #1
What can you say about the position of the shortest person?
If the position of the shortest person is i, how many people would be in front of the shortest person?
   Hide Hint #2
Once you fix the position of the shortest person, what can you say about the position of the second shortest person?
```


### Day7 : Coin Change 2

###  [Coin Change 2](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day7-coinChange2.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day7-coinChange2_test.go)
```
You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.
Example 1:
Input: amount = 5, coins = [1, 2, 5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
Example 2:
Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.
Example 3:
Input: amount = 10, coins = [10]
Output: 1
Note:
You can assume that
0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer
```


### Day8 : Power of Two

###  [Power of Two](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day8-powerofTwo.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day8-powerofTwo_test.go)
```
Given an integer, write a function to determine if it is a power of two.
Example 1:
Input: 1
Output: true
Explanation: 20 = 1
Example 2:
Input: 16
Output: true
Explanation: 24 = 16
Example 3:
Input: 218
Output: false
```


### Day9 : Is Subsequence

###  [Is Subsequence](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day9-isSubsequence.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day9-isSubsequence_test.go)
```
Given a string s and a string t, check if s is subsequence of t.
A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).
Follow up:
If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you want to check one by one to see if T has its subsequence. In this scenario, how would you change your code?
Credits:
Special thanks to @pbrother for adding this problem and creating all test cases.
Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false
Constraints:
0 <= s.length <= 100
0 <= t.length <= 10^4
Both strings consists only of lowercase characters.
```


### Day10 : Search Insert Position

###  [Search Insert Position](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day10-searchInsertPosition.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day10-searchInsertPosition_test.go)
```
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You may assume no duplicates in the array.
Example 1:
Input: [1,3,5,6], 5
Output: 2
Example 2:
Input: [1,3,5,6], 2
Output: 1
Example 3:
Input: [1,3,5,6], 7
Output: 4
Example 4:
Input: [1,3,5,6], 0
Output: 0
```


### Day11 : Sort Colors

###  [Sort Colors](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day11-sortColors.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day11-sortColors_test.go)
```
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.
Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
Note: You are not suppose to use the library's sort function for this problem.
Example:
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Follow up:
A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?
```


### Day12 : Insert Delete GetRandom O(1)

###  [Insert Delete GetRandom O(1)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day12-insertDeleteGetRandomO(1).go)
```
Design a data structure that supports all following operations in average O(1) time.

insert(val): Inserts an item val to the set if not already present.
remove(val): Removes an item val from the set if present.
getRandom: Returns a random element from current set of elements. Each element must have the same probability of being returned.
Example:

// Init an empty set.
RandomizedSet randomSet = new RandomizedSet();

// Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomSet.insert(1);

// Returns false as 2 does not exist in the set.
randomSet.remove(2);

// Inserts 2 to the set, returns true. Set now contains [1,2].
randomSet.insert(2);

// getRandom should return either 1 or 2 randomly.
randomSet.getRandom();

// Removes 1 from the set, returns true. Set now contains [2].
randomSet.remove(1);

// 2 was already in the set, so return false.
randomSet.insert(2);

// Since 2 is the only number in the set, getRandom always return 2.
randomSet.getRandom();
```


### Day13 : Largest Divisible Subset

###  [Largest Divisible Subset](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day13-largestDivisibleSubset.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day13-largestDivisibleSubset_test.go)
```
Given a set of distinct positive integers, find the largest subset such that every pair (Si, Sj) of elements in this subset satisfies:
Si % Sj = 0 or Sj % Si = 0.
If there are multiple solutions, return any subset is fine.
Example 1:
Input: [1,2,3]
Output: [1,2] (of course, [1,3] will also be ok)
Example 2:
Input: [1,2,4,8]
Output: [1,2,4,8]
```


### Day14 : Cheapest Flights Within K Stops

###  [Cheapest Flights Within K Stops](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day14-cheapestFlightsWithinKStops.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day14-cheapestFlightsWithinKStops_test.go)
```
There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.
Now given all the cities and flights, together with starting city src and the destination dst, your task is to find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.
Example 1:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
Output: 200
Explanation:
The graph looks like this:
        0
 (100) / \ (500)
      1 - 2
      (100)
The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.
Example 2:
Input:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
Output: 500
Explanation:
The graph looks like this:
        0
 (100) / \ (500)
      1 - 2
      (100)
The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as marked blue in the picture.
Constraints:
The number of nodes n will be in range [1, 100], with nodes labeled from 0 to n - 1.
The size of flights will be in range [0, n * (n - 1) / 2].
The format of each flight will be (src, dst, price).
The price of each flight will be in the range [1, 10000].
k is in the range of [0, n - 1].
There will not be any duplicated flights or self cycles.
```


### Day15 : Search in a Binary Search Tree

###  [Search in a Binary Search Tree](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day15-searchinaBinarySearchTree.go)
```
Given the root node of a binary search tree (BST) and a value. You need to find the node in the BST that the node's value equals the given value. Return the subtree rooted with that node. If such node doesn't exist, you should return NULL.
For example,
Given the tree:
        4
       / \
      2   7
     / \
    1   3
And the value to search: 2
You should return this subtree:
      2
     / \
    1   3
In the example above, if we want to search the value 5, since there is no node with value 5, we should return NULL.
Note that an empty tree is represented by NULL, therefore you would see the expected output (serialized tree format) as [], not null.
```


### Day16 : Validate IP Address

###  [Validate IP Address](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day16-validateIPAddress.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day16-validateIPAddress_test.go)
```
Write a function to check whether an input string is a valid IPv4 address or IPv6 address or neither.
IPv4 addresses are canonically represented in dot-decimal notation, which consists of four decimal numbers, each ranging from 0 to 255, separated by dots ("."), e.g.,172.16.254.1;
Besides, leading zeros in the IPv4 is invalid. For example, the address 172.16.254.01 is invalid.
IPv6 addresses are represented as eight groups of four hexadecimal digits, each group representing 16 bits. The groups are separated by colons (":"). For example, the address 2001:0db8:85a3:0000:0000:8a2e:0370:7334 is a valid one. Also, we could omit some leading zeros among four hexadecimal digits and some low-case characters in the address to upper-case ones, so 2001:db8:85a3:0:0:8A2E:0370:7334 is also a valid IPv6 address(Omit leading zeros and using upper cases).
However, we don't replace a consecutive group of zero value with a single empty group using two consecutive colons (::) to pursue simplicity. For example, 2001:0db8:85a3::8A2E:0370:7334 is an invalid IPv6 address.
Besides, extra leading zeros in the IPv6 is also invalid. For example, the address 02001:0db8:85a3:0000:0000:8a2e:0370:7334 is invalid.
Note: You may assume there is no extra space or special characters in the input string.
Example 1:
Input: "172.16.254.1"
Output: "IPv4"
Explanation: This is a valid IPv4 address, return "IPv4".
Example 2:
Input: "2001:0db8:85a3:0:0:8A2E:0370:7334"
Output: "IPv6"
Explanation: This is a valid IPv6 address, return "IPv6".
Example 3:
Input: "256.256.256.256"
Output: "Neither"
Explanation: This is neither a IPv4 address nor a IPv6 address.
```


### Day17 : Validate IP Address

###  [Validate IP Address](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day17-surroundedRegions.go)
```
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
A region is captured by flipping all 'O's into 'X's in that surrounded region.
Example:
X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:
X X X X
X X X X
X X X X
X O X X
Explanation:
Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
```


### Day18 : H-Index II

###  [H-Index II](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day18-h-IndexII.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day18-h-IndexII_test.go)
```
Given an array of citations sorted in ascending order (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.
According to the definition of h-index on Wikipedia: "A scientist has index h if h of his/her N papers have at least h citations each, and the other N − h papers have no more than h citations each."
Example:
Input: citations = [0,1,3,5,6]
Output: 3
Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had
             received 0, 1, 3, 5, 6 citations respectively.
             Since the researcher has 3 papers with at least 3 citations each and the remaining
             two with no more than 3 citations each, her h-index is 3.
Note:
If there are several possible values for h, the maximum one is taken as the h-index.
Follow up:
This is a follow up problem to H-Index, where citations is now guaranteed to be sorted in ascending order.
Could you solve it in logarithmic time complexity?
   Hide Hint #1
Expected runtime complexity is in O(log n) and the input is sorted.
```


### Day19 : Longest Duplicate Substring

###  [Longest Duplicate Substring](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day19-longestDuplicateSubstring.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day19-longestDuplicateSubstring_test.go)
```
Given a string S, consider all duplicated substrings: (contiguous) substrings of S that occur 2 or more times.  (The occurrences may overlap.)
Return any duplicated substring that has the longest possible length.  (If S does not have a duplicated substring, the answer is "".)
Example 1:
Input: "banana"
Output: "ana"
Example 2:
Input: "abcd"
Output: ""
Note:
2 <= S.length <= 10^5
S consists of lowercase English letters.
   Hide Hint #1
Binary search for the length of the answer. (If there's an answer of length 10, then there are answers of length 9, 8, 7, ...)
   Hide Hint #2
To check whether an answer of length K exists, we can use Rabin-Karp 's algorithm.
```


### Day20 : Permutation Sequence

###  [Permutation Sequence](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day20-permutationSequence.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day20-permutationSequence_test.go)
```
The set [1,2,3,...,n] contains a total of n! unique permutations.
By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.
Note:
Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.
Example 1:
Input: n = 3, k = 3
Output: "213"
Example 2:
Input: n = 4, k = 9
Output: "2314"
```


### Day21 : Dungeon Game

###  [Dungeon Game](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day21-dungeonGame.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day21-dungeonGame_test.go)
```
The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.
The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.
Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms; other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).
In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.
Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.
For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.
-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)
Note:
The knight's health has no upper bound.
Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

```


### Day22 : Single Number II

###  [Single Number II](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day22-singleNumberII.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day22-singleNumberII_test.go)
```
Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.
Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
Example 1:
Input: [2,2,3,2]
Output: 3
Example 2:
Input: [0,1,0,1,0,1,99]
Output: 99
```


### Day23 : Count Complete Tree Nodes

###  [Count Complete Tree Nodes](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day23-countCompleteTreeNodes.go)
```
Given a complete binary tree, count the number of nodes.
Note:
Definition of a complete binary tree from Wikipedia:
In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.
Example:
Input:
    1
   / \
  2   3
 / \  /
4  5 6
Output: 6
```


### Day24 : Unique Binary Search Trees

###  [Unique Binary Search Trees](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day24-uniqueBinarySearchTrees.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day24-uniqueBinarySearchTrees_test.go)
```
Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?
Example:
Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:
   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```


### Day25 : Find the Duplicate Number

###  [Find the Duplicate Number](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day25-findtheDuplicateNumber.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day25-findtheDuplicateNumber_test.go)
```
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.
Example 1:
Input: [1,3,4,2,2]
Output: 2
Example 2:
Input: [3,1,3,4,2]
Output: 3
Note:
1.You must not modify the array (assume the array is read only).
2.You must use only constant, O(1) extra space.
3.Your runtime complexity should be less than O(n2).
4.There is only one duplicate number in the array, but it could be repeated more than once.
```


### Day26 : Find the Duplicate Number

###  [Find the Duplicate Number](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day26-sumRoottoLeafNumbers.go)
```
Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.
An example is the root-to-leaf path 1->2->3 which represents the number 123.
Find the total sum of all root-to-leaf numbers.
Note: A leaf is a node with no children.
Example:
Input: [1,2,3]
    1
   / \
  2   3
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
Example 2:
Input: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.
```


### Day27 : Perfect Squares

###  [Perfect Squares](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day27-perfectSquares.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day27-perfectSquares_test.go)
```
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.
Example 1:
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
```


### Day28 : Reconstruct Itinerary

###  [Reconstruct Itinerary](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day28-reconstructItinerary.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day28-reconstructItinerary_test.go)
```
Given a list of airline tickets represented by pairs of departure and arrival airports [from, to], reconstruct the itinerary in order. All of the tickets belong to a man who departs from JFK. Thus, the itinerary must begin with JFK.
Note:
If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string. For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
All airports are represented by three capital letters (IATA code).
You may assume all tickets form at least one valid itinerary.
One must use all the tickets once and only once.
Example 1:
Input: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
Output: ["JFK", "MUC", "LHR", "SFO", "SJC"]
Example 2:
Input: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"].
             But it is larger in lexical order.
```


### Day29 : Unique Paths

###  [Unique Paths](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day29-uniquePaths.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/June2020/day29-uniquePaths_test.go)
```
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
How many possible unique paths are there?
Above is a 7 x 3 grid. How many possible unique paths are there?
Example 1:
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Right -> Down
2. Right -> Down -> Right
3. Down -> Right -> Right
Example 2:
Input: m = 7, n = 3
Output: 28
Constraints:
1 <= m, n <= 100
It's guaranteed that the answer will be less than or equal to 2 * 10 ^ 9.
```
