package main

import (
	"context"
	"fmt"
	"time"
)

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Canceled")
			return
		}
	}
}

func mainFake() {
	ctx, cancel := context.WithCancel(context.Background())
	go tickOperation(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

}
