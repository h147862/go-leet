package main

import (
	"fmt"
	m "go-leet/GroupAnagrams"
)

func main() {
	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := m.Group(s)
	fmt.Println(res)
}
