package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackStr(str string) (string, error) {
	// используем strings.Builder, поскольку он лучше подходит для конкатенации строк
	var res strings.Builder
	// Количество повторений символа
	var cnt int
	// Предыдущий символ
	var prevSymbol rune
	// Признак, указывающий на экранированность предыдущего символа
	var isPrevShielded bool

	if len(str) == 0 {
		return str, nil
	}
	for _, symbol := range str {

		if unicode.IsDigit(symbol) { // Если текущий символ - цифра
			cnt, _ = strconv.Atoi(string(symbol))
			// Если первый символ цифра или встречается цифра 0 или идут две цифры подряд без экранирования, то возвращаем ошибку
			if prevSymbol == 0 || cnt == 0 || (unicode.IsDigit(prevSymbol) && isPrevShielded == false) {
				return "", fmt.Errorf("некорректная строка")
			}
		} else {
			cnt = 1
		}

		if prevSymbol == '\\' { // Если предыдущий символ \, то проставляем признак экранированности для текущего символа
			if isPrevShielded {
				// Записываем в результирующую строку символ \
				res.WriteString(strings.Repeat(string(prevSymbol), cnt))
				// Если предыдущий символ \ и он экранирован, то для текущего символа признак экранированности будет false
				isPrevShielded = false
			} else {
				isPrevShielded = true
			}
		} else if !unicode.IsDigit(prevSymbol) || (unicode.IsDigit(prevSymbol) && isPrevShielded) { // Иначе если предыдущий символ - буква или экранированная цифра
			// Записываем в результирующую строку предыдущий символ
			if prevSymbol != 0 {
				res.WriteString(strings.Repeat(string(prevSymbol), cnt))
			}
			isPrevShielded = false
		}
		prevSymbol = symbol
	}

	// Выводим последний символ, если это буква или экранированная цифра
	cnt = 1
	if !unicode.IsDigit(prevSymbol) {
		res.WriteString(strings.Repeat(string(prevSymbol), cnt))
	} else if unicode.IsDigit(prevSymbol) && isPrevShielded {
		res.WriteString(strings.Repeat(string(prevSymbol), cnt))
	}
	return res.String(), nil
}

func main() {
	str := "a4bc2d5e3"
	fmt.Println("Исходная строка:", str)
	res, err := unpackStr(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат:", res)
	}

	str = "abcd"
	fmt.Println("Исходная строка:", str)
	res, err = unpackStr(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат:", res)
	}

	str = "45"
	fmt.Println("Исходная строка:", str)
	res, err = unpackStr(str)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Результат:", res)
	}

	str = ""
	fmt.Println("Исходная строка:", str)
	res, err = unpackStr(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат:", res)
	}

	str = "qwe\\4\\5"
	fmt.Println("Исходная строка:", str)
	res, err = unpackStr(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат:", res)
	}

	str = "qwe\\45"
	fmt.Println("Исходная строка:", str)
	res, err = unpackStr(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат:", res)
	}

	str = "qwe\\\\5"
	fmt.Println("Исходная строка:", str)
	res, err = unpackStr(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат:", res)
	}
}
