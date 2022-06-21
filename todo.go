package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string     `json:"title" bson:"title"`
	Finished bool	 `json:"finished" bson:"finished"`
	Posted time.Time `json:"posted" bson:"posted"` 
}

func NewTodo(title string, isFinished bool) *Todo {
	todo := Todo{Title: title, Finished: isFinished}
	todo.Posted = time.Now()
	return &todo
}