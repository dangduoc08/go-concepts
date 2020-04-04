package synchronous

import (
	"fmt"
	"runtime"
	"sync"
)

func parallel(wg *sync.WaitGroup) {
	fmt.Println("A")
	fmt.Println("B")
	fmt.Println("C")
	wg.Done()
}

// MaxProcesses func
func MaxProcesses(numP int) {
	runtime.GOMAXPROCS(numP)

	wg := &sync.WaitGroup{}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go parallel(wg)
	}
	wg.Wait()
}
