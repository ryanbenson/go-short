package main

import (
  "os"
  "testing"
  "strings"
  "net/http"
  "github.com/labstack/echo"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestHandler(t *testing.T) {
  os.Setenv("ENV", "testing")
  g := Goblin(t)
  heartbeatResp := `{"message":"Success"}`
  var sampleRedirectStr string


  //special hook for gomega
  RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

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
      g.Assert(ct).Equal("application/json; charset=UTF-8")
      sampleRedirectStr = b[12:24]
      Ω(sampleRedirectStr).Should(HaveLen(12))
    })
  })

  g.Describe("Checks and redirects", func() {
    e := echo.New()
    Init(e)

    g.It("should give a 404 for a bad redirect string", func() {
      c, _, _ := request("GET", "/will-not-work", e, nil)
      g.Assert(c).Equal(http.StatusNotFound)
    })

    // TODO: need to figure out how to test this without it breaking
    // g.It("should redirect with a valid redirect string", func() {
    //   // urlToRedirect := "/" + sampleRedirectStr
    //   c, _, _ := request("GET", urlToRedirect, e, nil)
    //   g.Assert(c).Equal(http.StatusMovedPermanently)
    // })
  })
}
