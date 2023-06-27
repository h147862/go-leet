// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Example 1:

// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Example 2:

// Input: digits = ""
// Output: []
// Example 3:

// Input: digits = "2"
// Output: ["a","b","c"]

package LetterCombinationsofaPhoneNumber

func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}
	cur := []string{}
	r := []rune(digits)
	ch := []string{}
	for _, v := range r {
		ch = append(ch, string(v))
	}
	lookup := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	used := make(map[int]map[string]bool)
	for idx, c := range ch {
		used[idx] = make(map[string]bool)
		for _, v := range lookup[c] {
			used[idx][v] = false
		}
	}
	LetterCombinations(len(ch), used, &res, 0, &cur)
	return res
}

func LetterCombinations(length int, used map[int]map[string]bool,
	res *[]string, idx int, cur *[]string) {

	if idx == length {
		s := ""
		for _, v := range *cur {
			s += v
		}
		*res = append(*res, s)
		return
	}
	for d, isUsed := range used[idx] {
		if !isUsed {
			*cur = append(*cur, d)
			used[idx][d] = true
			LetterCombinations(length, used, res, idx+1, cur)
			*cur = (*cur)[:len(*cur)-1]
			used[idx][d] = false
		} else {
			continue
		}
	}
}
