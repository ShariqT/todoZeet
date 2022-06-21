package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/siredwin/pongorenderer/renderer"
)

func main() {

	MainRenderer := renderer.Renderer{Debug: true}

	var db []*Todo
	db = append(db, NewTodo("dfsdf", false))
	db = append(db, NewTodo("dsdfssdfsdfs", false))

	routeHandlers := &Handler{DB: db}
	// routeHandlers := &Handler{}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = MainRenderer
	e.Static("/static", "assets")
	e.GET("/", routeHandlers.listTodos)
	// e.POST("/update", routeHandlers.updateTodo)
	e.Logger.Info(e.Start(":3000"))
}
