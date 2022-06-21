package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/siredwin/pongorenderer/renderer"
	"gopkg.in/mgo.v2"
	"os"
)

func main() {

	MainRenderer := renderer.Renderer{Debug: true}

	db, err := mgo.Dial(os.Getenv("DATABASE"))
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
