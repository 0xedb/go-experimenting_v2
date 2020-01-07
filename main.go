package main

import (
	"fmt"
	"sync"
  "time"
)

func writeSomething(msg string, num int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < num; i++ {
			fmt.Println(msg)
      time.Sleep(100 * time.Microsecond)
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
