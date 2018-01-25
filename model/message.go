package model

import "encoding/json"

type Message struct {
  Name string
  Body string
}

func NewMessage(name string, body string) Message {
  return Message{name, body}
}

func (m Message) ToJSON() ([]byte, error) {
  return json.Marshal(m)
}
