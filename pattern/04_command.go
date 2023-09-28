package main

import "fmt"

// Отправитель команды - человек
type Human struct {
	command Command
}

// Приказать выполнить команду животному
func (b *Human) makeOrder() {
	b.command.execute()
}

// Интерфейс команды
type Command interface {
	execute()
}

// Конкретная команда - издать голос
type VoiceCommand struct {
	animal Animal
}

// онкретная команда - лежать
type LayCommand struct {
	animal Animal
}

// Выполнение команды
func (c *VoiceCommand) execute() {
	c.animal.voice()
}
func (c *LayCommand) execute() {
	c.animal.lay()
}

// Интерфейс животного
type Animal interface {
	voice()
	lay()
}

// Конкретное животное - собака
type Dog struct {
}

func (d *Dog) voice() {
	fmt.Println("Собака лает")
}
func (c *Dog) lay() {
	fmt.Println("Собака лежит")
}

// Конкретное животное - кошка
type Cat struct {
}

func (c *Cat) voice() {
	fmt.Println("Кошка мяукает")
}
func (c *Cat) lay() {
	fmt.Println("Кошка лежит")
}

func main() {
	dog := &Dog{}
	cat := &Cat{}
	// Создание голосовой команды для собаки
	voiceCommandForDog := &VoiceCommand{
		animal: dog,
	}
	// Создание команды лежать для собаки
	layCommandForDog := &LayCommand{
		animal: dog,
	}
	// Создание голосовой команды для кошки
	voiceCommandForCat := &VoiceCommand{
		animal: cat,
	}
	// Создание команды лежать для кошки
	layCommandForCat := &LayCommand{
		animal: cat,
	}
	// Создание отправителя команды - человека
	human := &Human{}
	// установление команды, который человек должен приказать животному
	human.command = voiceCommandForDog
	// Отдача приказа животному на выполнение команды
	human.makeOrder()

	human.command = layCommandForDog
	human.makeOrder()

	human.command = voiceCommandForCat
	human.makeOrder()

	human.command = layCommandForCat
	human.makeOrder()
}

/*
Паттерн "команда" следует использовать в случае, когда нужно чтобы объекты одного слоя приложения не вызывали
напрямую методы другого слоя, а делали это через специальный интерфейс "команду"

 Плюсы паттерна:
- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
- Упрощает добавление новых операций без изменения существующего кода
- Обеспечивает возможность построения сложных команд из простых операций
Минусы паттерна:
- Усложняет код программы из-за введения множества дополнительных классов.
*/
