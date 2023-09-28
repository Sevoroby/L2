package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParseAndValidateEventId_Error(t *testing.T) {
	r := httptest.NewRequest("GET", "/?id=invalid", nil)
	_, err := parseAndValidateEventID(r)
	if err == nil {
		t.Error("Ожидалась ошибка, получено nil")
	}

	expectedError := "Параметр 'id' должен быть числом"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Неожиданная ошибка. Ожидалось: %q, Получено: %q", expectedError, err.Error())
	}
}

func TestParseAndValidateDate_Error(t *testing.T) {
	r := httptest.NewRequest("GET", "/?date=invalid", nil)
	_, err := parseAndValidateDate(r)
	if err == nil {
		t.Error("Ожидалась ошибка, получено nil")
	}

	expectedError := "Параметр 'date' должен быть в формате 'YYYY-MM-DD'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Неожиданная ошибка. Ожидалось: %q, Получено: %q", expectedError, err.Error())
	}
}

func TestParseAndValidateTitle_Error(t *testing.T) {
	r := httptest.NewRequest("POST", "/?title=", nil)
	_, err := parseAndValidateTitle(r)
	if err == nil {
		t.Error("Ожидалась ошибка, получено nil")
	}

	expectedError := "Параметр 'title' обязателен"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Неожиданная ошибка. Ожидалось: %q, Получено: %q", expectedError, err.Error())
	}
}

func TestParseAndValidateUserId_Error(t *testing.T) {
	r := httptest.NewRequest("GET", "/?user_id=invalid", nil)
	_, err := parseAndValidateUserID(r)
	if err == nil {
		t.Error("Ожидалась ошибка, получено nil")
	}

	expectedError := "Параметр 'user_id' должен быть числом"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Неожиданная ошибка. Ожидалось: %q, Получено: %q", expectedError, err.Error())
	}
}
