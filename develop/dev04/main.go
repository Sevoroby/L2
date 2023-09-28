package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramSets(words []string) map[string][]string {
	// Мапа для уникальных слов
	uniqueWords := make(map[string]struct{})
	// Мапа с отсортированным ключом
	sortedWords := make(map[string][]string)
	// Итоговая мапа
	res := make(map[string][]string)

	for _, word := range words {
		// Проверка на повторяющиеся слова
		if _, inMap := uniqueWords[word]; !inMap {
			uniqueWords[word] = struct{}{}
		} else {
			continue
		}
		// Сортируем буквы слова в нижнем регистре
		sortedWord := sortString(strings.ToLower(word))
		// Добавляем в мапу с отсортированным ключом элемент, где ключ - отсортированное слово, а значение - срез исходных слов
		sortedWords[sortedWord] = append(sortedWords[sortedWord], word)

		fmt.Println(sortedWords)
	}

	// Проходимся по мапе с отсортированным ключом
	for _, sortWord := range sortedWords {
		// Если в множестве более одного элемента, то записываем в итоговую мапу
		if len(sortWord) > 1 {
			// Выставляем в итоговой мапе в качестве ключа первый элемент множества
			res[sortWord[0]] = sortWord
			// Сортировка среза слов по возрастанию
			sort.Strings(sortWord)
		}
	}

	return res
}

func sortString(s string) string {
	sortedRunes := []rune(s)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func main() {
	// Пример использования
	words := []string{"слиток", "столик", "тяпка", "листок", "пятак", "тяпка", "глыба", "пятка"}
	// Поиск множеств анаграмм
	anagramSets := findAnagramSets(words)

	for key, word := range anagramSets {
		fmt.Printf("Ключ: %s; Слова: %s\n", key, strings.Join(word, ", "))
	}
}
