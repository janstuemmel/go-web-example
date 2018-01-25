package controller

import (
  "net/http"

  . "github.com/janstuemmel/go-web-example/model"
)

type DataProvider interface {
  Get() Message
}

type Controller struct {
  provider DataProvider
}

func NewController(provider DataProvider) Controller {
  return Controller{provider}
}

func (c Controller) JsonRoute(w http.ResponseWriter, r *http.Request) {

  m := c.provider.Get()
  b, _ := m.ToJSON()

  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  w.Write(b)
}

func (c Controller) RootRoute(w http.ResponseWriter, r *http.Request) {

  w.Write([]byte("Welcome home!"))
}
