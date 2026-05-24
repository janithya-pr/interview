package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	// named function goroutine
	go Task("A", &wg)
	go Task("B", &wg)

	// anonymous function goroutine
	for i := 1; i <= 3; i++ {
		go func(num int) {
			fmt.Println("C", num)
		}(i)
	}
	fmt.Println("Main function")

	wg.Wait()
}

func Task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(2 * time.Second)
	}
}