package main

import (
	"fmt"
)

// Интерфейс состояния
type State interface {
	execute(*Thread)
}

// Конкретное состояние "ожидание"
type Waiting struct{}

// Выполнение состояния и переход к следующему
func (w *Waiting) execute(t *Thread) {
	fmt.Println("Выполняется состояние \"ожидание\"")
	t.setState(&Runnable{})
}

// Конкретное состояние "готов к запуску"
type Runnable struct{}

func (r *Runnable) execute(t *Thread) {
	fmt.Println("Выполняется состояние \"готов к запуску\"")
	t.setState(&Running{})
}

// Конкретное состояние "запущен"
type Running struct{}

func (r *Running) execute(t *Thread) {
	fmt.Println("Выполняется состояние \"запущен\"")
}

// Процесс, который использует состояния
type Thread struct {
	state State
}

// Установить состояние потока
func (t *Thread) setState(state State) {
	t.state = state
}

// Выполнить текущее состояние потока
func (t *Thread) executeState() {
	t.state.execute(t)
}

func main() {
	// Создание объекта потока
	thread := Thread{}
	// Создание объекта начального состояния
	waitingState := &Waiting{}
	// Установка начального состояния
	thread.setState(waitingState)
	// Выполнение текущего состояния
	thread.executeState()
	thread.executeState()
	thread.executeState()
}

/*
Паттерн "состояние" следует использовать в том случае, когда у вас есть объект,
поведение которого кардинально меняется в зависимости от внутреннего состояния,
причём типов состояний много, и их код часто меняется.
Плюсы паттерна:
- Избавляет от множества больших условных операторов машины состояний
- Позволяет легко добавлять новые состояния и изменять существующие без изменения кода контекста.
Минусы паттерна:
- Может неоправданно усложнить код, если состояний мало и они редко меняются
- Вводит дополнительные сложности в управлении состояниями объекта
*/
