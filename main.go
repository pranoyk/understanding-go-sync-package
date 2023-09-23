package main

import (
	"fmt"
	"sync"
)

func helloWorld() {
	fmt.Println("Hello World!")
}

func main() {
	once := sync.OnceFunc(helloWorld)
	for i := 0; i < 5; i++ {
		once()
	}
}