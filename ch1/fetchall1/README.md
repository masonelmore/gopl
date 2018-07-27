# Exercise 1.10
"Find a web site that produces a large amount of data.  Investigate caching by running [fetchall](https://github.com/adonovan/gopl.io/blob/b725d6015f980e94734da37e35ba0d943fc7532f/ch1/fetchall/main.go) twice in succession to see whether the reported time changes much.  Do you get the same content each time?  Modify `fetchall` to print its output to a  file so it can be examined."

# Results
I was able to observe a faster response when requesting `https://www.reddit.com` twice in succession.

```
$ go run main.go https://www.reddit.com
2018-06-12 11:24:21.5074152 -0500 CDT m=+0.003000101
0.82s   331901  https://www.reddit.com
0.82s elapsed

$ go run main.go https://www.reddit.com
2018-06-12 11:24:24.5515894 -0500 CDT m=+0.004000301
0.20s     1064  https://www.reddit.com
0.20s elapsed
```