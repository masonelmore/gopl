# Description
"Try `fetchall` with longer argument lists, such as samples from the top million web sites available at alexa.com.  How does the program behave if a web site just doesn't respond? (Section 8.9 describes mechanisms for coping in such cases.)"

[fetchall](https://github.com/masonelmore/gopl/blob/293c85111d03029f4612ca7f23ca92d05a34f1d5/ch1/ex1.10/main.go)

# Results
I initially tried running it with 10000 urls, but it was taking way too long to complete.  I imagine that trying to send that many requests at once might be a bit excessive.  I toned it down to 100 urls and I was still able to catch a ton of errors.
```
$ go run testdata/generate.go <path to alexa csv> 100 | xargs go run main.go
$ cat results
2018-06-12 20:47:55.1884472 -0500 CDT m=+0.003000201
0.00s    ERROR  http://googleusercontent.com
0.15s     3180  http://t.co
0.16s    11443  http://google.co.jp
0.17s    10865  http://google.co.in
0.18s    11427  http://google.com.hk
0.18s   366986  http://coinmarketcap.com
0.18s   123421  http://bing.com
0.19s    29286  http://twitch.tv
0.19s    11412  http://google.co.id
0.20s    62096  http://diply.com
0.20s    10814  http://google.com.pk
0.20s    11406  http://google.com.tr
0.20s    11419  http://google.de
0.21s    10839  http://google.com.au
0.21s    11438  http://google.com.sa
0.22s    11416  http://google.com.ua
0.23s    11461  http://google.com.br
0.24s    43614  http://msn.com
0.23s     3638  http://live.com
0.24s     1064  http://reddit.com
0.23s    10834  http://google.co.uk
0.24s   300992  http://wikia.com
0.25s    11391  http://google.es
0.25s    16870  http://deloton.com
0.25s    70770  http://wordpress.com
0.25s    11456  http://google.com.mx
0.26s    10869  http://google.ca
0.26s    11417  http://google.fr
0.30s    11545  http://google.com
0.30s    11409  http://google.pl
0.30s    11476  http://google.com.ar
0.35s    50073  http://apple.com
0.36s    11420  http://google.ru
0.36s    76274  http://tumblr.com
0.37s    11428  http://google.com.tw
0.39s   261073  http://stackoverflow.com
0.41s    54808  http://github.com
0.44s   149751  http://twitter.com
0.45s   416382  http://whatsapp.com
0.51s   118608  http://adobe.com
0.53s       81  http://baidu.com
0.55s   153676  http://imgur.com
0.55s    75576  http://wikipedia.org
0.55s    64051  http://pinterest.com
0.55s  1644566  http://espn.com
0.57s    14548  http://popads.net
0.59s    21344  http://instagram.com
0.60s    11411  http://google.it
0.62s    94591  http://blogspot.com
0.62s   298300  http://pornhub.com
0.63s    71769  http://netflix.com
0.63s   245788  http://qq.com
0.64s   148820  http://360.cn
0.66s   195617  http://imdb.com
0.68s    45520  http://linkedin.com
0.70s   502399  http://yahoo.com
0.76s   264092  http://ebay.com
0.76s    66081  http://office.com
0.79s    38522  http://aliexpress.com
0.78s   218577  http://sohu.com
0.90s   449735  http://facebook.com
0.91s   244311  http://tmall.com
0.93s   595290  http://amazon.com
0.94s   159558  http://xhamster.com
0.96s     6515  http://amazon.co.jp
0.98s     7867  http://tianya.cn
1.04s    43230  http://coccoc.com
1.06s   250297  http://bongacams.com
1.06s   141998  http://gmw.cn
1.08s   104907  http://nicovideo.jp
1.14s   244311  http://list.tmall.com
1.16s   124121  http://jd.com
1.20s   162383  http://microsoft.com
1.23s    89714  http://yandex.ru
1.24s   242009  http://taobao.com
1.31s     5751  http://pixnet.net
1.31s    44877  http://paypal.com
1.33s   198812  http://livejasmin.com
1.35s     8909  http://vk.com
1.45s    66075  http://xnxx.com
1.50s   283273  http://mail.ru
1.51s    20028  http://yahoo.co.jp
1.65s   496182  http://youtube.com
1.74s   663598  http://amazon.in
1.76s   110490  http://xvideos.com
1.79s   438831  http://amazon.co.uk
1.80s   385248  http://amazon.de
1.94s   591833  http://sina.com.cn
2.31s   124720  http://csdn.net
2.43s    21651  http://alipay.com
2.63s    23753  http://soso.com
3.16s   544641  http://hao123.com
3.18s   291174  http://so.com
3.29s    90601  http://dropbox.com
3.34s    95241  http://weibo.com
3.49s   133738  http://savefrom.net
6.26s    ERROR  http://detail.tmall.com
25.74s   174610  http://ok.ru
30.00s    ERROR  http://microsoftonline.com
30.81s     4822  http://thepiratebay.org
30.83s elapsed
```