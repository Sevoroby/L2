package main

import (
	"testing"
)

func TestParseSelectedColumns(t *testing.T) {
	input := "1,4-7,9,10"
	expected := []int{1, 4, 5, 6, 7, 9, 10}

	result := parseSelectedColumns(input)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Неверное значние, оижадось: %d, Получено: %d", expected[i], v)
		}
	}
}
