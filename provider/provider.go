package provider

import . "github.com/janstuemmel/go-web-example/model"

type Provider struct {}

func NewProvider() *Provider {
  return &Provider{}
}

func (p *Provider) Get() Message {

  return NewMessage("Bob", "Hello World")
}
