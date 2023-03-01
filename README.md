# go-httpprd

Package **httpprd** provides tools for **HTTP products**, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-httpprd

[![GoDoc](https://godoc.org/github.com/reiver/go-httpprd?status.svg)](https://godoc.org/github.com/reiver/go-httpprd)

## Example:

```go
import "github.com/reiver/go-httpprd"

// ...

var s string = "HTTP/1.1"

name, version, err := httpprd.Parse(s)
if nil != err {
	return err
}

fmt.Printf("HTTP-product name:    %q\n", name)
fmt.Printf("HTTP-product version: %q\n", version)


// Output:
// HTTP-product name:    "HTTP"
// HTTP-product version: "1.1"
```

## Another Example

```go
import "github.com/reiver/go-httpprd"

// ...

var s string = "finger/beta"

name, version, err := httpprd.Parse(s)
if nil != err {
	return err
}

fmt.Printf("HTTP-product name:    %q\n", name)
fmt.Printf("HTTP-product version: %q\n", version)


// Output:
// HTTP-product name:    "finger"
// HTTP-product version: "beta"
```
