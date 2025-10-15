package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
s:= "Привет"

fmt.Println(ReverseRunes(s))

}

// Задача 7
// Напиши функцию ReverseRunes(s string) string, которая переворачивает строку по рунам (а не по байтам!). Проверь на "Привет", "你好", "🙂👍".
func ReverseRunes(s string) string {
	var  revesNumber  = []rune(s)
	res := make([]rune, len(revesNumber))
	for i:= 0; i <= utf8.RuneCountInString(s)-1 ; i++{
		res[i] = revesNumber[len(revesNumber)-1-i]
	}
	return  string(res)

}