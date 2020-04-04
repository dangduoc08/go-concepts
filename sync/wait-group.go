package sync

import (
	"fmt"
	"sync"
	"time"
)

func parallelWaitGroup(wg *sync.WaitGroup) {
	fmt.Println("A")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("B")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("C")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("D")
	wg.Done()
}

// WaitGroup comment
func WaitGroup() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go parallelWaitGroup(wg)
	}
	wg.Wait()
}
