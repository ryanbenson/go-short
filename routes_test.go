package main

import (
  "os"
  "net/http"
  "net/http/httptest"
  "testing"
  "io"
  "github.com/labstack/echo"
  . "github.com/franela/goblin"
)

func request(method, path string, e *echo.Echo, data io.Reader) (int, string, string) {
  req := httptest.NewRequest(method, path, data)
  req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
  rec := httptest.NewRecorder()
  e.ServeHTTP(rec, req)
  return rec.Code, rec.Body.String(), rec.HeaderMap["Content-Type"][0]
}

func TestRoute(t *testing.T) {
  os.Setenv("ENV", "testing")
  g := Goblin(t)

  g.Describe("Routes", func() {
    e := echo.New()
    Init(e)

    g.It("should respond to the heartbeat", func() {
      c, _, ct := request("GET", "/", e, nil)
      g.Assert(c).Equal(http.StatusOK)
      g.Assert(ct).Equal("application/json; charset=UTF-8")
    })
  })
}
