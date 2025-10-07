package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main(){
	fmt.Println("Аргументы командной строки")
	fmt.Println(os.Args)

	if len(os.Args) != 2{
		fmt.Println("Пожалуйста, укажите путь к файлу")
		return
	}

	filepath := os.Args[1]
	fmt.Println("Анализирую файл:", filepath)

	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Содержимое файла:", content)

	textContent := string(content)
	fmt.Println("Содержмое файла в виде текста:", textContent)
	fmt.Println("Количество символов:", utf8.RuneCountInString(textContent))
	fmt.Println("Количество строк:", strings.Count(textContent, "\n") + 1)
	//dasdad
}	