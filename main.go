package main

import (
	"fmt"
	"sync"
)

type BookKeeper struct {
	name        string
	id          string
	phoneNumber string
}

func NewBookKeeper(name, id, phoneNumber string) *BookKeeper {
	return &BookKeeper{
		name:        name,
		id:          id,
		phoneNumber: phoneNumber,
	}
}

func (bk *BookKeeper) helloWorld() {
	fmt.Println("Hello World!")
}

func (bk *BookKeeper) getName() string {
	return bk.name
}

func (bk *BookKeeper) getNameAndPhoneNumber() (string, string) {
	return bk.name, bk.phoneNumber
}

func main() {
	bk :=  NewBookKeeper("John", "123456789", "1234567890")

	// sync.OnceFunc
	once := sync.OnceFunc(bk.helloWorld)
	for i := 0; i < 5; i++ {
		once()
	}

	// sync.OnceValue
	onceValue := sync.OnceValue[string](bk.getName)
	for i := 0; i < 5; i++ {
		fmt.Println(onceValue())
	}

	// sync.OnceValues
	onceValues := sync.OnceValues[string, string](bk.getNameAndPhoneNumber)
	for i := 0; i < 5; i++ {
		fmt.Println(onceValues())
	}
}
