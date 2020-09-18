package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}

func mySleepAndTalk(ctx context.Context, dur time.Duration, msg string) {
	select {
	case <-time.After(dur):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("done")
	}

}
