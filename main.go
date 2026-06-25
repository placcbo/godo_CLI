package main

import "fmt"

type Todo struct {
	Title string
	ID    int
	Done  bool
}

type TodoList struct {
	items  []Todo
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{
		items:  []Todo{},
		nextID: 1,
	}
}

func (t1 *TodoList) Add(title string) {
	todo := Todo{
		ID:    t1.nextID,
		Title: title,
	}

	t1.items = append(t1.items, todo)
	fmt.Println(todo)
}

func main() {
	todos := NewTodoList()
	todos.Add("Learn Go")

	fmt.Println(todos.items)
}
