package main

import (
	"context"
	"fmt"
	"time"
)

func mainFake() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
	}()

	select {
	case <-done:
		fmt.Println("Done task")
	case <-ctxWithTimeout.Done():
		fmt.Println("Timeout")
	}
}
