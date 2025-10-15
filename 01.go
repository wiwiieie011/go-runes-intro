package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
var r rune = 'A'
s:= "Go!"
ru:= "Привет"
fmt.Println(r)
fmt.Println(string(r))
fmt.Println(len(s), utf8.RuneCountInString(s))
fmt.Println(utf8.RuneCountInString(ru),  len(ru))

}

// Создай новый проект на Go (go mod init ...). Все задания выполняй в этом проекте. Для каждой отдельной задачи создавай свой файл. Можно для всех использовать пакет main. Помни, что в пакете main может быть только одна функция main().

// Объяви переменную var r rune = 'A' и выведи её как число и как символ.
// Создай строку s := "Go!".
// Выведи: len(s) и результат utf8.RuneCountInString(s).
// Добавь строку с кириллицей: ru := "Привет" и снова сравни len(ru) и utf8.RuneCountInString(ru).