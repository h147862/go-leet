// https://leetcode.com/problems/letter-case-permutation/
// 784. Letter Case Permutation

// Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.

// Return a list of all possible strings we could create. Return the output in any order.

// Example 1:

// Input: s = "a1b2"
// Output: ["a1b2","a1B2","A1b2","A1B2"]

// Example 2:

// Input: s = "3z4"
// Output: ["3z4","3Z4"]

// Constraints:

// 1 <= s.length <= 12
// s consists of lowercase English letters, uppercase English letters, and digits.

package LetterCasePermutation

import (
	"strings"
	"unicode"
)

func letterCasePermutation(s string) []string {
	res := []string{}
	alphabet := []Alphabet{}
	strArr := []string{}
	for addr, r := range s {
		al := []string{}
		ascii := r - '0'
		str := string(r)
		strArr = append(strArr, str)
		if ascii >= 0 && ascii <= 9 {
			continue
		}
		al = append(al, str)
		tmp := ""
		if !unicode.IsUpper(r) {
			tmp = strings.ToUpper(str)
		} else {
			tmp = strings.ToLower(str)
		}
		al = append(al, tmp)
		alphabet = append(alphabet, Alphabet{
			Addr:   addr,
			StrArr: al,
		})
	}
	LetterCasePermutation(&strArr, alphabet, &res, 0)
	return res
}

func LetterCasePermutation(s *[]string, alphabet []Alphabet, res *[]string, start int) {
	if start == len(alphabet) {
		*res = append(*res, toString(*s))
		return
	}
	idx := start
	addr := alphabet[idx].Addr
	strArr := alphabet[idx].StrArr
	for _, v := range strArr {
		origin := (*s)[addr]
		(*s)[addr] = v
		LetterCasePermutation(s, alphabet, res, idx+1)
		(*s)[addr] = origin
	}
}

func toString(s []string) string {
	str := ""
	for _, v := range s {
		str += v
	}
	return str
}

type Alphabet struct {
	Addr   int
	StrArr []string
}
