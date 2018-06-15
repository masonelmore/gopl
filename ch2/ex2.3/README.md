# Description
"Rewrite `PopCount` to use a loop instead of a single expression.  Compare the performance of the two versions. (Section 11.4 shows hwo to compare the performance of different implementations systematically.)"

[popcount](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch2/popcount/main.go)

# Results
```
$ go test -bench=.
goos: windows
goarch: amd64
pkg: github.com/masonelmore/gopl/ch2/ex2.3/popcount
BenchmarkPopCount/Table-12              300000000                4.09 ns/op
BenchmarkPopCount/Loop-12               100000000               18.2 ns/op
PASS
ok      github.com/masonelmore/gopl/ch2/ex2.3/popcount  3.537s
```
