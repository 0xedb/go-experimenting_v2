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

func main() {
  bruno := &Human{"kofi", 39}
  // bruno_ptr := &bruno 
  fmt.Println(*bruno)
}

