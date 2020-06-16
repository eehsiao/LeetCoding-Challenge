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