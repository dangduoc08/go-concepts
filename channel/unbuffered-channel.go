package channel

import "fmt"

func add(a, b int, ch chan<- int) {
	ch <- a + b
}

// UnbufferedChannel func
func UnbufferedChannel(a, b int) {
	ch := make(chan int)

	go add(a, b, ch)

	result := <-ch

	fmt.Println("Result of add function:", result)
}
