package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	// A map to store the last index of each character
	lastIndex := make(map[rune]int)
	maxLength := 0
	start := 0

	// Iterate through the string with an end pointer
	for end := 0; end < len(s); end++ {
		// If the character was already seen, update the start pointer
		if lastPos, found := lastIndex[rune(s[end])]; found && lastPos >= start {
			start = lastPos + 1
		}

		// Update the last index of the current character
		lastIndex[rune(s[end])] = end

		// Calculate the current length of the substring
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

// Helper function to get the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test the function with some examples
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
}
