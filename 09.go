package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

fmt.Println(CleanPassword("пароль🙂 "))

}

// Задача 9. Очистка пароля
// Пользователь вводит пароль. Нужно:
// Удалить все пробельные символы (используй unicode.IsSpace).
// Запретить непечатные символы. (используй unicode.IsControl().
// Перевести латинские буквы в верхний регистр.
// Сохранить остальные руны как есть (в т.ч. эмодзи, кириллицу, знаки).
// Функция CleanPassword(s string) string. Обрати внимание, что функция принимает string и возвращает также string.

func CleanPassword(s string) string {
	b:= ""
	for _, v := range s {
		if  unicode.IsSpace(v) == false &&unicode.IsControl(v) == false{
				b += string(v) 
		}
	}
	c:= strings.ToUpper(b)
	return  c
}