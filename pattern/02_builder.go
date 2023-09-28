package main

import "fmt"

// Основной объект - дом
type House struct {
	Walls   int
	Windows int
	Doors   int
	Pool    bool
}
type Director struct {
}

// Постройка дома по шагам
func (d *Director) BuildHouse(builder HouseBuilder) {
	builder.reset()
	builder.setWalls()
	builder.setDoors()
	builder.setWindows()
	builder.setPool()
}

// Интерфейс строителя
type HouseBuilder interface {
	reset()
	setWalls()
	setDoors()
	setWindows()
	setPool()
	getResult() *House
}

// Конкретный строитель малого дома
type SmallHouseBuilder struct {
	house *House
}

// Создание нового объекта
func (hb *SmallHouseBuilder) reset() {
	hb.house = &House{}
}

func (hb *SmallHouseBuilder) setWalls() {
	hb.house.Walls = 4
}

func (hb *SmallHouseBuilder) setDoors() {
	hb.house.Doors = 3
}

func (hb *SmallHouseBuilder) setWindows() {
	hb.house.Windows = 5
}
func (hb *SmallHouseBuilder) setPool() {
	hb.house.Pool = false
}

// Возврат построенного объекта
func (hb *SmallHouseBuilder) getResult() *House {
	return hb.house
}

// Конкретный строитель большого дома
type LargeHouseBuilder struct {
	house *House
}

func (hb *LargeHouseBuilder) reset() {
	hb.house = &House{}
}

func (hb *LargeHouseBuilder) setWalls() {
	hb.house.Walls = 50
}

func (hb *LargeHouseBuilder) setDoors() {
	hb.house.Doors = 10
}

func (hb *LargeHouseBuilder) setWindows() {
	hb.house.Windows = 60
}
func (hb *LargeHouseBuilder) setPool() {
	hb.house.Pool = true
}

func (hb *LargeHouseBuilder) getResult() *House {
	return hb.house
}

func main() {
	// Создание директора
	director := &Director{}

	// Создание строителя дома
	builder := &SmallHouseBuilder{}
	// Вызов метода директора дял постройки малого дома
	director.BuildHouse(builder)
	res := builder.getResult()
	fmt.Printf("Построен маленький дом: стены-%v штук,двери-%v штук, окна-%v штук, бассейн - %v\n", res.Walls, res.Doors, res.Windows, res.Pool)
	// Вызов метода директора для постройки большого дома
	builder2 := &LargeHouseBuilder{}
	director.BuildHouse(builder2)
	res = builder2.getResult()
	fmt.Printf("Построен большой дом: стены-%v штук,двери-%v штук, окна-%v штук, бассейн - %v\n", res.Walls, res.Doors, res.Windows, res.Pool)
}

/* Паттерн "строитель" следует применять в том случае, когда нужно создавать большой объект вручную и указывать в конструкторе множество
параметров для создания объекта. Паттерн позволяет упростить процесс создания объекта для клиентского кода

Плюсы паттерна:
- Позволет создавать объекты по шагам
- Упрощает использование кода создания сложного объекта

Минусы паттерна:
- Усложнение кода из-за новых классов
- Клиент привязан к конкретным классам строителей, т.к. в интерфейсе директора может не быть метода получения результата
*/
