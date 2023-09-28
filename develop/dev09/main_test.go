package main

import (
	"fmt"
	"testing"
)

func TestDownloadOK(t *testing.T) {
	url := "https://habr.com/ru/articles/"
	outputFileName := "habr.html"
	err := download(url, outputFileName)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

}

func TestDownloadError(t *testing.T) {
	url := "http://test123"
	outputFileName := "habr.html"

	err := download(url, outputFileName)

	if err == nil {
		t.Error("Ожидалась ошибка, но её нет")
	} else {
		fmt.Println(err.Error())
	}

}
func TestCreateFileOK(t *testing.T) {
	url := "https://habr.com/ru/articles/"
	outputFileName := "habr.html"

	err := download(url, outputFileName)
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}
}
func TestCreateFileError(t *testing.T) {
	url := "https://habr.com/ru/articles/"
	outputFileName := "*("

	err := download(url, outputFileName)
	if err == nil {
		t.Error("Ожидалась ошибка, но её нет")
	} else {
		fmt.Println(err.Error())
	}

}
