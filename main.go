package main

import (
	"fmt"
	m "go-leet/LetterTilePossibilities"
)

func main() {

	tiles := "AAB"
	r := []rune(tiles)
	tile := [26]int{}
	for _, v := range r {
		tile[v-'A'] += 1
	}
	sum := m.NumTilePossibilities(tile)
	fmt.Println(sum)
}
