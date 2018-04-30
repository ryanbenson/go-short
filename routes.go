package main

import (
  "github.com/labstack/echo"
)

// SetupRoutes sets up the list of routes for the server
// all handlers are setup elsewhere
func SetupRoutes(e *echo.Echo) {
  e.GET("/", Heartbeat)
  e.POST("/new", CreateRedirect)
  e.GET("/:redirect", CheckAndRedirect)
}
