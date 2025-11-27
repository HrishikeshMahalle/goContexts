package core

import (
	"context"
	"fmt"
	"time"
)

func Start(ctx context.Context) {
	fmt.Println("Starting the application...")

	ch := make(chan bool)

	go func(ctx context.Context) {
		i := 0
		for {
			fmt.Println("Working...", i)
			time.Sleep(1 * time.Second)
			if i == 5 {
				ch <- false
				break
			}
			select {
			case <-ctx.Done():
				ch <- false
				return
			default:
				// fmt.Println("default")
			}
			i++
		}
	}(ctx)
	<-ch
	fmt.Println("Done")
}

func worker(ctx context.Context, i int) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", i, "cancelled")
			return
		default:
			fmt.Println("Worker", i, "working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func StartTwo(ctx context.Context, cancel context.CancelFunc) {

	for i := 0; i < 5; i++ {
		go worker(ctx, i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("cancelling context now")
	cancel()
	time.Sleep(1 * time.Second)
}
