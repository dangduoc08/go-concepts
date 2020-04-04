package channel

import "fmt"

// BufferedChannel func
func BufferedChannel() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}
