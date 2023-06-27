package main

import "fmt"

type Human struct {
	name string
	age  int
	Action
}

// Методы структуры Human
func (h Human) sayHello() {
	fmt.Printf("Hello, I am %s\n", h.name)
}

func (h Human) getAge() int {
	return h.age
}

type Action struct{}

// Метод пустой структуры Action, можем вызывать его из структуры Human, так как добавили поле типа Action в Human
// Можно добавить в другие структуры при необходимости
func (a Action) countToN(n int) {
	fmt.Println("Counting to", n)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

func main() {
	// Создание структуры Human
	h := Human{"Kirill", 22, Action{}}

	// Вызовы методов Human
	h.sayHello()
	fmt.Println(h.getAge())

	// Вызов метода структуры Action
	h.countToN(5)
	//h.Action.countToN(5)
}
