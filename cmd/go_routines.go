package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 300)
	}
}

func handlePanic() {
	// defer helps us get here, without this, the execution will hang if there is a problem
	if r := recover(); r != nil {
		fmt.Println("PANIC!")
	}
}

func printStuff() {
	defer wg.Done() // have to call this to ensure wg knows when its done, pop it off the stack
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

// If there is nothing to be executed first, it won't execute at all
func main() {
	wg.Add(1)
	go printStuff()
	wg.Wait() // wait until this is done
	// TODO learn about channels
	//go say("Hello")
	//say("There")
}
