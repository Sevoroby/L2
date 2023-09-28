package main

import (
	"testing"
)

func TestExecuteCommandError(t *testing.T) {
	err := executeCommand("test123")
	if err == nil {
		t.Errorf("Ожидалась ошибка: Неверная команда, но её нет")
	}
}
