# Exercise 2.5
"The expression `x&(x-1)` clears the rightmost non-zero bit of `x`.  Write a version of [popcount](https://github.com/masonelmore/gopl/blob/b9a3ac1f943c43dbc7c55ffe9d8201ee1f62a628/ch2/popcount2/popcount.go) that counts bits by using this fact, and asses its performance."

# Results
```
$ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/masonelmore/gopl/ch2/ex2.5/popcount
BenchmarkPopCount/Table-12              300000000                4.15 ns/op
BenchmarkPopCount/Loop-12               100000000               18.5 ns/op
BenchmarkPopCount/Shift-12              20000000                69.4 ns/op
BenchmarkPopCount/Clear-12              30000000                42.6 ns/op
PASS
ok      github.com/masonelmore/gopl/ch2/ex2.5/popcount  6.376s
```
