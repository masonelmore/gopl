# Description
"Modify the Lissajous server to read parameter values from the URL.  For example, you might arrange it so that a URL like `http://localhost:8000/?cycles=20` sets the number of cycles to 20 instead of the default 5.  Use the `strconv.Atoi` function to convert the string parameter into an integer.  You can see its documentation with `go doc strconv.Atoi`."

[lissajous](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch1/lissajous/main.go)
