// https://leetcode.com/problems/palindrome-partitioning/

// Given a string s, partition s such that every
// substring
//  of the partition is a
// palindrome
// . Return all possible palindrome partitioning of s.
// Example 1:

// Input: s = "aab"
// Output: [["a","a","b"],["aa","b"]]

// Example 2:

// Input: s = "a"
// Output: [["a"]]

package PalindromePartitioning

func partition(s string) [][]string {
	str := []string{}
	for _, v := range []rune(s) {
		str = append(str, string(v))
	}
	cur := []string{}
	res := [][]string{}
	Par(str, &cur, &res, len(s))
	return res
}

func Par(s []string, cur *[]string, res *[][]string, size int) {
	for n := len(s); n > 0; n-- {
		tmp := s[:n]
		if isPalidrome(tmp) {
			*cur = append(*cur, getString(tmp))
			if size == count(*cur) {
				*res = append(*res, append([]string{}, *cur...))
				*cur = (*cur)[:len(*cur)-1]
				continue
			} else {
				if len(tmp) < len(s) {
					s = s[len(tmp):]
					Par(s, cur, res, size)
					s = append(tmp, s...)
				}
			}
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}

func isPalidrome(s []string) bool {
	if len(s) == 1 {
		return true
	}
	for idx := 0; idx < len(s)/2; idx++ {
		if s[idx] != s[len(s)-idx-1] {
			return false
		}
	}
	return true
}

func getString(tmp []string) string {
	str := ""
	for _, v := range tmp {
		str += v
	}
	return str
}

func count(cur []string) int {
	cnt := 0
	for _, v := range cur {
		cnt += len(v)
	}
	return cnt
}
