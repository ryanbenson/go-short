package main

import(
  "github.com/labstack/echo"
)

func main() {
  e := echo.New()
  Init(e)
  Start(e)
}
