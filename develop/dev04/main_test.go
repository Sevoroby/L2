package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	words := []string{"слиток", "столик", "тяпка", "листок", "пятак", "тяпка", "глыба", "пятка"}

	anagramSets := findAnagramSets(words)

	expected := map[string][]string{
		"слиток": {"листок", "слиток", "столик"},
		"тяпка":  {"пятак", "пятка", "тяпка"},
	}

	if !reflect.DeepEqual(anagramSets, expected) {
		t.Errorf("Ожидаемый результат: %+v, полученный результат: %+v", expected, anagramSets)
	}
}
