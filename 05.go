package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println(CountRunes("你好"))
}

// Задача 5
// Напиши функцию CountRunes(s string) int,
// которая возвращает количество рун в строке. Проверь на строках: "Go", "Привет", "你好", "🙂"
func CountRunes(s string) int {
	return utf8.RuneCountInString(s)
}