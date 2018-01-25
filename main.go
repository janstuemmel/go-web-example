package main

import (
  "net/http"

  . "github.com/janstuemmel/go-web-example/controller"
  . "github.com/janstuemmel/go-web-example/provider"
)


func main() {


  provider := NewProvider()
  controller := NewController(provider);

  // handle root
  http.HandleFunc("/", controller.RootRoute);

  // handle api
  http.HandleFunc("/test.json", controller.JsonRoute)

  // start server
  http.ListenAndServe(":8000", nil)
}
