package main

import (
  "os"
  "testing"
  "strings"
  "net/http"
  "github.com/labstack/echo"
  . "github.com/franela/goblin"
)

func TestHandler(t *testing.T) {
  os.Setenv("ENV", "testing")
  g := Goblin(t)
  heartbeatResp := `{"message":"Success"}`

  g.Describe("Home page heartbeat", func() {

    e := echo.New()
    Init(e)
    c, b, ct := request("GET", "/", e, nil)

    g.It("should receive a 200 status", func() {
      g.Assert(c).Equal(http.StatusOK)
    })

    g.It("should receive a success JSON", func() {
      g.Assert(b).Equal(heartbeatResp)
      g.Assert(ct).Equal("application/json; charset=UTF-8")
    })
  })

  g.Describe("Add URL", func() {
    e := echo.New()
    Init(e)

    g.It("should return Bad Request with no data", func() {
      c, b, ct := request("POST", "/new", e, nil)
      g.Assert(c).Equal(http.StatusBadRequest)
      g.Assert(b).Equal(`{"message":"Bad Request"}`)
      g.Assert(ct).Equal("application/json; charset=UTF-8")
    })

    g.It("should receive a 201 created and response", func() {
      fileJSON := `{"url":"http://ryanbensonmedia.com"}`
      c, b, ct := request("POST", "/new", e, strings.NewReader(fileJSON))
      g.Assert(c).Equal(http.StatusCreated)
      g.Assert(b).Equal(`{"message":"Success"}`)
      g.Assert(ct).Equal("application/json; charset=UTF-8")
    })
  })
}
