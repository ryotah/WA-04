package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/ryotah/WA-04/numutils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide numbers")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	messages := make(chan string)

	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println(arg, err)
			os.Exit(1)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			messages <- fmt.Sprintf("%d %s", num, numutils.FizzBuzz(num))
		}()
	}

	go func() {
		// Move wg.Wait() into a goroutine to prevent the error
		// "fatal error: all goroutines are asleep - deadlock!"
		wg.Wait()
		close(messages)
	}()

	for msg := range messages {
		fmt.Println(msg)
	}
}
