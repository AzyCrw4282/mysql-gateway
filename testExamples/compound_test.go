package main

import (
	"fmt"
	"time"
)

var svr = "http://localhost:8080"

func main() {
	start := time.Now()

	fmt.Println(time.Since(start))
}
