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