package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		fmt.Println(hello())
		time.Sleep(time.Second)
	}
}

func hello() string {
	return "Hello, World!"
}
