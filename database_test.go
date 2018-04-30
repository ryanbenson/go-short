package main

import (
  "os"
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestDatabase(t *testing.T) {
  os.Setenv("ENV", "testing")
  g := Goblin(t)

  //special hook for gomega
  RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

  g.Describe("Connects to the database", func() {

    g.It("should create a valid connection to Redis", func() {
      success, err := Connect()
      Ω(success).Should(Equal("PONG"))
      Ω(err).ShouldNot(HaveOccurred())
    })
  })

  g.Describe("Helpers", func() {
    g.It("should create a random string for our keys", func() {
      str := GetRandKey()
      Ω(str).Should(HaveLen(16))
    })

    g.It("should fail to find a key/val", func() {
      found, value := Find("fd7a89f7das98")
      g.Assert(found).Equal(false)
      g.Assert(value).Equal("")
    })
  })
}
