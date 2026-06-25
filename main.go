package main

import "fmt"

type Todo struct {
	ID    int
	Title string
	Done  bool
}

type TodoList struct {
	items  []Todo
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{items: []Todo{}, nextID: 1}
}

func (tl *TodoList) Add(title string) Todo {
	t := Todo{ID: tl.nextID, Title: title}
	tl.items = append(tl.items, t)
	tl.nextID++
	return t
}

func main() {
	list := NewTodoList()

	t := list.Add("Learn Go")

	fmt.Println("added:", t)
	fmt.Println("all items:", list.items)
	fmt.Println("next ID:", list.nextID)
}
