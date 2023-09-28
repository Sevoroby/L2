package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	// Канал для завершения
	done := make(chan interface{})
	// Запуск в цикле горутин, в которых будут прослушиваться входные каналы
	for _, ch := range channels {
		go func(ch1 <-chan interface{}) {
			// Ожидание записи в канал
			<-ch1
			fmt.Println("Вычитано значение из канала")
			// Закрыть основной канал, чтобы продолжить выполнение в горутине main
			close(done)
		}(ch)
	}
	return done
}

func main() {
	// Проверочная функция
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	// Начало отсчёта
	start := time.Now()
	// Вычитание из первого закрывшегося канала
	<-or(
		sig(2*time.Hour),
		sig(30*time.Second),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(4*time.Second),
		sig(1*time.Minute),
	)

	fmt.Printf("Завершение спустя %v", time.Since(start))
}
