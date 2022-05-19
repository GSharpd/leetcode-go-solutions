// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4
// Example 2:

// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

package main

func search(n []int, t int) int {
	l := 0
	h := len(n) - 1

	for l <= h {
		m := (l + h) / 2

		if n[m] < t {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	if l == len(n) || n[l] != t {
		return -1
	} else {
		return l
	}
}
