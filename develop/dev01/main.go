package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	// Запрос точного времени с помощью библиотеки NTP
	res, err := ntp.Time("pool.ntp.org")
	if err != nil {
		// Вывод ошибки в STDERR
		fmt.Fprintln(os.Stderr, err.Error())
		// Ненулевой код выхода в OS
		os.Exit(1)
	}
	// Вывод текущего времени стандартной функцией
	fmt.Println("Текущее время:", time.Now())
	// Вывод точного времени, полученного с помощью библиотеки NTP
	fmt.Println("Точное время:", res)
}
