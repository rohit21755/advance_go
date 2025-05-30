package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userId := 10
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result: ", val)
	fmt.Println("time took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}

	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("third fetching data tooks to long")
		case resp := <-respch:
			return resp.value, resp.err
		}

	}

	// if err != nil {
	// 	return 0, err
	// }
	// return val, nil
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 150)
	return 666, nil
}
