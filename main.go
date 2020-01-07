package main

import (
	"fmt"
  "runtime"
)

 

func main() { 
	var mm  map[int]string
  if mm == nil {
    fmt.Println("mm is nil")
  }
  fmt.Println(mm)
  fmt.Println(runtime.NumCPU())
}
