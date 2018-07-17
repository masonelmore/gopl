# Description
"The expression `x&(x-1)` clears the rightmost non-zero bit of `x`.  Write a version of `PopCount` that counts bits by using this fact, and asses its performance."

[popcount](https://github.com/masonelmore/gopl/tree/14fe9551820d3a172398464b55f17c06912ed85d/ch2/ex2.4)

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
