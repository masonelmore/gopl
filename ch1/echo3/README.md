# Exercise 1.3
"Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses `strings.Join`.  (Section 1.6 illustrates part of the `time` package, and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)"

See [echo1](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch1/echo1/main.go) and [echo3](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch1/echo3/main.go) for the different versions.

# Benchmarks

```
$ go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: github.com/masonelmore/gopl/ch1/ex1.3
benchmark                            iter              time/iter        bytes alloc             allocs
---------                            ----              ---------        -----------             ------
BenchmarkEchoConcat/0-12        500000000             3.22 ns/op             0 B/op        0 allocs/op
BenchmarkEchoConcat/10-12         2000000           798.00 ns/op           416 B/op        9 allocs/op
BenchmarkEchoConcat/100-12          50000         30661.00 ns/op         30960 B/op       99 allocs/op
BenchmarkEchoConcat/1000-12           500       2756157.00 ns/op       3238064 B/op      999 allocs/op
BenchmarkEchoConcat/10000-12            5     202611580.00 ns/op     325934566 B/op     9999 allocs/op
BenchmarkEchoConcat/100000-12           1   13021744800.00 ns/op   30350007968 B/op   100082 allocs/op
BenchmarkEchoJoin/0-12          300000000             5.87 ns/op             0 B/op        0 allocs/op
BenchmarkEchoJoin/10-12           5000000           260.00 ns/op           160 B/op        2 allocs/op
BenchmarkEchoJoin/100-12          1000000          1888.00 ns/op          1280 B/op        2 allocs/op
BenchmarkEchoJoin/1000-12          100000         12110.00 ns/op         12288 B/op        2 allocs/op
BenchmarkEchoJoin/10000-12          10000        183210.00 ns/op        131072 B/op        2 allocs/op
BenchmarkEchoJoin/100000-12          1000       1871107.00 ns/op       1212416 B/op        2 allocs/op
ok      github.com/masonelmore/gopl/ch1/ex1.3   33.360s
```
