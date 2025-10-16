package main

import (
	"fmt"
	"strings"
	// "unicode/utf8"
)

func main() {
s:= "Привет"

fmt.Println(ReverseRunes(s))

}

// Задача 7
// Напиши функцию ReverseRunes(s string) string, которая переворачивает строку по рунам (а не по байтам!). Проверь на "Привет", "你好", "🙂👍".
func ReverseRunes(s string) string {
	var  revesNumber  = []rune(s)
	// res := make([]rune, len(revesNumber))
	var c strings.Builder
	for i:= len(revesNumber)-1; i >=  0 ; i--{
		c.WriteRune(revesNumber[i])
	}
	return  c.String()

}