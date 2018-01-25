package controller

import (
  "net/http"
)

type DataProvider interface {
  Get() ([]byte, error)
}

type Controller struct {
  provider DataProvider
}

func NewController(provider DataProvider) Controller {
  return Controller{provider}
}

func (c Controller) JsonRoute(w http.ResponseWriter, r *http.Request) {

  b, _ := c.provider.Get()

  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  w.Write(b)
}

func (c Controller) RootRoute(w http.ResponseWriter, r *http.Request) {

  w.Write([]byte("Welcome home!"))
}
