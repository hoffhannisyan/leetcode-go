package main

import (
	"fmt"

	"github.com/hoffhannisyan/leetcode/tasks"
)

func main() {
	fmt.Println("LeetCode Solutions in Go")
	fmt.Println("========================")

	// --- LeetCode #1: Two Sum ---
	/*
		printTitle(1, "Two Sum")

		nums := []int{2, 7, 11, 15}
		target := 9
		result := tasks.TwoSum(nums, target)
		fmt.Println("Input:", nums, "Target:", target)
		fmt.Println("Output:", result)
	*/

	// --- LeetCode #2: Add Two Numbers ---
	/*
		printTitle(2, "Add Two Numbers")

		l1 := &tasks.ListNode{Val: 2, Next: &tasks.ListNode{Val: 4, Next: &tasks.ListNode{Val: 3}}}
		l2 := &tasks.ListNode{Val: 5, Next: &tasks.ListNode{Val: 6, Next: &tasks.ListNode{Val: 4}}}
		resultList := tasks.AddTwoNumbers(l1, l2)

		fmt.Print("Output: [")
		for node := resultList; node != nil; node = node.Next {
			fmt.Print(node.Val)
			if node.Next != nil {
				fmt.Print(" ")
			}
		}
		fmt.Println("]")
	*/

	// --- LeetCode #3: Longest Substring Without Repeating Characters ---
	/*
		printTitle(3, "Longest Substring Without Repeating Characters")

		s := "pwwkew"
		length := tasks.LengthOfLongestSubstring(s)
		fmt.Println("Input:", s)
		fmt.Println("Output:", length)
	*/

	// --- LeetCode #4: Median of Two Sorted Arrays ---
	/*
		 printTitle(4, "Median of Two Sorted Arrays")

		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		median := tasks.FindMedianSortedArrays(nums1, nums2)
		fmt.Println("Input:", nums1, nums2)
		fmt.Println("Output:", median)
	*/

	// --- LeetCode #5: Longest Palindromic Substring ---

	printTitle(5, "Longest Palindromic Substring")

	s1 := "babad"
	fmt.Printf("Input: %q -> Output: %q\n", s1, tasks.LongestPalindrome(s1))

	s2 := "cbbd"
	fmt.Printf("Input: %q -> Output: %q\n", s2, tasks.LongestPalindrome(s2))

}

func printTitle(id int, title string) {
	fmt.Printf("\n--- LeetCode #%d: %s ---\n", id, title)
}
