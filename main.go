package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	const n = 1_000
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			ctx3, cancel3 := context.WithTimeout(context.Background(), n*time.Millisecond)
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

	all := make([]time.Duration, n)
	for i := 0; i < n; i++ {
		t0 := time.Now()
		ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Millisecond)
		ctx2, cancel2 := context.WithTimeout(ctx1, 1000*time.Millisecond)
		<-ctx2.Done()
		d := time.Since(t0)
		cancel1()
		cancel2()
		all[i] = d
	}

	ave, x, y := histogram(all, 10)
	fmt.Println("min =", y[0])
	fmt.Println("max =", y[1])
	fmt.Println("ave =", ave)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.Version())
}

func histogram(d []time.Duration, n int) (ave time.Duration, x []int, y []time.Duration) {
	x = make([]int, n)
	y = make([]time.Duration, n+1)
	min := d[0]
	max := min
	ave = min
	for _, v := range d[1:] {
		ave += v
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	ave /= time.Duration(len(d))
	distance := (max - min) / time.Duration(n)
	v := min
	for i := range y {
		y[i] = v
		v += distance
	}
	y[len(y)-1] = max // compansate division error
	for _, v := range d {
		i := int((v - min) / distance)
		if i >= n {
			i--
		}
		x[i]++
	}
	return
}
