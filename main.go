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

func main() {
  singl := Stringly("bruno edoh")
  singl.printData()
}