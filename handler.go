package main

import (
  "net/http"
  "github.com/labstack/echo"
)

type Response struct {
  Message string `json:"message"`
}

func Heartbeat(c echo.Context) error {
  success := &Response{Message: "Success"}
  return c.JSON(http.StatusOK, success)
}
