package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("hello")
		waitGrp.Done()

	}()

	go func() {
		fmt.Println("world")
		waitGrp.Done()

	}()

	waitGrp.Wait()
}
