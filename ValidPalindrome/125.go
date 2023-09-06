// https://leetcode.com/problems/valid-palindrome/
// 125. Valid Palindrome
// Easy

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// Constraints:

// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters

package ValidPalindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	return P(s)
}

func P(s string) bool {
	r := []rune(strings.ToLower(s))
	i := 0
	j := len(r) - 1
	for cnt := 0; cnt < len(r); cnt++ {
		for i < len(r) && !unicode.IsLetter(r[i]) && !unicode.IsNumber(r[i]) {
			i++
		}
		for j > 0 && !unicode.IsLetter(r[j]) && !unicode.IsNumber(r[j]) {
			j--
		}
		if i >= j {
			return true
		}
		if r[i] == r[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
