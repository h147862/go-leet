package main

import (
	"fmt"
	m "go-leet/LetterCombinationsofaPhoneNumber"
)

func main() {

	digits := "22"
	res := []string{}
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
	m.LetterCombinations(len(ch), used, &res, 0, &cur)
	fmt.Println(res)
}
