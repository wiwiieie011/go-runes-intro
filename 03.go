package main

import "fmt"

func main() {
	// 	Задача 3
	// Создай строку s := "Жук".
	// Выведи s[0], s[1], s[2] как числа (байты).
	// Добавь цикл for i := 0; i < len(s); i++ { fmt.Println(i, s[i]) } — исследуй что печатается.
	// Добавь цикл for i, r := range s { fmt.Println(i, r, string(r)) } — сравни с предыдущим.
	s := "Жук"
	for i, _ := range s {
		fmt.Println(i, s[i])
	}

	for i, j := range s {
		fmt.Println(i,j, string(j))
	}
}