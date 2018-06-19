# Description
"Follow the approach of the Lissajous example in Section 1.7, construct a web server that computes sufaces and write SVG data to the client.  The server must set the `Content-Type` header like this:
```go
w.Header().Set("Content-Type", "image/svg+xml")
```
[...] Allow the client to specify values like height, width, and color as HTTP request parameters."

