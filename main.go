package main

import (
	"fmt"
	m "go-leet/ValidAnagram"
)

func main() {
	s := "anagram"
	t := "nagaram"
	res := m.Is(s, t)
	fmt.Println(res)
}
