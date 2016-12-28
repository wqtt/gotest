package main

import (
	"fmt"
	"time"
)

func errorFunc() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%d\n", i)
		}()
	}
	time.Sleep(time.Second)
}

func correctFunc() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("%d\n", i)
		}(i)
	}
	time.Sleep(time.Second)
}

func main() {
	errorFunc()
	correctFunc()

	time.Sleep(time.Second)
}
