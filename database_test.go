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
}
