package main

import (
	"unicode"
	"fmt"
)

func main() {
a:= "Приве6763123123т"

fmt.Println(OnlyLetters(a))
	

}
// Задача 6
// Напиши функцию OnlyLetters(s string) string, которая возвращает новую строку только из букв. Используй for _, r := range s и unicode.IsLetter(r)
func  OnlyLetters(s string)string{
	var b string = ""
	for _, r := range s {
	if  unicode.IsLetter(r) == true{
		  b +=string(r)	
	}
	}
	return  b
}