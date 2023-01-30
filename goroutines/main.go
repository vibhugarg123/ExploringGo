package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
		Concurrency vs Parallelism

	 1. Concurrency is the task of running and managing the multiple computations at the same time.
	    Parallelism is the task of running multiple computations simultaneously.

	 2. Concurrency is achieved through the interleaving operation of processes on the central processing unit(CPU)
	    or in other words by the context switching.
	    Parallelism is achieved by through multiple central processing units(CPUs).

	 3. Concurrency can be done by using a single processing unit.
	    Parallelism can’t be done by using a single processing unit. it needs multiple processing units.

	 4. Concurrency increases the amount of work finished at a time.
	    Parallelism improves the throughput and computational speed of the system.

	 5. Concurrency deals a lot of things simultaneously.
	    Parallelism do a lot of things simultaneously.

	 6. Concurrency is the non-deterministic control flow approach.
	    Parallelism is deterministic control flow approach.

	 7. In concurrency debugging is very hard.
	    In parallelism, debugging is also hard but simpler than concurrency.

	    Thread vs GoRoutines

	 1. Threads are managed by OS with fixed stack 1 MB.
	    GoRoutines are managed by Go runtime with flexible stack 2 KB

	    DO NOT COMMUNICATE BY SHARING MEMORY. INSTEAD, SHARE MEMORY BY COMMUNICATING.
*/

/*
	     Basic way to create a go routine

			func main() {
				go greeter("hello")
				greeter("world")
			}

			func greeter(s string) {
				for i := 0; i < 6; i++ {
					// basic way to stop for a thread
					time.Sleep(3 * time.Millisecond)
					fmt.Println(s)
				}
			}
*/

// Advanced version of time.Sleep(), it's usually a pointer.
// - ADD()  : If 5 jobs are added, it will wait for 5 of them to finish.
// - DONE() : When your job is done
// - WAIT() : When you wait for your job to be completed.
var wg sync.WaitGroup

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	//usually goes to the end of method/function, please don't exit.
	wg.Wait()
}
func getStatusCode(endPoint string) {
	// Passes a signal when the function is completed.
	defer wg.Done()

	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endPoint)
}
