package main

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func BenchmarkParentTimedoutContext(b *testing.B) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			ctx3, cancel3 := context.WithTimeout(context.Background(), 10_000*time.Millisecond)
			defer cancel3()
			for {
				select {
				case <-ctx3.Done():
					return
				default:
					// runtime.Gosched()
				}
			}
		}()
	}
	for i := 0; i < b.N; i++ {
		ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Millisecond)
		ctx2, cancel2 := context.WithTimeout(ctx1, 1000*time.Millisecond)
		<-ctx2.Done()
		cancel1()
		cancel2()
	}
}
