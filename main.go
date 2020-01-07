package main

import (
	"fmt"
	"sync"
)

func writeSomething(msg string, num int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < num; i++ {
			fmt.Println(msg)
		}
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
  writeSomething("bruno", 10, &wg)
  writeSomething("kofi", 20, &wg)
  writeSomething("edoh", 14, &wg)
	wg.Wait()
}
