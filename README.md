wda-inspector [![Build Status](https://travis-ci.org/qa-dev/wda-inspector.png?branch=master)](https://travis-ci.org/qa-dev/wda-inspector)
=============

**wda-inspector** is a useful inspector of [qa-dev/WebDriverAgent](https://github.com/qa-dev/WebDriverAgent) (forked from [facebook/WebDriverAgent](https://github.com/facebook/WebDriverAgent))
written in Go.
 
Build
-----
Clone wda-inspector to `$GOPATH` according to [Golang code organization rules](https://golang.org/doc/code.html#Organization)

```
$ cd $GOPATH/src/github.com/qa-dev/wda-inspector
$ go build
```


Usage
-----

```
$ ./wda-inspector -h=127.0.0.1 -bundleId=com.apple.mobilesafari
```

Flags:
```
-bundleId string
    	Bundle Id
-h string
    	WDA host
-p string
    	WDA port. optional (default "8100")
-l string
    	Port to listen. optional (default "8888")    	
```
