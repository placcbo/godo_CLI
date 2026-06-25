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

func (t1 *TodoList) Add(title string) Todo {
	todo := Todo{
		Title: title,
		ID:    t1.nextID,
	}
	t1.items = append(t1.items, todo)
	t1.nextID++
	return todo
}

func (t1 *TodoList) MarkDone(id int) error {
	for i := range t1.items {
		if t1.items[i].ID == id {
			t1.items[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("todo %d not found", id)
}

func (t1 *TodoList) Remove(id int) error {
	for i := range t1.items {
		if t1.items[i].ID == id {
			t1.items = append(t1.items[:i], t1.items[i+1:]...)
				return nil
		}
	
	}

	return fmt.Errorf("todo %d not found", id)
}

func main() {
	todos := NewTodoList()
	todos.Add("learn go")
	todos.Add("Build API")
	todos.Add("deploy App")
	fmt.Println("before:")
	fmt.Println(todos.items)
	todos.Remove(2)
	fmt.Println("after:")
	fmt.Println(todos.items)
}
