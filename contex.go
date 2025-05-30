package main

import (
	"context"
	"fmt"
	"time"
)

type key string

func main() {
	start := time.Now()

	// 1. Base Context
	ctx := context.Background()

	// 2. Context with Value
	ctx = context.WithValue(ctx, key("userID"), 10)

	// 3. Context with Deadline (set 500ms from now)
	deadline := time.Now().Add(500 * time.Millisecond)
	ctx, cancelDeadline := context.WithDeadline(ctx, deadline)
	defer cancelDeadline()

	// Call function that uses WithTimeout and WithCancel
	val, err := fetchUserData(ctx)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", val)
	}

	// 4. Context TODO (Example of placeholder context)
	placeholderContext := context.TODO()
	fmt.Println("TODO context is still valid?", placeholderContext.Err() == nil)

	fmt.Println("Time took:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(parentCtx context.Context) (int, error) {
	// 5. WithTimeout: adds 200ms timeout on top of parent context
	ctx, cancelTimeout := context.WithTimeout(parentCtx, 200*time.Millisecond)
	defer cancelTimeout()

	// 6. WithCancel: allows manual cancellation (optional in this example)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	respCh := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respCh <- Response{value: val, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("context canceled or timed out: %v", ctx.Err())
		case resp := <-respCh:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(150 * time.Millisecond)
	return 666, nil
}
