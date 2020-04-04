package sync

import (
	"fmt"
	"sync"
	"time"
)

func parallelLock(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("A")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("B")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("C")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("D")
	wg.Done()
}

// Lock comment
func Lock() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	var mu *sync.Mutex = &sync.Mutex{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go parallelLock(wg, mu)
	}
	wg.Wait()
}
