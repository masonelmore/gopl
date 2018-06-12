# Description
"The function call to `io.Copy(dst, src)` reads from `src` and writes to `dst`.  Use it instead of `ioutil.ReadAll` to copy the response body to `os.Stdout` without requiring a buffer large enough to hold the entire stream.  Be sure to check the error result of `io.Copy`."

[fetch](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch1/fetch/main.go)
