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

min = 1.041773ms
max = 2.99132ms
ave = 2.08696ms
[950 0 0 0 0 0 1 0 0 49]
[1.041773ms 2.99132ms 4.940867ms 6.890414ms 8.839961ms 10.789508ms 12.739055ms 14.688602ms 16.638149ms 18.587696ms 20.537247ms]

```
