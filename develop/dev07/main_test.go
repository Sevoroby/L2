package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {

	// Проверка, что функция завершает выполнение после первого события на входных каналах
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(30*time.Second),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(4*time.Second),
		sig(1*time.Minute),
	)

	elapsed := time.Since(start)
	seconds := (int)(elapsed.Seconds())
	if seconds != 4 {
		t.Errorf("Ожидалось время в секундах: 4, получено: %v", elapsed)
	}
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
