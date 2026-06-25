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

func main() {
	list := TodoList{items: []Todo{}}
	list.items = append(list.items, Todo{1, "learn go", true})
	list.items = append(list.items, Todo{2, "buy milk", false})
	fmt.Println(list.items[0].Title)

}
