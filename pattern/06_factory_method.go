package main

import "fmt"

// Интерфейс автомобиля
type IAuto interface {
	setModel(model string)
	setSpeed(speed int)
	getModel() string
	getSpeed() int
}

// Основной класс автомобиля
type Auto struct {
	model    string
	maxSpeed int
}

func (a *Auto) setModel(model string) {
	a.model = model
}

func (a *Auto) getModel() string {
	return a.model
}

func (a *Auto) setSpeed(speed int) {
	a.maxSpeed = speed
}

func (a *Auto) getSpeed() int {
	return a.maxSpeed
}

// Фабричный метод, который позволяет получить готовый объект по типу
func getAuto(autoType string) (IAuto, error) {
	if autoType == "Car" {
		return newCar(), nil
	}
	if autoType == "Truck" {
		return newTruck(), nil
	}
	return nil, fmt.Errorf("Неправильный тип автомобиля")
}

// Легковой автомобиль
type Car struct {
	Auto
}

// Грузовик
type Truck struct {
	Auto
}

// Метод создания объекта грузовика
func newTruck() IAuto {
	return &Truck{
		Auto: Auto{
			model:    "ВАЗ-2108",
			maxSpeed: 100,
		},
	}
}

// Метод создания объекта легкового автомобиля
func newCar() IAuto {
	return &Car{
		Auto: Auto{
			model:    "ЗИЛ-157",
			maxSpeed: 200,
		},
	}
}
func main() {
	// Получить объект автомобиля
	car, _ := getAuto("Car")
	// Получить объект грузовика
	truck, _ := getAuto("Truck")
	fmt.Printf("Легковой автомобиль: Модель - %s, Максимальная скорость - %v км/ч\n", car.getModel(), car.getSpeed())
	fmt.Printf("Грузовик: %s, Максимальная скорость - %v км/ч\n", truck.getModel(), truck.getSpeed())
}

/*
Паттерн "фабричный метод" следует использовать в том случае, когда заранее неизвестны типы и зависимости объектов,
с которыми должен работать код
Плюсы паттерна:
- Упрощает добавление новых продуктов без изменения клиентского кода.
- Разделяет создание объектов и их использование, что позволяет сделать систему более гибкой и расширяемой.
Минусы паттерна:
- Вводит больше классов в систему, что увеличивает ее сложность.
*/
