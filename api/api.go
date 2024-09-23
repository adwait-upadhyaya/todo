package api

type TodoItemResponse struct {
	Id        int
	Task      string
	Completed bool
}

type TodoItemParams struct {
	Id int
}

type Error struct {
	Code    int
	Message string
}
