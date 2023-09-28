package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Считываем флаги
	timeout := *flag.Duration("timeout", 10*time.Second, "timeout for connection")
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Println("Использование:  go run main.go [--timeout=<timeout>] <host> <port>")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	err := setConnection(host, port, timeout)
	if err != nil {
		fmt.Println("Ошибка :" + err.Error())
	}
}

func setConnection(host string, port string, timeout time.Duration) error {
	// Подключение к сокету
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return fmt.Errorf("Не удалось подключиться к сокету: %v\n" + err.Error())
	}
	defer conn.Close()

	// Установка обработчика сигналов для завершения программы при получении Ctrl+C
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// Закрывать соединение при нажатии комбинации Ctrl+C
		<-signals
		fmt.Println("\n Закрытие соединения...")
		conn.Close()
		os.Exit(0)
	}()

	// Копирование данных из STDIN в сокет
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Ошибка копирования из STDIN в сокет: ", err.Error())
		}
	}()

	// Копирование из сокета в STDOUT
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		return fmt.Errorf("Ошибка копирования из сокета в STDOUT: %v\n" + err.Error())
	}
	return nil
}
