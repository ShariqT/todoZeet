package main

import (
	"time"
	"github.com/google/uuid"
)

type Todo struct {
	ID uuid.UUID	 `json:"id"`
	Title string     `json:"title"`
	Finished bool	 `json:"finished"`
	Posted time.Time `json:"posted"` 
}

func NewTodo(title string, isFinished bool) *Todo {
	todo := Todo{Title: title, Finished: isFinished}
	todo.Posted = time.Now()
	todo.ID = uuid.New()
	return &todo
}