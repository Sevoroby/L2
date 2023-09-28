package main

import (
	"testing"
)

func TestUnpackStr(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}

	for _, test := range tests {
		result, _ := unpackStr(test.input)

		if result != test.expected {
			t.Errorf("Ожидалось: %s, но получено: %s", test.expected, result)
		}
	}
}

func TestUnpackStrError(t *testing.T) {

	tests := []struct {
		input    string
		expected string
	}{
		{"45", "некорректная строка"},
	}

	for _, test := range tests {
		_, err := unpackStr(test.input)

		if err.Error() != "некорректная строка" {
			t.Errorf("Ожидалась ошибка: некорректная строка, но её нет")
		} else {
			t.Log(err.Error())
		}
	}
}
