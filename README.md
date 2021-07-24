# max = 22.717866ms

```go
t0 := time.Now()
ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Millisecond)
ctx2, cancel2 := context.WithTimeout(ctx1, 1000*time.Millisecond)
<-ctx2.Done()
d := time.Since(t0)
```


```sh
go run .

min = 1.046524ms
max = 22.717866ms
ave = 2.090832ms 
[950 0 0 0 0 0 1 0 47 2] 
[1.046524ms 3.213658ms 5.380792ms 7.547926ms 9.71506ms 11.882194ms 14.049328ms 16.216462ms 18.383596ms 20.55073ms 22.717866ms]
linux amd64 go1.16.6

min = 1.041773ms
max = 2.99132ms
ave = 2.08696ms
[950 0 0 0 0 0 1 0 0 49]
[1.041773ms 2.99132ms 4.940867ms 6.890414ms 8.839961ms 10.789508ms 12.739055ms 14.688602ms 16.638149ms 18.587696ms 20.537247ms]
linux amd64 go1.16.6


# CPU usage < 10%  // main.g: lines 11 to 25 commented

min = 1.005733ms
max = 2.71599ms
ave = 1.151698ms
[99980 13 4 1 0 1 0 0 0 1]
[1.005733ms 2.71599ms 4.426247ms 6.136504ms 7.846761ms 9.557018ms 11.267275ms 12.977532ms 14.687789ms 16.398046ms 18.108303ms]
linux amd64 go1.16.6


go test -bench=. -benchtime=10s
# BenchmarkParentTimedoutContext-8  589  17178588 ns/op  // 17ms

# commented CPU usage lines:
# BenchmarkParentTimedoutContext-8  10000  1117480 ns/op // 1.1ms

go test -bench=. -benchtime=1000x
# BenchmarkParentTimedoutContext-8  1000  1154665 ns/op
go test -bench=. -benchtime=10000x
# BenchmarkParentTimedoutContext-8  10000  1120488 ns/op

go test -benchmem -run=^$ -bench ^(BenchmarkParentTimedoutContext)$ my
# BenchmarkParentTimedoutContext-8  1094   1101580 ns/op  672 B/op  10 allocs/op

```
