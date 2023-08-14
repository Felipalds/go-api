package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.StampMicro))
	rand.Seed(now.UnixNano())
	//if we don't pass the seed, go will generate the same pseudorandom numbers
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(10))
	}
}
