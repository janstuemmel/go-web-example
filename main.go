package main

import (
  "net/http"
  "encoding/json"

  "github.com/janstuemmel/go-web-example/controller"
)

type Message struct {
  Name string
  Body string
  Time int64
}

type TestJsonProvider struct {}

func (t *TestJsonProvider) Get() ([]byte, error) {

  m := Message{ "Bob", "Hello World", 123 }

  b, err := json.Marshal(m)

  return b, err
}


func main() {

  provider := &TestJsonProvider{}

  // handle root
  http.HandleFunc("/", controller.RootRoute(provider));

  // handle api
  http.HandleFunc("/test.json", controller.JsonRoute(provider))

  // start server
  http.ListenAndServe(":8000", nil)
}
