package main

import (
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

// TodoList wraps the slice so we can hang methods off it.
type TodoList struct {
	items  []Todo
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{items: []Todo{}, nextID: 1}
}

// Add appends a new Todo and returns it. Pointer receiver because
// we're mutating the list's underlying slice and nextID counter.
func (tl *TodoList) Add(title string) Todo {
	t := Todo{ID: tl.nextID, Title: title}
	tl.items = append(tl.items, t)
	tl.nextID++
	return t
}

func (tl *TodoList) Done(id int) error {
	for i := range tl.items {
		if tl.items[i].ID == id {
			tl.items[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("todo %d not found", id)
}

func (tl *TodoList) Remove(id int) error {
	for i, t := range tl.items {
		if t.ID == id {
			// slice trick: cut the element by re-slicing around it
			tl.items = append(tl.items[:i], tl.items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo %d not found", id)
}

func main() {
	list := NewTodoList()

	if len(os.Args) < 2 {
		fmt.Println("usage: godo <add|list|done|remove> [args]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		t := list.Add(os.Args[2])
		fmt.Printf("added #%d — %q\n", t.ID, t.Title)
	case "done":
		id, _ := strconv.Atoi(os.Args[2])
		if err := list.Done(id); err != nil {
			fmt.Println("error:", err)
		}
	case "list":
		for _, t := range list.items {
			mark := " "
			if t.Done {
				mark = "x"
			}
			fmt.Printf("[%s] #%d  %s\n", mark, t.ID, t.Title)
		}
	default:
		fmt.Println("unknown command:", os.Args[1])
	}
}
