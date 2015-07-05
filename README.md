goss
====

Go Screenshot Package for OS X

Requirements
------------
* OS X >= 10.8
* Go >= 1.4

How to use
----------
```go
image, err := goss.Capture()
```

or

```go
ss, err := goss.New()
rgba := ss.ToRGBA()
```

Refs
----
* [Son of Grab](https://developer.apple.com/library/mac/samplecode/SonOfGrab/Introduction/Intro.html)
* [vova616/screenshot](https://github.com/vova616/screenshot)
