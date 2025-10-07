package main

import (
	"fmt"
	"os"

	//"sort"
	"bufio"
	"strings"
	"unicode"
)

// TextStats хранит всю статистику по тексту
type TextStats struct {
	charCount  int
	wordCount  int
	lineCount  int
	spaceCount int
}

func main() {
	//1. Проверить, что имя файла было передано
	if len(os.Args) != 2 {
		fmt.Println("Введите название файла, команда: go run main.go <название файла>")
		os.Exit(1)
	}
	filename := os.Args[1]

	//2. Анализировать файл
	stats, err := analyzeFile(filename)
	if err != nil {
		fmt.Println("Ошибка анализа файла:", err)
		os.Exit(1)
	}

	//3. Вывести основную статистику
	printStats(stats, filename)
}

// analyzeFile читает файл по указанному пути
// и возвращает структуру со статистикой
func analyzeFile(fileName string) (TextStats, error) {
	//Открыть файл для чтения
	file, err := os.Open(fileName)
	if err != nil {
		return TextStats{}, err
	}
	defer file.Close() //Гарантируем закрытие файла

	//Инициализируем структуру для хранения статистики
	stats := TextStats{}

	//Создать сканер для чтения файла по словам
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)
	//Читать файл слово за словом
	for scanner.Scan() {
		word := scanner.Text()
		stats.wordCount++

		cleanedWord := strings.ToLower(word)
		cleanedWord = strings.TrimFunc(cleanedWord, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
	}

	//Переоткрыть файл
	file.Seek(0, 0)
	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		stats.lineCount++
		stats.charCount += len(lineScanner.Text())
		//Считать пробелы в строке
		for _, char := range lineScanner.Text() {
			if unicode.IsSpace(char) {
				stats.spaceCount++
			}
		}
	}

	stats.charCount += stats.lineCount

	//Проверить на ошибки во время сканирования
	if err := scanner.Err(); err != nil {
		return TextStats{}, err
	}
	if err := lineScanner.Err(); err != nil {
		return TextStats{}, err
	}

	return stats, nil

}

// printStats выводит общую статистику в консоль
// Красиво выводим статистику на экран
func printStats(stats TextStats, filename string) {
	fmt.Printf("---Анализ файла ---%s\n", filename)
	fmt.Printf("Количество строк:                %d\n", stats.lineCount)
	fmt.Printf("Количество слов:                 %d\n", stats.wordCount)
	fmt.Printf("Количество символов:             %d\n", stats.charCount)
	fmt.Printf("Количество пробельных символолв: %d\n", stats.spaceCount)
}
