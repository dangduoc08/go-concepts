package main

import (
	"testing"
)

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// c := NewStudent()
		// c.CreatedAt = time.Now()

		_ = NewStudent()
	}
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = studentPool.Get().(*Student)
	}
}

func BenchmarkWithAndReset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := studentPool.Get().(*Student)
		s.Reset()
	}
}

func BenchmarkWithAndResetAndPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := studentPool.Get().(*Student)
		s.Reset()
		studentPool.Put(s)
	}
}
