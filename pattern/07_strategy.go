package main

import "fmt"

// Интерфейс стратегии приготовления еды
type CookingStrategy interface {
	cook(f *Food)
}

// Конкретная стратегия - жарка
type Frying struct {
}

// Метод жарки
func (fr *Frying) cook(f *Food) {
	fmt.Printf("%s жарится...\n", f.Name)
}

// Конкретная стратегия - варка на пару
type Steaming struct {
}

func (s *Steaming) cook(f *Food) {
	fmt.Printf("%s варится на пару...\n", f.Name)
}

// Конкретная стратегия - варение
type Boiling struct {
}

func (b *Boiling) cook(f *Food) {
	fmt.Printf("%s варится...\n", f.Name)
}

// Продукты
type Food struct {
	Name string
}

// Приготовитель продуктов
type FoodMaker struct {
	cookingStr CookingStrategy
}

// Установить стратегию для готовки
func (fm *FoodMaker) setCookingStrategy(ck CookingStrategy) {
	fm.cookingStr = ck
}

// Приготовить еду в заисимости от выбранной стратегии
func (fm *FoodMaker) cookFood(f *Food) {
	fm.cookingStr.cook(f)
}
func main() {
	// Создание объекта стратегии - жарка
	frying := &Frying{}
	// Создание объекта приготовителя еды
	foodMaker := &FoodMaker{}
	// Установить стратегию жарки
	foodMaker.setCookingStrategy(frying)
	// Создание объекта продуктов - мясо
	meat := &Food{Name: "Мясо"}
	// Приготовить еду
	foodMaker.cookFood(meat)

	// Создание объекта стратегии - варение
	boiling := &Boiling{}

	foodMaker.setCookingStrategy(boiling)
	// Создание объекта продуктов - картошка
	potato := &Food{Name: "Картошка"}
	foodMaker.cookFood(potato)
}

/*
 Паттерн "стратегия" следует использовать в том случае, когда различные варианты алгоритма представлены в виде длинных условных конструкций

Плюсы паттерна:
- Упрощает добавление новых стратегий без изменения клиентского кода.
- Изолирует код и данные алгоритмов от остальных классов.

Минусы паттерна:
- Увеличивает количество классов в системе, что может повысить ее сложность
- Может усложнить программу из-за дополнительной абстракции и введения зависимостей между классами стратегий и клиентским кодом.
*/
