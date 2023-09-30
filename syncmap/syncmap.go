package syncmap

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	m sync.Map
}

type AnotherStruct struct {
	name        string
	id          string
	phoneNumber string
}

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

func (sm *SyncMap) storeAnotherStruct() {
	// store method is not type safe. We have to type convert when we load the value
	sm.m.Store("anotherStruct", AnotherStruct{
		name:        "John",
		id:          "123456789",
		phoneNumber: "1234567890",
	})
}

func (sm *SyncMap) loadAnotherStruct() AnotherStruct {
	value, ok := sm.m.Load("anotherStruct")
	if !ok {
		return AnotherStruct{}
	}
	// we have to type cast the value to AnotherStruct
	anotherStruct, ok := value.(AnotherStruct)
	if ok {
		return anotherStruct
	} else {
		return AnotherStruct{}
	}
}

func (sm *SyncMap) deleteAnotherStruct() {
	sm.m.Delete("anotherStruct")
}

func RunSyncMap() {
	sm := NewSyncMap()
	sm.storeAnotherStruct()
	var anotherStruct1 AnotherStruct
	var anotherStruct2 AnotherStruct
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		anotherStruct1 =  sm.loadAnotherStruct()
	}()
	go func() {
		defer wg.Done()
		anotherStruct2 =  sm.loadAnotherStruct()
	}()
	wg.Wait()
	fmt.Printf("value from first fetch, name: %s\n",anotherStruct1.name)
	fmt.Printf("value from second fetch, name: %s\n",anotherStruct2.name)
	sm.deleteAnotherStruct()
	anotherStruct := sm.loadAnotherStruct()
	fmt.Printf("value after delete, name: %s\n",anotherStruct.name)
}
