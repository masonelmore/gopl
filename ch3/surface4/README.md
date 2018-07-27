# Exercise 3.4
"Follow the approach of the [Lissajous](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch1/lissajous/main.go) example in Section 1.7, construct a web server that computes sufaces and write SVG data to the client.  The server must set the `Content-Type` header like this:
```go
w.Header().Set("Content-Type", "image/svg+xml")
```
[...] Allow the client to specify values like height, width, and color as HTTP request parameters."
