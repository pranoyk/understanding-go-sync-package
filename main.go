package main

import (
	"github.com/pranoyk/understanding-go-sync-package/syncmap"
	"github.com/pranoyk/understanding-go-sync-package/synconce"
)

func main() {
	synconce.RunBookKeeper()
	syncmap.RunSyncMap()
}
