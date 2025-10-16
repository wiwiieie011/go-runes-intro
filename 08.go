package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Интукод"
	fmt.Println(SubstrRunes(s, 2, 4))

}

// Функция SubstrRunes(s string, start, length int) string — вернуть подстроку по индексам рун. Проверь на "Интукод": возьми 4 руны, начиная с руны с индексом 2. Должно получиться туко.
func SubstrRunes(s string, start, length int) string {
	c:= []rune(s)
	var bildString strings.Builder
	for i:= start ; i < length; i++ {
		bildString.WriteRune(c[i]) 
	}
	return  bildString.String()
}