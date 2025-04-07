package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	exampleTimeout()
}

func exampleTimeout() {
	//creates an empty context
	ctx := context.Background()

	//Returns a context and a cancel function
	//If the timeout is set to 2 seconds the ctxWithTimeout will be done
	//If it is set to 4 seconds it will the done channel will be closed and will print calle the api
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	//make it defer the cancel function
	defer cancel()

	//make a channel to handle the contexts and the cancel function
	done := make(chan struct{})

	//create a goroutine
	go func() {
		//make it wait 3 seconds
		time.Sleep(3 * time.Second)
		//make it close the channel
		close(done)
	}()

	//Select is a keyword for async things like channels, contexts
	select {
	case <-done:
		fmt.Println("Called the api")
	case <-ctxWithTimeout.Done():
		fmt.Println("oh no my timeout expired!", ctxWithTimeout.Err())
	}
}
