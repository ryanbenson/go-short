package main

import (
  "github.com/labstack/echo"
)

func SetupRoutes(e *echo.Echo) {
  e.GET("/", Heartbeat)
}
