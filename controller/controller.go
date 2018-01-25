package controller

import (
  "io"
  "net/http"
)

type DataProvider interface {
  Get() ([]byte, error)
}


func RootRoute(provider DataProvider) func(w http.ResponseWriter, r *http.Request) {

  return func(w http.ResponseWriter, r *http.Request) {

    io.WriteString(w, "Welcome home!")
  }
}


func JsonRoute(provider DataProvider) func(w http.ResponseWriter, r *http.Request) {

  return func(w http.ResponseWriter, r *http.Request) {

      b, _ := provider.Get()

      w.WriteHeader(http.StatusOK)
      w.Header().Set("Content-Type", "application/json")
      w.Write(b)
  }
}
