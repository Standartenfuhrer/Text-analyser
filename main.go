package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

type TextStats struct {
	lineCount   int
	wordCount   int
	charCount   int
	letterCount int
	spaceCount  int
	wordFreq    map[string]int
}

type WordStat struct {
	Word  int
	Count int
}

func main() {
	fmt.Println("Аргументы командной строки")
	fmt.Println(os.Args)

	if len(os.Args) != 2 {
		fmt.Println("Пожалуйста, укажите путь к файлу")
		return
	}

	//Получаем путь к файлу
	filePath := os.Args[1]
	fmt.Println("Анализирую файл:", filePath)

	//Получить содержимое файла
	content, err := os.ReadFile(filePath)
	if err != nil {
		//Если есть ошибка = выводим ее и завершаем программу
		log.Fatal(err)
	}

	//Вывести содержимое файла
	fmt.Println("Содержимое файла:")
	fmt.Println(content)

	//Преобразовать срез байт в строку
	textContent := string(content)
	fmt.Println("Содержимое файла в виде текста:")
	fmt.Println(textContent)

	//Анализируем текс с помощью вункции
	stats := analizeText(textContent)

	//Смотрим результаты анализа
	printStats(stats)
}

func analizeText(content string) TextStats {
	var stats TextStats

	//Проверить на пустой файл
	if content == "" {
		return stats //Возвращаем пустую структуру
	}

	//Посчитать символы
	stats.charCount = len(content)

	//Посчитать строки
	stats.letterCount = strings.Count(content, "/n") + 1

	//Посчитать слова
	words := strings.Fields(content)
	for _, word := range words {
		if len(word) > 0 {
			stats.wordCount++
			stats.wordFreq[word]++
		}
	}

	//Посчитать количество букв и пробельных символов
	for _, r := range content {
		if unicode.IsLetter(r) {
			stats.letterCount++
		}
		if unicode.IsSpace(r) {
			stats.spaceCount++
		}
	}

	return stats
}

func getTopWords(freqMap map[string]int, topN int) []WordStat {
	var WordStats []WordStat
	for word, count := range freqMap {
		WordStats = append(WordStats, WordStat{Word: word, Count: count})
	}

	sort.Slice(WordStats, func(i, j int) bool {
		return WordStats[i].Count > WordStats[j].Count
	})

	if len(WordStats) < topN {
		return WordStats
	}

}

// Красиво выводим статистику на экраны
func printStats(stats TextStats) {
	fmt.Println("---Анализ текста---")
	fmt.Printf("Количество строк:                %d\n", stats.lineCount)
	fmt.Printf("Количество слов:                 %d\n", stats.wordCount)
	fmt.Printf("Количество символов:             %d\n", stats.charCount)
	fmt.Printf("Количество букв:                 %d\n", stats.letterCount)
	fmt.Printf("Количество пробельных символолв: %d\n", stats.spaceCount)
}

// Подключил гит на хате
