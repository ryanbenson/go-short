package main

import (
	"os"
	"testing"
  "net/http"
	"github.com/labstack/echo"
  . "github.com/franela/goblin"
)

var heartbeatResp string = `{"message":"Success"}`

func TestHandler(t *testing.T) {
	os.Setenv("ENV", "testing")
  g := Goblin(t)

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
}
