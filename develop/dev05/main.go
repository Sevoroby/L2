package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Keys - Ключи, которые подаются при запуске программы
type Keys struct {
	After        int
	Before       int
	Context      int
	Count        bool
	IgnoreCase   bool
	Invert       bool
	Fixed        bool
	PrintLineNum bool
}

func parseArgs(args []string) Keys {
	keys := Keys{}
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-A":
			keys.After, _ = strconv.Atoi(args[i+1])
			i++
		case "-B":
			keys.Before, _ = strconv.Atoi(args[i+1])
			i++
		case "-C":
			keys.Context, _ = strconv.Atoi(args[i+1])
			i++
		case "-c":
			keys.Count = true
		case "-i":
			keys.IgnoreCase = true
		case "-v":
			keys.Invert = true
		case "-F":
			keys.Fixed = true
		case "-n":
			keys.PrintLineNum = true
		}
	}

	return keys
}

func filterFile(options Keys, pattern string, filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Читаем входной файл построчно
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// совпадения по паттерну
	matches := make([]string, 0)
	count := 0
	// Если флаг -F, то ищем точное совпадение со строкой
	if options.Fixed {
		pattern = "^" + pattern + "$"
	}
	// Если флаг -i, то не учитываем регистр
	if options.IgnoreCase {
		pattern = "(?i)" + pattern
	}
	// Компиляция регулярного выражения по паттерну
	regExp, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	// Строки, идущие после искомой
	var afterLines []string
	// Строки, идущие до искомой
	var beforeLines []string

	// Сканирование строк
	for scanner.Scan() {
		line := scanner.Text()
		// Поиск совпадения по паттерну
		match := regExp.MatchString(line)
		// Если флаг -v, то вместо совпадения по паттерну ищем исключения
		if options.Invert {
			match = !match
		}
		// Если совпадение
		if match {
			count++
			if options.After > 0 {
				afterLines = nil
				matches = append(matches, line)
			} else if options.Before > 0 {
				matches = append(matches, beforeLines[len(beforeLines)-options.Before:]...)
				matches = append(matches, line)
				beforeLines = nil
			} else if options.Context > 0 {
				matches = append(matches, beforeLines[len(beforeLines)-options.Context:]...)
				matches = append(matches, line)
				beforeLines = nil
				afterLines = nil
			} else {
				// Если флаг -n. то выводим номер строки
				if options.PrintLineNum {
					line = fmt.Sprintf("%d:%s", count, line)
				}
				matches = append(matches, line)
			}
		} else { // Если не нашли совпадение
			// Если флаг -A, то к совпадению добавляем следующие за ним строки
			if options.After > 0 {
				afterLines = append(afterLines, line)
				if len(afterLines) == options.After && len(matches) != 0 {
					matches = append(matches, afterLines...)
				}
			}
			// Если флаг -B, то к совпадению добавляем предыдущие строки
			if options.Before > 0 {
				beforeLines = append(beforeLines, line)
			}

			// Если флаг -C, то к совпадению добавляем предыдущие и следующие строки
			if options.Context > 0 {
				beforeLines = append(beforeLines, line)
				afterLines = append(afterLines, line)
				if len(afterLines) == options.Context && len(matches) != 0 {
					matches = append(matches, afterLines...)
				}
			}
		}
	}

	// Если флаг -c, то выводим количество совпадений
	if options.Count {
		res := strconv.Itoa(count)
		return []string{res}, nil
	}
	return matches, nil
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Использование: go run main.go [flags] pattern filename")
		return
	}
	// Парсинг аргументов
	options := parseArgs(args)
	filename := args[len(args)-1]
	pattern := args[len(args)-2]

	// Фильтрация файла
	matches, err := filterFile(options, pattern, filename)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	// Вывод совпадений
	for _, match := range matches {
		fmt.Println(match)
	}
}
