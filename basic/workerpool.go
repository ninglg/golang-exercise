package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int) {
	fmt.Println(i)
}

func workerpool() {
	fmt.Println("========== workerpool ==========")
	start := time.Now()
	channel := make(chan int, 10)
	var wg sync.WaitGroup
	var workerPoolSize = 10
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for a := range channel {
					process(a)
				}
			}()
		}
	}()

	for i := 0; i < 30; i++ {
		channel <- i
	}
	close(channel)
	wg.Wait()

	elasped := time.Since(start)

	fmt.Printf("Took %s\n", elasped)
}
