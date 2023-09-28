package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(url string, outputFileName string) error {
	// Отправка Get-запроса на указанный url
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Ошибка при выполнении Get-запроса: %v", err)
	}
	defer resp.Body.Close()

	// Если статус ответа от сервера неверный
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Неверный статус ответа от сервера %v", resp.Status)
	}

	// Создание файла и запись в него результата
	err = createAndWriteOutputFile(resp.Body, outputFileName)
	if err != nil {
		return err
	}

	return nil
}
func createAndWriteOutputFile(body io.ReadCloser, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Ошибка при создании файла: %v", err)
	}
	// Копирование тела ответа от сервера в файл
	_, err = io.Copy(file, body)
	if err != nil {
		return fmt.Errorf("Ошибка при сохранении файла: %v", err)
	}

	fmt.Printf("Результат успешно сохранён в файле %s\n", filename)
	defer file.Close()
	return nil
}
func main() {

	if len(os.Args) != 2 {
		fmt.Println("Необходимо указать url")
		return
	}
	// Считываем url в качестве аргумента при запуске программы
	url := os.Args[1]
	filename := "output.html"
	err := download(url, filename)
	if err != nil {
		fmt.Println("Ошибка при скачивании страницы:", err)
	}
}
