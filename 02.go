package main

import "fmt"

func main() {
	// 	Задача 2
	// Попробуй объявить руну так: var bad rune = "A" (двойные кавычки).
	// Запусти, перепиши в комментариях текст ошибки и объясни, почему так нельзя.
	// потому что руну объявляются в одинарных кавычках
	// Исправь на var good rune = 'A'.

	// var bad rune = "А"
	var good rune = 'A'
	fmt.Println(good)

}