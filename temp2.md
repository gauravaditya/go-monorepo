# Data Structures and Algorithms Practice Problems

## Arrays and Strings

### 1. Two Sum

**Problem Statement:** Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`. You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

### 2. Valid Anagram

**Problem Statement:** Given two strings `s` and `t`, return true if `t` is an anagram of `s`, and false otherwise. An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, using all the original letters exactly once.

Example:

```
Input: s = "anagram", t = "nagaram"
Output: true
```

### 3. Maximum Subarray

**Problem Statement:** Given an integer array `nums`, find the contiguous subarray with the largest sum and return its sum.

Example:

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
```

### 4. Product of Array Except Self

**Problem Statement:** Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

Example:

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Explanation: 
answer[0] = 2*3*4 = 24
answer[1] = 1*3*4 = 12
answer[2] = 1*2*4 = 8
answer[3] = 1*2*3 = 6
```

## Hash Tables

### 1. Contains Duplicate

**Problem Statement:** Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.

Example:

```
Input: nums = [1,2,3,1]
Output: true
```

### 2. First Unique Character

**Problem Statement:** Given a string `s`, find the first non-repeating character and return its index. If it does not exist, return -1.

Example:

```
Input: s = "leetcode"
Output: 0
Explanation: The first non-repeating character is 'l' with index 0.
```

### 3. Group Anagrams

**Problem Statement:** Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

Example:

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

### 4. LRU Cache

**Problem Statement:** Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initialize the LRU cache with positive size `capacity`.
- `int get(int key)` Return the value of the `key` if the key exists, otherwise return -1.
- `void put(int key, int value)` Update the value of the `key` if the `key` exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity` from this operation, evict the least recently used key.

## Linked Lists

### 1. Reverse Linked List

**Problem Statement:** Given the head of a singly linked list, reverse the list, and return the reversed list.

Example:

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

### 2. Detect Cycle

**Problem Statement:** Given head of a linked list, determine if the linked list has a cycle in it. Return true if there is a cycle, false otherwise.

Example:

```
Input: head = [3,2,0,-4], where -4 points back to 2
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
```

### 3. Merge Two Sorted Lists

**Problem Statement:** You are given the heads of two sorted linked lists `list1` and `list2`. Merge the two lists into one sorted list and return the head of the merged linked list.

Example:

```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

### 4. Remove Nth Node From End

**Problem Statement:** Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example:

```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

## Stacks and Queues

### 1. Valid Parentheses

**Problem Statement:** Given a string `s` containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Example:

```
Input: s = "()[]{}"
Output: true
```

### 2. Implement Queue using Stacks

**Problem Statement:** Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (`push`, `peek`, `pop`, and `empty`).

### 3. Min Stack

**Problem Statement:** Design a stack that supports push, pop, top, and retrieving the minimum element in constant time. Implement the `MinStack` class:

- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element `val` onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

### 4. Daily Temperatures

**Problem Statement:** Given an array of integers `temperatures` represents the daily temperatures, return an array `answer` such that `answer[i]` is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

Example:

```
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
```

## Trees and Binary Search Trees

### 1. Maximum Depth of Binary Tree

**Problem Statement:** Given the root of a binary tree, return its maximum depth. A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Example:

```
Input: root = [3,9,20,null,null,15,7]
Output: 3
```

### 2. Validate Binary Search Tree

**Problem Statement:** Given the root of a binary tree, determine if it is a valid binary search tree (BST). A valid BST is defined as follows:

- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

### 3. Lowest Common Ancestor

**Problem Statement:** Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree. According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself)."

### 4. Binary Tree Level Order Traversal

**Problem Statement:** Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example:

```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
```

## Heaps/Priority Queues

### 1. Kth Largest Element

**Problem Statement:** Given an integer array `nums` and an integer `k`, return the kth largest element in the array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example:

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```

### 2. Merge K Sorted Lists

**Problem Statement:** You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.

Example:

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
```

### 3. Top K Frequent Elements

**Problem Statement:** Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

Example:

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

### 4. Find Median from Data Stream

**Problem Statement:** Design a data structure that supports adding integers and finding the median of all elements so far. Implement the MedianFinder class:

- `MedianFinder()` initializes the MedianFinder object.
- `void addNum(int num)` adds the integer `num` from the data stream to the data structure.
- `double findMedian()` returns the median of all elements so far.

## Graphs

### 1. Number of Islands

**Problem Statement:** Given an `m x n` 2D binary grid `grid` which represents a map of '1's (land) and '0's (water), return the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example:

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

### 2. Course Schedule

**Problem Statement:** There are a total of `numCourses` courses you have to take, labeled from 0 to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you must take course `bi` first if you want to take course `ai`. Return true if you can finish all courses. Otherwise, return false.

Example:

```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are 2 courses to take. To take course 1 you should have finished course 0. So it is possible.
```

### 3. Pacific Atlantic Water Flow

**Problem Statement:** There is an `m x n` rectangular island that borders both the Pacific Ocean and Atlantic Ocean. The Pacific Ocean touches the island's left and top edges, and the Atlantic Ocean touches the island's right and bottom edges. The island is partitioned into a grid of square cells. You are given an `m x n` integer matrix `heights` where `heights[r][c]` represents the height above sea level of the cell at coordinate `(r, c)`. The island receives a lot of rain, and the rain water can flow to neighboring cells directly north, south, east, and west if the neighboring cell's height is less than or equal to the current cell's height. Water can flow from any cell adjacent to an ocean into the ocean. Return a list of grid coordinates where water can flow to both the Pacific and Atlantic oceans.

### 4. Network Delay Time

**Problem Statement:** You are given a network of `n` nodes, labeled from 1 to `n`. You are also given times, a list of travel times as directed edges `times[i] = (ui, vi, wi)`, where `ui` is the source node, `vi` is the target node, and `wi` is the time it takes for a signal to travel from source to target. We will send a signal from a given node `k`. Return the time it takes for all the `n` nodes to receive the signal. If it is impossible for all the `n` nodes to receive the signal, return -1.

Example:

```
Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2
```

## Sorting and Searching

### 1. Merge Sorted Array

**Problem Statement:** You are given two integer arrays `nums1` and `nums2`, sorted in non-decreasing order, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively. Merge `nums1` and `nums2` into a single array sorted in non-decreasing order. The final sorted array should not be returned by the function, but instead be stored inside the array `nums1`. To accommodate this, `nums1` has a length of `m + n`, where the first `m` elements denote the elements that should be merged, and the last `n` elements are set to 0 and should be ignored. `nums2` has a length of `n`.

Example:

```
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
```

### 2. Search in Rotated Sorted Array

**Problem Statement:** There is an integer array `nums` sorted in ascending order (with distinct values). Prior to being passed to your function, `nums` is possibly rotated at an unknown pivot index `k` (1 <= k < nums.length) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (0-indexed). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index 3 and become `[4,5,6,7,0,1,2]`. Given the array `nums` after the possible rotation and an integer `target`, return the index of `target` if it is in `nums`, or -1 if it is not in `nums`.

Example:

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

### 3. Find First and Last Position

**Problem Statement:** Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value. If `target` is not found in the array, return `[-1, -1]`. You must write an algorithm with O(log n) runtime complexity.

Example:

```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

### 4. Search a 2D Matrix

**Problem Statement:** Write an efficient algorithm that searches for a value `target` in an `m x n` integer matrix `matrix`. This matrix has the following properties:

- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

Example:

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

## Dynamic Programming

### 1. Climbing Stairs

**Problem Statement:** You are climbing a staircase. It takes `n` steps to reach the top. Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example:

```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

### 2. Coin Change

**Problem Statement:** You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money. Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1. You may assume that you have an infinite number of each kind of coin.

Example:

```
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
```

### 3. Longest Increasing Subsequence

**Problem Statement:** Given an integer array `nums`, return the length of the longest strictly increasing subsequence. A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.

Example:

```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

### 4. House Robber

**Problem Statement:** You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night. Given an integer array `nums` representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Example:

```
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3). Total amount you can rob = 1 + 3 = 4.
```

## Recursion and Backtracking

### 1. Generate Parentheses

**Problem Statement:** Given `n` pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example:

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

### 2. Subsets

**Problem Statement:** Given an integer array `nums` of unique elements, return all possible subsets (the power set). The solution set must not contain duplicate subsets. Return the solution in any order.

Example:

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

### 3. Letter Combinations of Phone Number

**Problem Statement:** Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order. A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

```
2 -> abc
3 -> def
4 -> ghi
5 -> jkl
6 -> mno
7 -> pqrs
8 -> tuv
9 -> wxyz
```

Example:

```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

### 4. Word Search

**Problem Statement:** Given an `m x n` grid of characters `board` and a string `word`, return `true` if `word` exists in the grid. The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```