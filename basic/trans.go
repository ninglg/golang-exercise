package main

import "fmt"

func transFunc() {
	var sum = 17
	var count = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean的值为: %f\n", mean)
}
