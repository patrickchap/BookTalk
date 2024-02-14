package main

import (
	"github/patrickchap/booktalk/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()
	router.Static("/static", "static")

	homeHandler := handler.HomeHandler{}
	browseHandler := handler.BrowseHandler{}

	router.GET("/Browse", browseHandler.Index)
	router.GET("/", homeHandler.Index)

	router.Start("0.0.0.0:8080")
}
