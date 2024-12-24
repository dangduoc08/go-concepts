package main

import (
	"sync"
	"time"
)

type Student struct {
	Name      string
	Age       int32
	CreatedAt time.Time
}

func NewStudent() *Student {
	s := &Student{
		Name:      "John Doe",
		Age:       29,
		CreatedAt: time.Now(),
	}
	// s.CreatedAt = time.Now()
	return s
}

func (s *Student) Reset() {
	s.Name = ""
	s.Age = 0
}

var studentPool = sync.Pool{
	New: func() any {
		s := NewStudent()
		s.CreatedAt = time.Now()
		return s
	},
}

func main() {
}
