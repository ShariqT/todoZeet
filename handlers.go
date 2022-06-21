package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/labstack/echo"
	"github.com/flosch/pongo2"
	"strconv"
	"fmt"
)

type Handler struct {
	DB *mgo.Session
}

func (h *Handler) listTodos (c echo.Context) error {
	db := h.DB.Clone()
	defer db.Close()
	var todos []*Todo
	if err := db.DB("todolist").C("todos").Find(bson.M{}).All(&todos); err != nil {
		data := pongo2.Context{}
		data["error"] = err
		return c.Render(500, "templates/500.html", data)
	}
	data := pongo2.Context{}
	data["todos"] = todos
	return c.Render(200, "templates/index.html", data)
}

func (h *Handler) updateTodo (c echo.Context) error {
	db := h.DB.Clone()
	defer db.Close()
	todoID := bson.ObjectIdHex(c.FormValue("ID"))
	newFinishedValue, _ := strconv.ParseBool(c.FormValue("Finished"))
	var foundTodo *Todo
	if err := db.DB("todolist").C("todos").Find(bson.M{"_id":todoID}).One(&foundTodo); err != nil {
		data := pongo2.Context{}
		data["error"] = err.Error()
		fmt.Println(err)
		return c.Render(500, "templates/500.html", data)
	}
	foundTodo.Finished = newFinishedValue
	if err := db.DB("todolist").C("todos").UpdateId(todoID, foundTodo); err != nil {
		data := pongo2.Context{}
		data["error"] = err
		return c.Render(500, "templates/500.html", data)
	}
	data := pongo2.Context{}
	data["todo"] = foundTodo
	data["action"] = "Updated task status to " + c.FormValue("Finished")
	return c.Render(200, "templates/status.html", data)
	
}