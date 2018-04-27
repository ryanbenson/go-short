package main

import (
  "os"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

// Init just sets up our server and starts up basic things
func Init(e *echo.Echo) {
  e.Use(middleware.Logger())
  SetupRoutes(e)
  Connect()
}

// Start runs the HTTP server
func Start(e *echo.Echo) {
  e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
