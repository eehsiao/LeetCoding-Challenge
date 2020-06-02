# 2020-May-LeetCoding-Challenge
[Challenge site](https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/)

[My leetcode](https://leetcode.com/eehsiao/)

## Week 1: May 1st–May 7th

### Day1 : First Bad Version

###  [First Bad Version](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day1-firstBadVersion.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day1-firstBadVersion_test.go)
```
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.
Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.
Example:
Given n = 5, and version = 4 is the first bad version.
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version. 
```

### Day2 : Jewels and Stones

###  [Jewels and Stones](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day2-jewelsandStones.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day2-jewelsandStones_test.go)
```
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".
Example 1:
Input: J = "aA", S = "aAAbbbb"
Output: 3
Example 2:
Input: J = "z", S = "ZZ"
Output: 0
Note:
S and J will consist of letters and have length at most 50.
The characters in J are distinct.
   Hide Hint #1
For each stone, check if it is a jewel.
```


### Day3 : Ransom Note

###  [Ransom Note](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day3-ransomNote.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day3-ransomNote_test.go)
```
Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.
Each letter in the magazine string can only be used once in your ransom note.
Note:
You may assume that both strings contain only lowercase letters.
canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
```

### Day4 : Number Complement

###  [Number Complement](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day4-numberComplement.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day4-numberComplement_test.go)
```
Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.
Example 1:
Input: 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
Example 2:
Input: 1
Output: 0
Explanation: The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.

Note:
The given integer is guaranteed to fit within the range of a 32-bit signed integer.
You could assume no leading zero bit in the integer’s binary representation.
This question is the same as 1009: https://leetcode.com/problems/complement-of-base-10-integer/
```

### Day5 : First Unique Character in a String

###  [First Unique Character in a String](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day5-firstUniqueCharacterinaString.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day5-firstUniqueCharacterinaString_test.go)
```
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
Examples:
s = "leetcode"
return 0.
s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
its like https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week4-firstUniqueNumber.go
```


### Day6 : Majority Element

###  [Majority Element](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day6-majorityElement.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day6-majorityElement_test.go)
```
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
You may assume that the array is non-empty and the majority element always exist in the array.
Example 1:
Input: [3,2,3]
Output: 3
Example 2:
Input: [2,2,1,1,1,2,2]
Output: 2
```

### Day7 : Cousins in Binary Tree

###  [Cousins in Binary Tree](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day7-cousinsinBinaryTree.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day7-cousinsinBinaryTree_test.go)
```
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
Return true if and only if the nodes corresponding to the values x and y are cousins.
Example 1:
      1
     / \
    2   3
   /
  4
Input: root = [1,2,3,4], x = 4, y = 3
Output: false
Example 2:
    1
   / \
  2   3
   \   \
    4   5
Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true
Example 3:
    1
   / \
  2   3
   \
    4
Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false
Note:
The number of nodes in the tree will be between 2 and 100.
Each node has a unique integer value from 1 to 100.
```

### Day8 : Check If It Is a Straight Line

###  [Check If It Is a Straight Line](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day8-checkIfItIsaStraightLine.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day8-checkIfItIsaStraightLine_test.go)
```
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.
Example 1:
https://assets.leetcode.com/uploads/2019/10/15/untitled-diagram-2.jpg
Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true
Example 2:
https://assets.leetcode.com/uploads/2019/10/09/untitled-diagram-1.jpg
Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false
Constraints:
2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates contains no duplicate point.
   Hide Hint #1
If there're only 2 points, return true.
   Hide Hint #2
Check if all other points lie on the line defined by the first 2 points.
   Hide Hint #3
Use cross product to check collinearity.
```

### Day9 : Valid Perfect Square

###  [Valid Perfect Square](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day9-validPerfectSquare.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day9-validPerfectSquare_test.go)
```
Given a positive integer num, write a function which returns True if num is a perfect square else False.
Note: Do not use any built-in library function such as sqrt.
Example 1:
Input: 16
Output: true
Example 2:
Input: 14
Output: false
```

### Day10 : Find the Town Judge

###  [Find the Town Judge](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day10-findtheTownJudge.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day10-findtheTownJudge_test.go)
```
In a town, there are N people labelled from 1 to N.  There is a rumor that one of these people is secretly the town judge.
If the town judge exists, then:
The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.
If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.
Example 1:
Input: N = 2, trust = [[1,2]]
Output: 2
Example 2:
Input: N = 3, trust = [[1,3],[2,3]]
Output: 3
Example 3:
Input: N = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1
Example 4:
Input: N = 3, trust = [[1,2],[2,3]]
Output: -1
Example 5:
Input: N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
Output: 3
Note:
1 <= N <= 1000
trust.length <= 10000
trust[i] are all different
trust[i][0] != trust[i][1]
1 <= trust[i][0], trust[i][1] <= N
```

### Day11 : Flood Fill

###  [Flood Fill](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day11-floodFill.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day11-floodFill_test.go)

another solution

###  [Flood Fill](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day11-floodFill2.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day11-floodFil2l_test.go)

```
An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).
Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.
To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.
At the end, return the modified image.
Example 1:
Input:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation:
From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
by a path of the same color as the starting pixel are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected
to the starting pixel.
Note:
The length of image and image[0] will be in the range [1, 50].
The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
The value of each color in image[i][j] and newColor will be an integer in [0, 65535].
   Hide Hint #1
Write a recursive function that paints the pixel if it's the correct color, then recurses on neighboring pixels.
```

### Day12 : Single Element in a Sorted Array

###  [Single Element in a Sorted Array](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day12-singleElementinaSortedArray.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day12-singleElementinaSortedArray_test.go)
```
You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once. Find this single element that appears only once.
Example 1:
Input: [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:
Input: [3,3,7,7,10,11,11]
Output: 10
Note: Your solution should run in O(log n) time and O(1) space.
```

### Day13 : Remove K Digits

###  [Remove K Digits](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day13-removeKDigits.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day13-removeKDigits_test.go)
```
Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.
Note:
The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.
Example 1:
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
Example 2:
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
Example 3:
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
```

### Day14 : Remove K Digits

###  [Remove K Digits](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day14-implementTrie.go)
```
Implement a trie with insert, search, and startsWith methods.
Example:
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:
You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.

```

### Day15 : Maximum Sum Circular Subarray

###  [Maximum Sum Circular Subarray](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day15-maximumSumCircularSubarray.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day15-maximumSumCircularSubarray_test.go)
```
Given a circular array C of integers represented by A, find the maximum possible sum of a non-empty subarray of C.
Here, a circular array means the end of the array connects to the beginning of the array.  (Formally, C[i] = A[i] when 0 <= i < A.length, and C[i+A.length] = C[i] when i >= 0.)
Also, a subarray may only include each element of the fixed buffer A at most once.  (Formally, for a subarray C[i], C[i+1], ..., C[j], there does not exist i <= k1, k2 <= j with k1 % A.length = k2 % A.length.)
Example 1:
Input: [1,-2,3,-2]
Output: 3
Explanation: Subarray [3] has maximum sum 3
Example 2:
Input: [5,-3,5]
Output: 10
Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10
Example 3:
Input: [3,-1,2,-1]
Output: 4
Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4
Example 4:
Input: [3,-2,2,-3]
Output: 3
Explanation: Subarray [3] and [3,-2,2] both have maximum sum 3
Example 5:
Input: [-2,-3,-1]
Output: -1
Explanation: Subarray [-1] has maximum sum -1
Note:
-30000 <= A[i] <= 30000
1 <= A.length <= 30000
   Hide Hint #1
For those of you who are familiar with the Kadane's algorithm, think in terms of that. For the newbies, Kadane's algorithm is used to finding the maximum sum subarray from a given array. This problem is a twist on that idea and it is advisable to read up on that algorithm first before starting this problem. Unless you already have a great algorithm brewing up in your mind in which case, go right ahead!
   Hide Hint #2
What is an alternate way of representing a circular array so that it appears to be a straight array? Essentially, there are two cases of this problem that we need to take care of. Let's look at the figure below to understand those two cases:
   Hide Hint #3
The first case can be handled by the good old Kadane's algorithm. However, is there a smarter way of going about handling the second case as well?
```

### Day16 : Odd Even Linked List

###  [Odd Even Linked List](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day16-oddEvenLinkedList.go)
```
Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.
You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.
Example 1:
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
Example 2:
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL
Note:
The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
```

### Day17 : Find All Anagrams in a String

###  [Find All Anagrams in a String](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day17-findAllAnagramsinaString.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day17-findAllAnagramsinaString_test.go)
```
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
The order of output does not matter.
Example 1:
Input:
s: "cbaebabacd" p: "abc"
Output:
[0, 6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:
Input:
s: "abab" p: "ab"
Output:
[0, 1, 2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
```

### Day18 : Permutation in String

###  [Permutation in String](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day18-permutationinString.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day18-permutationinString_test.go)
```
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
Example 1:
Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:
Input:s1= "ab" s2 = "eidboaoo"
Output: False
Note:
The input strings only contain lower case letters.
The length of both given strings is in range [1, 10,000].
   Hide Hint #1
Obviously, brute force will result in TLE. Think of something else.
   Hide Hint #2
How will you check whether one string is a permutation of another string?
   Hide Hint #3
One way is to sort the string and then compare. But, Is there a better way?
   Hide Hint #4
If one string is a permutation of another string then they must one common metric. What is that?
   Hide Hint #5
Both strings must have same character frequencies, if one is permutation of another. Which data structure should be used to store frequencies?
   Hide Hint #6
What about hash table? An array of size 26?
```

### Day19 : Online Stock Span

###  [Online Stock Span](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day19-onlineStockSpan.go)
```
Write a class StockSpanner which collects daily price quotes for some stock, and returns the span of that stock's price for the current day.
The span of the stock's price today is defined as the maximum number of consecutive days (starting from today and going backwards) for which the price of the stock was less than or equal to today's price.
For example, if the price of a stock over the next 7 days were [100, 80, 60, 70, 60, 75, 85], then the stock spans would be [1, 1, 1, 2, 1, 4, 6].
Example 1:
Input: ["StockSpanner","next","next","next","next","next","next","next"], [[],[100],[80],[60],[70],[60],[75],[85]]
Output: [null,1,1,1,2,1,4,6]
Explanation:
First, S = StockSpanner() is initialized.  Then:
S.next(100) is called and returns 1,
S.next(80) is called and returns 1,
S.next(60) is called and returns 1,
S.next(70) is called and returns 2,
S.next(60) is called and returns 1,
S.next(75) is called and returns 4,
S.next(85) is called and returns 6.
Note that (for example) S.next(75) returned 4, because the last 4 prices
(including today's price of 75) were less than or equal to today's price.
Note:
Calls to StockSpanner.next(int price) will have 1 <= price <= 10^5.
There will be at most 10000 calls to StockSpanner.next per test case.
There will be at most 150000 calls to StockSpanner.next across all test cases.
The total time limit for this problem has been reduced by 75% for C++, and 50% for all other languages.
```

### Day20 : Kth Smallest Element in a BST

###  [Kth Smallest Element in a BST](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day20-kthSmallestElementinaBST.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day20-kthSmallestElementinaBST_test.go)
```
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
Example 1:
Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?
   Hide Hint #1
Try to utilize the property of a BST.
   Hide Hint #2
Try in-order traversal. (Credits to @chan13)
   Hide Hint #3
What if you could modify the BST node's structure?
   Hide Hint #4
The optimal runtime complexity is O(height of BST).
```

### Day21 : Kth Smallest Element in a BST

###  [Kth Smallest Element in a BST](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day21-countSquareSubmatriceswithAllOnes.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day21-countSquareSubmatriceswithAllOnes_test.go)
```
Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
Example 1:
Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation:
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.
Example 2:
Input: matrix =
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
Output: 7
Explanation:
There are 6 squares of side 1.
There is 1 square of side 2.
Total number of squares = 6 + 1 = 7.
Constraints:
1 <= arr.length <= 300
1 <= arr[0].length <= 300
0 <= arr[i][j] <= 1
   Hide Hint #1
Create an additive table that counts the sum of elements of submatrix with the superior corner at (0,0).
   Hide Hint #2
Loop over all subsquares in O(n^3) and check if the sum make the whole array to be ones, if it checks then add 1 to the answer.
```

### Day22 : Sort Characters By Frequency

###  [Sort Characters By Frequency](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day22-sortCharactersByFrequency.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day22-sortCharactersByFrequency_test.go)

[same as](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week3-constructBinarySearchTreefromPreorderTraversal.go)
```
Given a string, sort it in decreasing order based on the frequency of characters.
Example 1:
Input:
"tree"
Output:
"eert"
Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
Example 2:
Input:
"cccaaa"
Output:
"cccaaa"
Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.
Example 3:
Input:
"Aabb"
Output:
"bbAa"
Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
```

### Day23 : Interval List Intersections

###  [Interval List Intersections](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day23-intervalListIntersections.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day23-intervalListIntersections_test.go)
```
Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.
Return the intersection of these two interval lists.
(Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.  The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)
Example 1:
Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.
Note:
0 <= A.length < 1000
0 <= B.length < 1000
0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.
```

### Day24 : Construct Binary Search Tree from Preorder Traversal

###  [Construct Binary Search Tree from Preorder Traversal](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day24-constructBinarySearchTreefromPreorderTraversal.go)

[Same as : 30-Day-LeetCoding-Challenge > week3-constructBinarySearchTreefromPreorderTraversal](https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week3-constructBinarySearchTreefromPreorderTraversal.go)

```
Return the node node of a binary search tree that matches the given preorder traversal.
(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)
Example 1:
Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]
       8
      / \
     5   10
    / \    \
   1   7    12
Note:
1 <= preorder.length <= 100
The values of preorder are distinct.
```

### Day25 : Uncrossed Lines

###  [Uncrossed Lines](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day25-uncrossedLines.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day25-uncrossedLines_test.go)
```
We write the integers of A and B (in the order they are given) on two separate horizontal lines.
Now, we may draw connecting lines: a straight line connecting two numbers A[i] and B[j] such that:
A[i] == B[j];
The line we draw does not intersect any other connecting (non-horizontal) line.
Note that a connecting lines cannot intersect even at the endpoints: each number can only belong to one connecting line.
Return the maximum number of connecting lines we can draw in this way.
Example 1:
1 4 2
|  \
1 2 4
Input: A = [1,4,2], B = [1,2,4]
Output: 2
Explanation: We can draw 2 uncrossed lines as in the diagram.
We cannot draw 3 uncrossed lines, because the line from A[1]=4 to B[2]=4 will intersect the line from A[2]=2 to B[1]=2.
Example 2:
Input: A = [2,5,1,2,5], B = [10,5,2,1,5,2]
Output: 3
Example 3:
Input: A = [1,3,7,1,7,5], B = [1,9,2,5,1]
Output: 2
Note:
1 <= A.length <= 500
1 <= B.length <= 500
1 <= A[i], B[i] <= 2000
   Hide Hint #1
Think dynamic programming. Given an oracle dp(i,j) that tells us how many lines A[i:], B[j:] [the sequence A[i], A[i+1], ... and B[j], B[j+1], ...] are uncrossed, can we write this as a recursion?
```

### Day26 : Possible Bipartition

###  [Possible Bipartition](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day26-possibleBipartition.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day26-possibleBipartition_test.go)
```
Given a set of N people (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.
Each person may dislike some other people, and they should not go into the same group.
Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.
Return true if and only if it is possible to split everyone into two groups in this way.
Example 1:
Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4], group2 [2,3]
Example 2:
Input: N = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false
Example 3:
Input: N = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
Output: false
Note:
1 <= N <= 2000
0 <= dislikes.length <= 10000
1 <= dislikes[i][j] <= N
dislikes[i][0] < dislikes[i][1]
There does not exist i != j for which dislikes[i] == dislikes[j].
```

### Day27 : Counting Bits

###  [Counting Bits](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day27-countingBits.go) [(Test Case)](https://github.com/eehsiao/LeetCoding-Challenge/blob/master/May2020/day27-countingBits_test.go)
```
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.
Example 1
Input: 2
Output: [0,1,1]
Example 2:
Input: 5
Output: [0,1,1,2,1,2]
Follow up:
It is very easy to come up with a solution with run time O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a single pass?
Space complexity should be O(n).
Can you do it like a boss? Do it without using any builtin function like __builtin_popcount in c++ or in any other language.
   Hide Hint #1
You should make use of what you have produced already.
   Hide Hint #2
Divide the numbers in ranges like [2-3], [4-7], [8-15] and so on. And try to generate new range from previous.
   Hide Hint #3
Or does the odd/even status of the number help you in calculating the number of 1s?
```