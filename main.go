package main

import "fmt"

type Node struct {
  data string
  next *Node
}

func (v Node) printData() string {
  return v.data;
}

func main() {
  node := Node{"bruno", nil}
  fmt.Println(node.printData())
}