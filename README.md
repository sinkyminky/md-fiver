## MD-Fiver

This project get urls and parallel execution count from cli then requests the given urls with in 
given parallel execution count and print MD5 hash of the responses with the urls.

PS: You do not have to provide parallel execution count. Default count is 10.

## Installation ##

- Golang must be installed to your computer
- Go to under cmd/mdfiver folder and run ``go build``

## Usage ##

```ssh
$ ./mdfiver -parallel 3 google.com facebook.com http://google.com https://reddit.com https://www.twitter.com
```
Output:
```ssh
http://facebook.com   c3bc6ad2c92d18ca614daff370d7155a
http://google.com   62671d6663bb3d4403315ebe6911480b
http://google.com   ea74b81162fd3e8582e476bf689bf6e8
https://www.twitter.com   be15fa0fba0889f2c3b9cce5016cec8d
https://reddit.com   3634242f2d9fdf889f338b3ed6b0cadb
```
OR:
```ssh
$ ./mdfiver google.com facebook.com http://google.com https://reddit.com https://www.twitter.com
```
Output:
```ssh
http://google.com   db6a445a21eb022a7a4cfce1bc46b70b
http://google.com   1e6b180ea80e4fa48afdca4eeb25ffc2
http://facebook.com   bf136c9e5b4fa3bc7dced9a9ade365e4
https://www.twitter.com   699f04410f7db64fa8f083bf49007a2e
https://reddit.com   ce4095d3337ba50f3ac4af5f5aa881f5
```


