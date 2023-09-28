package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fields string
	var delimiter string
	var separated bool
	// Считывание флагов
	flag.StringVar(&fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "только строки с разделителем")
	flag.Parse()

	// Парсинг выбираемых полей
	selectedColumns := parseSelectedColumns(fields)

	// Чтение вводимых строк
	StdinReader := bufio.NewReader(os.Stdin)
	readLines(StdinReader, selectedColumns, delimiter, separated)
}

func parseSelectedColumns(fields string) []int {
	var res []int
	// Могут быть перечислены через запятую
	for _, field := range strings.Split(fields, ",") {
		// Чтение диапазонов строк
		if strings.Contains(field, "-") {
			rangeFields := strings.Split(field, "-")
			if len(rangeFields) == 2 {
				start, _ := parseField(rangeFields[0])
				end, _ := parseField(rangeFields[1])

				for i := start; i <= end; i++ {
					res = append(res, i)
				}
			}
		} else {
			column, _ := parseField(field)
			res = append(res, column)
		}
	}

	return res
}

func parseField(s string) (int, error) {
	field := strings.TrimSpace(s)
	col, err := strconv.Atoi(field)

	if err != nil || col <= 0 {
		return 0, fmt.Errorf("Некорректное поле: %s", field)
	}

	return col, nil
}

func readLines(reader *bufio.Reader, selectedColumns []int, delimiter string, separated bool) {
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			return
		}

		line = strings.TrimRight(line, "\n")

		// Если стоит флаг -s и в строке нет выбранного разделителя, то пропускаем строку
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		columns := strings.Split(line, delimiter)
		result := make([]string, len(selectedColumns))
		// Проходимся по выбранным полям
		for i, col := range selectedColumns {
			if col-1 < len(columns) {
				result[i] = columns[col-1]
			}
		}

		fmt.Println(strings.Join(result, delimiter))
	}
}
