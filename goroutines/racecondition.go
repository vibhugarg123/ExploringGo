package main

import (
	"fmt"
	"sync"
)

/*
To check if race condition exists-
  - use go run racecondition.go --race

Use mutex when you are trying to access critical section
*/
func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	score := []int{0}

	// 3 denotes the number of goroutines
	wg.Add(3)

	func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One R")

		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two R")

		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three R")

		mut.Lock()
		score = append(score, 3)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
