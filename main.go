package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "github.com/siredwin/pongorenderer/renderer"
)

func main() {
	var db []*Todo
	db = append(db, NewTodo("dfsdf", false))
	db = append(db, NewTodo("dsdfssdfsdfs", false))

	routeHandlers := &Handler{DB: db}
	e := echo.New()
	// e.Use(middleware.Logger())
	e.GET("/", routeHandlers.listTodos)
	// e.POST("/update", routeHandlers.updateTodo)
	e.Start(":3000")
}
