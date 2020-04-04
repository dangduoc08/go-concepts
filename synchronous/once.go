package synchronous

import (
	"fmt"
	"sync"
	"time"
)

var flag int32 = 0

func parallelOnce(wg *sync.WaitGroup, once *sync.Once) {
	once.Do(func() {
		flag++
		fmt.Println("Flag:", flag)
	})
	fmt.Println("A")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("B")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("C")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("D")
	wg.Done()
}

// Once func
func Once() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	var once *sync.Once = &sync.Once{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go parallelOnce(wg, once)
		wg.Wait()
	}
}
