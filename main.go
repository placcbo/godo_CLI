package main

type Todo struct {
	Title string
	ID    int
	Done  bool
}

type TodoList struct {
	items  []Todo
	NextId int
}

func NewTodoList() *TodoList {
	return &TodoList{items: []Todo{}, NextId: 1}
}

func (t1 *TodoList) Add(title string) Todo {
	t := Todo{ID: t1.NextId, Title: title}
	t1.items = append(t1.items, t)
	t1.NextId++
	return t
}
