package main

import (
  "math/rand"
  "github.com/shomali11/xredis"
)

var client *xredis.Client
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const keyLen = 16

type Redirection struct {
  Url  string `json:"url" form:"url" query:"url"`
}

// Connect starts a connection to Redis and binds the client to the global var
// it also runs a Ping to make sure it works
func Connect() (string, error) {
  client = xredis.DefaultClient()
  return client.Ping()
}

// Save a new Url
func Save(url string) (string, error) {
  key := GetRandKey()
  _, err := client.SetNx(key, url)
  return key, err
}

// RandStringBytes gives us a random string
func GetRandKey() string {
    b := make([]byte, keyLen)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}
