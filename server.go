package main

import (
  "os"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
  e.Use(middleware.Logger())
  SetupRoutes(e)
  Connect()
}

func Start(e *echo.Echo) {
  e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
