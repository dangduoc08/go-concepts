package channel

import "fmt"

// Select func
func Select() {
	pub := make(chan int)
	sub := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			pub <- i
		}
		sub <- true
	}()

	for {
		select {
		case pubVal := <-pub:
			fmt.Println(pubVal)
		case <-sub:
			fmt.Println("Done")
			return
		}
	}
}
