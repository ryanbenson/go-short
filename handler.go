package main

import (
  "net/http"
  "github.com/labstack/echo"
)

// Response is a basic string to respond back to the request
type Response struct {
  Message string `json:"message"`
}

// Heartbeat is a ping/pong check to make sure the API is healthy
// it just returns a basic message
func Heartbeat(c echo.Context) error {
  success := &Response{Message: "Success"}
  return c.JSON(http.StatusOK, success)
}

// CreateRedirect saves a new redirect in the database
func CreateRedirect(c echo.Context) error {
  r := new(Redirection)
  if err := c.Bind(r); err != nil {
    badRequestMessage := &Response{Message: "Bad Request"}
    return c.JSON(http.StatusBadRequest, badRequestMessage)
  }
  urlKey, err := Save(r.URL)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, &Response{Message: "Internal Error"})
  }
  success := &Response{Message: urlKey}
  return c.JSON(http.StatusCreated, success)
}

// CheckAndRedirect validates the request, and redirects or throws 404 at you
func CheckAndRedirect(c echo.Context) error {
  redir := c.Param("redir")
  found, destURL := Find(redir)
  if found == false || destUrl == "" {
    notFoundMessage := &Response{Message: "Page Not Found"}
    return c.JSON(http.StatusNotFound, notFoundMessage)
  }
  return c.Redirect(http.StatusMovedPermanently, destURL)
}
