package main

import "fmt"

type Node struct {
  data string
  next *Node
}

type Stringly string

func (v Stringly) printData()  {
  fmt.Println(">>>>>>>>>>", v)
}

type Human struct {
  name string
  age uint16
}

func (human *Human) setInfo(name string, age uint16) {
  human.name = name
  human.age = age
}

type walkable interface {
  kind() string
  distance() float64
}

func (vv Human) kind() string {
  return "I am kinda kind"
}

func (vv Human) distance() float64 {
  return 14.23
}

func main() {
  bruno := &Human{"kofi", 39} 
  fmt.Println(bruno.distance())
  fmt.Println(bruno.kind())
}

