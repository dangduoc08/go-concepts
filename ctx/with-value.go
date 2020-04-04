package ctx

import (
	"context"
	"fmt"
)

// WithValue func
func WithValue() {
	ctx := context.Background()
	ch := make(chan context.Context)
	arr := []string{
		"Gia tri 1",
		"Gia tri 2",
		"Gia tri 3",
		"Gia tri 4",
		"Gia tri 5",
		"Gia tri 6",
		"Gia tri 7",
		"Gia tri 8",
		"Gia tri 9",
		"Gia tri 10",
	}

	for _, val := range arr {
		go func(val string, ch chan<- context.Context, ctx context.Context) {
			ctx = context.WithValue(ctx, val, val)
			ch <- ctx
		}(val, ch, ctx)
	}

	i := 0

	for val := range ch {
		i++
		fmt.Println(val)

		if i == len(arr) {
			close(ch)
		}
	}
}
