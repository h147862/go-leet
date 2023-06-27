// https://leetcode.com/problems/letter-tile-possibilities/

// You have n  tiles, where each tile has one letter tiles[i] printed on it.

// Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles.

// Example 1:

// Input: tiles = "AAB"
// Output: 8
// Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".

// Example 2:

// Input: tiles = "AAABBC"
// Output: 188
// Example 3:

// Input: tiles = "V"
// Output: 1

// Constraints:

// 1 <= tiles.length <= 7
// tiles consists of uppercase English letters.

package LetterTilePossibilities

func numTilePossibilities(tiles string) int {
	r := []rune(tiles)
	tile := [26]int{}
	for _, v := range r {
		tile[v-'A'] += 1
	}
	return NumTilePossibilities(tile)
}

func NumTilePossibilities(tiles [26]int) int {
	sum := 0
	for idx, t := range tiles {
		if t == 0 {
			continue
		}
		sum += 1
		tiles[idx]--
		sum += NumTilePossibilities(tiles)
		tiles[idx]++
	}
	return sum
}
