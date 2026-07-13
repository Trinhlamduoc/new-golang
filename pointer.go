package main

import (
	"fmt"
	// "runtime"
	// "math"
)

func main() {

	fmt.Printf("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")

}
