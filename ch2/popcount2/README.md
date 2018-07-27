# Exercise 2.4
"Write a version of [popcount](https://github.com/masonelmore/gopl/blob/b9a3ac1f943c43dbc7c55ffe9d8201ee1f62a628/ch2/popcount1/popcount.go) that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time.  Compare its performance to the table-lookup version."

# Results
```
$ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/masonelmore/gopl/ch2/ex2.4/popcount
BenchmarkPopCount/Table-12              300000000                4.17 ns/op
BenchmarkPopCount/Loop-12               100000000               18.4 ns/op
BenchmarkPopCount/Shift-12              20000000                69.5 ns/op
PASS
ok      github.com/masonelmore/gopl/ch2/ex2.4/popcount  5.063s
```