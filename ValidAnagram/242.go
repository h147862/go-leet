// https://leetcode.com/problems/valid-anagram/
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:

// Input: s = "rat", t = "car"
// Output: false

package ValidAnagram

func isAnagram(s string, t string) bool {
	return Is(s, t)
}

func Is(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	str1 := []rune(s)
	str2 := []rune(t)
	cnt1 := make(map[string]int)
	cnt2 := make(map[string]int)
	for idx := 0; idx < len(str1); idx++ {
		s1 := string(str1[idx])
		s2 := string(str2[idx])
		if _, exist := cnt1[s1]; !exist {
			cnt1[s1] = 1
		} else {
			cnt1[s1] += 1
		}
		if _, exist := cnt2[s2]; !exist {
			cnt2[s2] = 1
		} else {
			cnt2[s2] += 1
		}
	}
	for k, v := range cnt1 {
		if v2, exist := cnt2[k]; !exist {
			return false
		} else if v2 != v {
			return false
		}
	}
	return true
}
