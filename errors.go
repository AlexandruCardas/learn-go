package main

import (
	"errors"
	"fmt"
	"os"
)

func isGreaterThanTen(num int) error {
	if num < 10 {
		return errors.New("something bad happened")
	}
	return nil
}

func openFile() error {
	f, err := os.Open("missingFile.txt") // check with go doc os.Open

	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

func doThings() {
	defer fmt.Println("First line but do this last!")
	defer fmt.Println("Do this second to last!")
	fmt.Println("Things and stuff should happen first")
}

// Recover will only hold a value if there is a panic
func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked but everyone's fine")
		fmt.Println(r)
	}
}

func doMoreThings() {
	defer recoverFromPanic()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("Panic!")
		}
	}
}

func main() {
	var num = 9
	doThings()
	// more idiomatic, the error is in the scope of the if block only
	doMoreThings()
	if err := isGreaterThanTen(num); err != nil {
		fmt.Println(fmt.Errorf("%d is not greater than ten", num))
		// or
		/*panic(err) // stop execution of the whole program*/
		// or
		//log.Fatal(err) // send this to a log
	}

	if err := openFile(); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}
