package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func Count(num *int) {
	for range 100000 {
		mu.Lock()
		*num++
		mu.Unlock()
	}
}

func Subtract(num *int) {
	for range 100000 {
		mu.Lock()
		*num--
		mu.Unlock()
	}
}
func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(6)

	var m sync.Map

	m.Store("hello", 1)

	go func() {
		defer wg.Done()
		Count(&counter)
	}()

	go func() {
		defer wg.Done()
		Subtract(&counter)
	}()

	go func() {
		defer wg.Done()
		Count(&counter)
	}()

	go func() {
		defer wg.Done()
		Subtract(&counter)
	}()

	go func() {
		defer wg.Done()
		Count(&counter)
	}()

	go func() {
		defer wg.Done()
		Subtract(&counter)
	}()

	wg.Wait()
	fmt.Println(counter)
}
