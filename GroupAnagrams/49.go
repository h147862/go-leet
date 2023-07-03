// https://leetcode.com/problems/group-anagrams/

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Example 2:
// Input: strs = [""]
// Output: [[""]]

// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]

package GroupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	return Group(strs)
}

func Group(strs []string) [][]string {
	dict := make(map[string][]int)
	for idx, s := range strs {
		key := getKey(s)
		if _, exist := dict[key]; !exist {
			dict[key] = []int{idx}
		} else {
			dict[key] = append(dict[key], idx)
		}
	}
	res := [][]string{}
	for _, v := range dict {
		tmp := []string{}
		for _, idx := range v {
			tmp = append(tmp, strs[idx])
		}
		res = append(res, tmp)
	}
	return res
}

func getKey(s string) string {
	key := ""
	strs := []string{}
	for _, v := range s {
		strs = append(strs, string(v))
	}
	sort.Strings(strs)
	for _, r := range strs {
		v := []rune(r)
		key += string(v[0]-'a') + "-"
	}
	return key
}
