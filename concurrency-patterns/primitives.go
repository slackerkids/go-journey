package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func DoWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing some work")
		}
	}
}

// go concurrency model is Fork Join model. Where goroutine task after it's completed will be joined into main
// channels is FIFO queue like datastructure created to communicate between goroutines
// select is a switch like syntax that allows to wait multiple channels and depending what is done to show. (if multiple chanells will be ready returns randomly on fairness principle.)

func main() {
	// go someFunc("1")

	// myChan := make(chan string)
	// anotherChan := make(chan string)

	// go func() {
	// 	myChan <- "data"
	// }()

	// go func() {
	// 	// time.Sleep(time.Second * 1)
	// 	anotherChan <- "another data"
	// }()

	// select {
	// case msgFromMyChan := <-myChan:
	// 	fmt.Println(msgFromMyChan)
	// case msgFromAnotherChan := <-anotherChan:
	// 	fmt.Println(msgFromAnotherChan)
	// }

	// Blocking line of code, because main waits a channel to be filled
	// msg := <- myChan

	// charChan := make(chan string, 3)
	// chars := []string{"a", "b", "c"}

	// for _, s := range chars {
	// 	select {
	// 	case charChan <- s:
	// 	}
	// }

	done := make(chan bool)

	go DoWork(done)

	time.Sleep(time.Second * 3)

	close(done)
}
