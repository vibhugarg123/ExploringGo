package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
Mutex: Mutual exclusion lock. Zero value is unlocked mutex.
  - Locks memory till particular go routine is working. Till it's not finished writing to a shared resource,
    it'll not allow any other go routine to write to it.
*/
var wgroup sync.WaitGroup
var mut sync.Mutex // usually a pointer
var signals = []string{"test"}

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCodeUtil(web)
		wgroup.Add(1)
	}

	//usually goes to the end of method/function, please don't exit.
	wgroup.Wait()
	fmt.Println(signals)
}
func getStatusCodeUtil(endPoint string) {
	// Passes a signal when the function is completed.
	defer wgroup.Done()

	res, err := http.Get(endPoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}

	mut.Lock()
	signals = append(signals, endPoint)
	mut.Unlock()
	
	fmt.Printf("%d status code for %s\n", res.StatusCode, endPoint)
}
