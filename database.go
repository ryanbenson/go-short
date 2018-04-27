package main

import (
  "github.com/shomali11/xredis"
)

var client *xredis.Client

// Connect starts a connection to Redis and binds the client to the global var
// it also runs a Ping to make sure it works
func Connect() (string, error) {
  client = xredis.DefaultClient()
  defer client.Close()
  return client.Ping()
}
