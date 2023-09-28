package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Line - Данные о строке
type Line struct {
	StrValue     string
	SortedColumn string  // столбец для сортировки
	NumValue     float64 // числовое представление строки
}

func parseLines(filename string, flag string, flagvalue string) ([]*Line, error) {
	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Создаём сканер файла
	scanner := bufio.NewScanner(file)
	lines := make([]*Line, 0)

	// Сканируем построчно файл
	for scanner.Scan() {
		value := scanner.Text()
		// По возможности парсим строку в числовом виде
		numVal, _ := strconv.ParseFloat(value, 32)
		// Добавляем в срез Line данные по каждой строке
		lines = append(lines, &Line{
			StrValue:     value,
			SortedColumn: getSortKey(value, flag, flagvalue),
			NumValue:     numVal,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getSortKey(value string, flag string, flagvalue string) string {
	// Делим строку на столбцы
	fields := strings.Fields(value)
	sortKey := fields[0]
	// Если указан флаг -k, то выбираем необходимый столбец для сортировки
	if len(fields) > 1 && flag == "-k" {
		flagValueInt, _ := strconv.Atoi(flagvalue)
		sortKey = fields[flagValueInt]
	}

	return sortKey
}

func checkKeys(lines []*Line, keys []string) []*Line {
	// Проходимся по срезу Lines
	for _, key := range keys {
		switch key {
		// Если указан флаг -r, то выполняем сортировку в обратном порядке
		case "-r":
			sort.Slice(lines, func(i, j int) bool {
				return lines[i].StrValue > lines[j].StrValue
			})
		// Если указан флаг -u, то сортируем без повторяющихся строк
		case "-u":
			uniqueLines := make([]*Line, 0)
			seen := make(map[string]bool)
			for _, line := range lines {
				if !seen[line.StrValue] {
					uniqueLines = append(uniqueLines, line)
					seen[line.StrValue] = true
				}
			}
			lines = uniqueLines
		// Если указан флаг -b, то сортируем строки без хвостовых пробелов
		case "-b":
			for _, line := range lines {
				for string(line.StrValue[len(line.StrValue)-1]) == " " {
					line.StrValue = strings.TrimSuffix(line.StrValue, " ")
				}
			}
		// Если указан флаг -h, то сортируем строки по числовым значениям с учётом суффиксов
		case "-h":
			for _, line := range lines {
				numStr := strings.TrimSuffix(line.StrValue, " ")
				numSuffix := string(numStr[len(numStr)-1])
				numValue, _ := strconv.ParseFloat(numStr[:len(numStr)-1], 32)
				switch numSuffix {
				// Если суффикс K, то умножаем число на 1000
				case "K":
					numValue *= 1000
				// Если суффикс M, то умножаем число на 1000000
				case "M":
					numValue *= 1000000
				// Если суффикс M, то умножаем число на 1000000000
				case "G":
					numValue *= 1000000000
				}
				line.NumValue = numValue

			}
			sort.Slice(lines, func(i, j int) bool {
				return lines[i].NumValue < lines[j].NumValue
			})
		// Если указан флаг -k, то сортируем строки по указанному столбцу
		case "-k":
			sort.Slice(lines, func(i, j int) bool {
				return lines[i].SortedColumn < lines[j].SortedColumn
			})
		// Если указан флаг -n, то сортируем строки по числовому значению
		case "-n":
			sort.Slice(lines, func(i, j int) bool {
				return lines[i].NumValue < lines[j].NumValue
			})
		// Если указан флаг -M, то сортируем строки по месяцам
		case "-M":
			sort.Slice(lines, func(i, j int) bool {
				t1, _ := time.Parse("January", strings.TrimSuffix(lines[i].StrValue, " "))
				t2, _ := time.Parse("January", strings.TrimSuffix(lines[j].StrValue, " "))
				return t1.Before(t2)
			})
		}
	}

	return lines
}

func writeLines(lines []*Line, filename string) error {
	// Создём файл для записи результатов
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// Записываем строки
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line.StrValue)
	}
	return writer.Flush()
}

func sortFile(inputFile, outputFile string, flags []string) error {
	var lines []*Line
	var err error
	// Проверка для флага -k
	if len(flags) > 1 {
		// Парсим строки файла
		lines, err = parseLines(inputFile, flags[0], flags[1])
	} else {
		lines, err = parseLines(inputFile, "", "")
	}
	if err != nil {
		return err
	}
	// Если флагов нет, то выполняем простую сортировку
	if len(flags) == 0 {
		sort.Slice(lines, func(i, j int) bool {
			return lines[i].StrValue < lines[j].StrValue
		})
	} else {
		lines = checkKeys(lines, flags)
	}
	// Записываем результаты в выходной файл
	if err = writeLines(lines, outputFile); err != nil {
		return err
	}
	return nil
}

func main() {
	// Считывание флагов
	flags := os.Args[1:]
	if len(flags) < 1 {
		fmt.Println("Использование: go run main.go [flags] inputFile")
		return
	}
	// Файл для вывода результата
	outputFile := "output.txt"
	// Входной файл
	inputFile := flags[len(flags)-1]

	flags = flags[:len(flags)-1]
	// Сортировка строк файла
	err := sortFile(inputFile, outputFile, flags)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат сортировки записан в файл -", outputFile)
}
