package main

import (
	"github.com/shomali11/xredis"
)

var client *xredis.Client

func Connect() (string, error) {
	client = xredis.DefaultClient()
	defer client.Close()
  return client.Ping()
}
