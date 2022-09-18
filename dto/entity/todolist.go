package entity

type TodoList struct {
	ID        string // random identity generated uniquely
	Title     string // tittle of todolist
	Desc      string // description of the todolist
	Done      bool   // status of the todolist
	Category  string // category of the todolist
	Priority  int    // 0=low, 1=med, 2=high
	DueDate   string // deadline of the todolist
	CreatedAt int64  // database creation date
	UpdatedAt int64  // database modification date
}
