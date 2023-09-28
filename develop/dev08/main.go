package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func changeDirectory(args []string) {
	// Если аргументов нет, перейти в домашнюю директорию
	if len(args) == 1 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Не удалось получить домашнюю директорию:", err)
			return
		}
		err = os.Chdir(homeDir)
		if err != nil {
			fmt.Println("Не удалось изменить директорию:", err)
		}
	} else {
		// Иначе перейти в указанную директорию
		err := os.Chdir(args[1])

		if err != nil {
			fmt.Println("Не удалось изменить директорию:", err)

		}
	}
}

func printWorkingDirectory() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Не удалось получить текущую директорию:", err)
		return
	}
	fmt.Println(wd)
}

func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func kill(args []string) {
	if len(args) != 2 {
		fmt.Println("Использование: kill <pid>")
		return
	}
	var cmd *exec.Cmd
	// В зависимости от системы выбираем нужную команду и выполняем
	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill /PID " + args[1])
	} else {
		cmd = exec.Command("kill " + args[1])
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Ошибка при выполнении команды kill:", err)
		fmt.Println(string(out))
	}
}

func printProcessList() {
	var cmd *exec.Cmd
	// В зависимости от системы выбираем нужную команду и выполняем
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	} else {
		cmd = exec.Command("ps")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Ошибка при выполнении команды ps:", err)
		fmt.Println(string(out))
		return
	}
	fmt.Println(string(out))
}

func executeCommand(command string) error {
	args := strings.Split(command, " ")
	switch args[0] {
	case "cd":
		// Смена директории
		changeDirectory(args)
	case "pwd":
		// Вывод текущей директории
		printWorkingDirectory()
	case "echo":
		// Вывод аргументов в STDOUT
		echo(args)
	case "kill":
		// Убить таск
		kill(args)
	case "ps":
		// Вывести список процессов
		printProcessList()
	default:
		return fmt.Errorf("Неверная команда")
	}
	return nil
}

func main() {
	// Считывание ввода с клавиатуры
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		if command == "\\quit" {
			break
		}
		readLine(command)
	}
}

func readLine(command string) {
	// Если строка содержит символ '|', то делим строку на части
	if strings.Contains(command, "|") {
		commands := strings.Split(command, "|")

		// Проходимся по всем командам
		for _, cmd := range commands {
			cmd = strings.TrimSpace(cmd)

			// Выполняем команду
			executeCommand(cmd)

		}
	} else {
		// Если строка не содержит символ '|', то просто выполняем команду
		executeCommand(command)
	}
}
