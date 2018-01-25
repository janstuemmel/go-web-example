package main

import (
  "net/http"
  "encoding/json"

  . "github.com/janstuemmel/go-web-example/controller"
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
  controller := NewController(provider);

  // handle root
  http.HandleFunc("/", controller.RootRoute);

  // handle api
  http.HandleFunc("/test.json", controller.JsonRoute)

  // start server
  http.ListenAndServe(":8000", nil)
}
