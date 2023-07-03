package main

import (
	"fmt"
	m "go-leet/PalindromePartitioning"
)

func main() {
	s := "aab"
	str := []string{}
	for _, v := range []rune(s) {
		str = append(str, string(v))
	}
	cur := []string{}
	res := [][]string{}
	m.Par(str, &cur, &res, len(s))
	fmt.Println(res)
}
