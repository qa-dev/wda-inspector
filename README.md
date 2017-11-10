# WDA-inspector [![Build Status](https://travis-ci.org/qa-dev/wda-inspector.png?branch=master)](https://travis-ci.org/qa-dev/wda-inspector)
=============

**wda-inspector** is a useful inspector of [qa-dev/WebDriverAgent](https://github.com/qa-dev/WebDriverAgent) (forked from [facebook/WebDriverAgent](https://github.com/facebook/WebDriverAgent))
written in Go.
 
## Build
Clone wda-inspector to `$GOPATH` according to [Golang code organization rules](https://golang.org/doc/code.html#Organization)

```
$ cd $GOPATH/src/github.com/qa-dev/wda-inspector
$ go build
```


## Usage

```
$ ./wda-inspector -h=127.0.0.1 -bundleId=com.apple.mobilesafari
```
_After run of wda-inspector, use link [http://127.0.0.1:8888] for open inspector in web browser._


### Flags:

| Name              | Type   | Required | Default | Description                      |
| ----------------- | ------ |--------- | --------|--------------------------------- |
| -bundleId         | string | true     |         |Bundle Id                         |
| -h                | string | true     |         |WDA host                          |
| -p                | string | false    | "8100"  |WDA port                          |
| -l                | string | false    | "8888"  |Port for handling                 |
