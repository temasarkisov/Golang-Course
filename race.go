// A data race happens when there are two memory accesses in a program
// where both target the same location, performed conc. by the same thread,
// are not reads and non-synchronous.
// The following two function calls can result in a data race condition
// when incrementing the count variable of the program.
/*
package main

import (
	"sync"
	"time"
)

var count int
var mut sync.Mutex

func IncCount() {
	mut.Lock
	count++
	mut.Unlock
}

func main() {
	go IncCount()
	go IncCount()
	time.Sleep(1 * time.Second)
}
*/