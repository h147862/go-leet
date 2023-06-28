package main

import (
	"fmt"
	m "go-leet/LetterCasePermutation"
	"strings"
	"unicode"
)

func main() {
	s := "a1b2"
	res := []string{}
	alphabet := []m.Alphabet{}
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
		alphabet = append(alphabet, m.Alphabet{
			Addr:   addr,
			StrArr: al,
		})
	}
	m.LetterCasePermutation(&strArr, alphabet, &res, 0)
	fmt.Println(res)
}
