# Exercise 4.7
"Modify [reverse](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch4/rev/main.go) to reverse the characters of a `[]byte` slice that represents a UTF-8 encoded string, in place.  Can you do it without allocating new memory?"

# Benchmarks
Reversing the entire `[]byte` slice before reversing the multi-byte runes seems to perform better than reversing the multi-byte runes first.  My first guess is that `utf8.DecodeRune` in `reverse` is slowing things down a tad.  I will have to fire up pprof to verify that.

|Benchmark|reverse|reverse2|
|-|-|-|
|ASCII short|33.1 ns/op|13.5 ns/op|
|ASCII medium|129 ns/op|41.5 ns/op|
|ASCII long|726 ns/op|236 ns/op|
|ASCII super-long|6425 ns/op|1935 ns/op|
|UTF8 short|57.3 ns/op|49.6 ns/op|
|UTF8 medium|219 ns/op|185 ns/op|
|UTF8 long|1257 ns/op|1090 ns/op|
|UTF8 super-long|11295 ns/op|9655 ns/op|
