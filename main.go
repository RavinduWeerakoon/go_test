package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if rand.Intn(2) == 0 { // 50% chance to fail
		fmt.Println("Job failed")
		os.Exit(1)
	}
	time.Sleep(30 * time.Second)
}
