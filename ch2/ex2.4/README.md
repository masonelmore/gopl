# Description
"Write a version of `PopCount` that counts bits by shifting its argument through 64 bit positions, testing the rightmost bit each time.  Compare its performance to the table-lookup version."

[popcount](https://github.com/masonelmore/gopl/tree/f66a00fa58dfd353779e621a86f724f29dc34a9d/ch2/ex2.3)

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