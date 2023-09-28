package main

import "fmt"

// Интерфейс уровней
type Level interface {
	execute(*Player)
	setNext(Level)
}

// Конкретный уровень 1
type Level1 struct {
	next Level
}

// Прохождение уровня 1, в котором происходит передача управления следующему уровню по цепочке
func (l1 *Level1) execute(p *Player) {
	if p.level1Done {
		fmt.Println("Уровень 1 уже пройден")
		l1.next.execute(p)
		return
	}
	fmt.Println("Прохождение уровня 1")
	p.level1Done = true
	p.experience += 25
	fmt.Printf("Текущий опыт игрока %s: %v \n", p.nickname, p.experience)

	l1.next.execute(p)
}

// Установление следующего уровня
func (l1 *Level1) setNext(next Level) {
	l1.next = next
}

type Level2 struct {
	next Level
}

func (l2 *Level2) execute(p *Player) {
	if p.level2Done {
		fmt.Println("Уровень 2 уже пройден")
		l2.next.execute(p)
		return
	}
	fmt.Println("Прохождение уровня 2")
	p.level2Done = true
	p.experience += 50
	fmt.Printf("Текущий опыт игрока %s: %v \n", p.nickname, p.experience)

	l2.next.execute(p)
}

func (l2 *Level2) setNext(next Level) {
	l2.next = next
}

type Level3 struct {
	next Level
}

func (l3 *Level3) execute(p *Player) {
	if p.level3Done {
		fmt.Println("Уровень 3 уже пройден")
		l3.next.execute(p)
		return
	}
	fmt.Println("Прохождение уровня 3")
	p.level3Done = true
	p.experience += 70
	fmt.Printf("Текущий опыт игрока %s: %v \n", p.nickname, p.experience)

	l3.next.execute(p)
}

func (l3 *Level3) setNext(next Level) {
	l3.next = next
}

type Level4 struct {
	next Level
}

func (l4 *Level4) execute(p *Player) {
	if p.level4Done {
		fmt.Println("Уровень 4 уже пройден")
	}
	fmt.Println("Прохождение уровня 4")
	p.level4Done = true
	p.experience += 100
	fmt.Printf("Текущий опыт игрока %s: %v \n", p.nickname, p.experience)
}

func (l4 *Level4) setNext(next Level) {
	l4.next = next
}

// Игрок
type Player struct {
	nickname   string
	experience int
	level1Done bool
	level2Done bool
	level3Done bool
	level4Done bool
}

func main() {
	// Создание объекта уровня 4
	level4 := &Level4{}

	// Создание объекта уровня 3
	level3 := &Level3{}
	// Установка следующего уровня для уровня 3
	level3.setNext(level4)

	// Создание объекта уровня 2
	level2 := &Level2{}
	// Установка следующего уровня для уровня 2
	level2.setNext(level3)

	// Создание объекта уровня 1
	level1 := &Level1{}
	// Установка следующего уровня для уровня 1
	level1.setNext(level2)

	// Создание объекта игрока
	player := &Player{nickname: "Ivan123"}
	fmt.Printf("Начальный опыт игрока %s: %v \n", player.nickname, player.experience)

	// Пройти первый уровень
	level1.execute(player)
}

/* Паттерн "цепочка обязанностей" следует применять в том случае, когда есть несколько объектов, которые могут обработать запрос,
и вы хотите передать запрос только одному из них.

Плюсы паттерна:
- Уменьшает зависимость между клиентом и обработчиками
- Упрощает добавление новых обработчиков в цепочку без изменения клиентского кода
- Позволяет реализовать различные варианты обработки запросов.
Минусы паттерна:
- Запрос может остаться не обработанным
- Увеличивает сложность системы, особенно когда имеется большое количество обработчиков или когда цепочки становятся длинными и сложными
*/
