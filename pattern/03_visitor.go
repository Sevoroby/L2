package main

import "fmt"

// Интерфейс фруктов - основного объекта
type FruitAccepter interface {
	Accept(Visitor)
}
type Fruit struct {
	price int
}
type Banana struct {
	Fruit
}

// Метод для передачи объекта посетителя и использования его
func (c *Banana) Accept(v Visitor) {
	v.VisitBananas(c)
}

type Apple struct {
	Fruit
}

func (r *Apple) Accept(v Visitor) {
	v.VisitApples(r)
}

type Orange struct {
	Fruit
}

func (t *Orange) Accept(v Visitor) {
	v.VisitOranges(t)
}

// Интерфейс посетителя, который обходит все необходимые структуры
type Visitor interface {
	VisitBananas(*Banana)
	VisitApples(*Apple)
	VisitOranges(*Orange)
}

// Конкретный посетитель для покупки фруктов
type BuyVisitor struct {
}

func (bv *BuyVisitor) VisitBananas(b *Banana) {
	fmt.Printf("Покупка бананов по цене %v \n", b.price)
}

func (bv *BuyVisitor) VisitApples(a *Apple) {
	fmt.Printf("Покупка яблок по цене %v \n", a.price)
}

func (bv *BuyVisitor) VisitOranges(o *Orange) {
	fmt.Printf("Покупка апельсинов по цене %v \n", o.price)
}

type ReducePriceVisitor struct {
}

func (rpv *ReducePriceVisitor) VisitBananas(c *Banana) {
	fmt.Printf("Текущая стоиомость бананов: %v \n", c.price)
	c.price -= 10
	fmt.Printf("Изменённая стоиомость бананов: %v \n", c.price)
}

func (rpv *ReducePriceVisitor) VisitApples(a *Apple) {
	fmt.Printf("Текущая стоиомость яблок: %v \n", a.price)
	a.price -= 10
	fmt.Printf("Изменённая стоиомость яблок: %v \n", a.price)
}

func (rpv *ReducePriceVisitor) VisitOranges(o *Orange) {
	fmt.Printf("Текущая стоиомость апельсинов: %v \n", o.price)
	o.price -= 10
	fmt.Printf("Изменённая стоиомость апельсинов: %v \n", o.price)
}

func main() {
	// Создание массива фруктов
	fruits := []FruitAccepter{
		&Banana{Fruit{price: 100}},
		&Apple{Fruit{price: 70}},
		&Orange{Fruit{price: 80}},
	}

	// Создание посетителя для изменения цены на фрукты
	visitor1 := &ReducePriceVisitor{}
	for _, fruit := range fruits {
		fruit.Accept(visitor1)
	}
	// Создание посетителя для покупки фруктов
	visitor2 := &BuyVisitor{}
	for _, fruit := range fruits {
		fruit.Accept(visitor2)
	}
}

/* Паттерн "посетитель" следует использовать в том случае, когда нужно добавить новую функциональность структуре, не изменяя её.

Плюсы паттерна "посетитель":
- Позволяет добавлять новые операции к объектам без изменения их классов.
- Упрощает добавление операций, работающих со сложными структурами объектов.
- Может накапливать состояние при обходе структуры элементов

Минусы паттерна "посетитель":
- Использование паттерна не оправдано, если иерархия элементов часто меняется
- Может привести к нарушению инкапсуляции элементов
*/
