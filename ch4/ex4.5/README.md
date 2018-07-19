# Description
"Write an in-place function to eliminate adjacent duplicates in a `[]string` slice."

# Benchmarks
The `dupRunAppend` and `dupRunCopy` functions perform the best when there are long runs of duplicates, which makes sense because they are making fewer calls to `append` and `copy` respectively.  They all perform about the same when there are no adjacent duplicates, except for `dupAppend` by a small margin, which is the shortest function.
||dupAppend|dupCopy|dupRunAppend|dupRunCopy|
|-|-|-|-|-|
|LongSingleRun|816 ns/op|769 ns/op|164 ns/op|150 ns/op|
|RepeatingRuns|451 ns/op|438 ns/op|204 ns/op|201 ns/op|
|NoRuns|228 ns/op|246 ns/op|248 ns/op|251 ns/op|
