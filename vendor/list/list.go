package list

type TodoList struct {
	Name  string
	Items []Task
}

type Task struct {
	Name        string
	Description string
	Done        bool
}

var CurrentList TodoList
