# Description
"Write a program `wordfreq` to report the frequency of each word in an input text file.  Call `input.Split(bufio.ScanWords)` before the first call to `Scan` to break the input into words instead of lines"

# Results
```
$ go run main.go testdata\lorem_ipsum
Freq    Word
4       id
3       massa
3       pulvinar
2       aliquam
1       Fusce
2       finibus
4       Nulla
3       Duis
2       Suspendisse
4       ullamcorper
4       ut
3       turpis
3       rutrum
4       neque
1       facilisis
1       laoreet
4       commodo
3       tortor
5       pretium
1       Sed
[...]
```