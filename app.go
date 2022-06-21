package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"gopkg.in/mgo.v2"
	"github.com/siredwin/pongorenderer/renderer"
	// "github.com/flosch/pongo2"
)


func main() {

	MainRenderer := renderer.Renderer{Debug:true}

	db, err := mgo.Dial("localhost")
	if err != nil {
		panic("Could not establish connection the database")
	}
	defer db.Close()

	routeHandlers := &Handler{DB: db}
	// routeHandlers := &Handler{}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = MainRenderer
	e.Static("/static", "assets")
	e.GET("/", routeHandlers.listTodos)
	e.POST("/update", routeHandlers.updateTodo)
	e.Logger.Info(e.Start(":3000"))
}