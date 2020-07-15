# go-html-selector

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-html-selector?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-html-selector)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-html-selector)](https://goreportcard.com/report/github.com/thewizardplusplus/go-html-selector)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-html-selector.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-html-selector)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-html-selector/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-html-selector)

## Installation

Prepare the directory:

```
$ mkdir --parents "$(go env GOPATH)/src/github.com/thewizardplusplus/"
$ cd "$(go env GOPATH)/src/github.com/thewizardplusplus/"
```

Clone this repository:

```
$ git clone https://github.com/thewizardplusplus/go-html-selector.git
$ cd go-html-selector
```

Install dependencies with the [dep](https://golang.github.io/dep/) tool:

```
$ dep ensure -vendor-only
```

## Examples

`htmlselector.SelectTags()`:

```go
package main

import (
	"fmt"
	"log"
	"strings"

	htmlselector "github.com/thewizardplusplus/go-html-selector"
)

func main() {
	reader := strings.NewReader(`
		<ul>
			<li>
				<a href="http://example.com/1">1</a>
				<video
					src="http://example.com/1.1"
					poster="http://example.com/1.2">
				</video>
			</li>
			<li>
				<a href="http://example.com/2">2</a>
				<video
					src="http://example.com/2.1"
					poster="http://example.com/2.2">
				</video>
			</li>
		</ul>
	`)

	filters := []htmlselector.Filter{
		{
			Tag:        []byte("a"),
			Attributes: [][]byte{[]byte("href")},
		},
		{
			Tag:        []byte("video"),
			Attributes: [][]byte{[]byte("src"), []byte("poster")},
		},
	}

	tags, err := htmlselector.SelectTags(reader, filters)
	if err != nil {
		log.Fatal(err)
	}

	for _, tag := range tags {
		fmt.Printf("<%s>:\n", tag.Name)
		for _, attribute := range tag.Attributes {
			fmt.Printf("  %s=%q\n", attribute.Name, attribute.Value)
		}
	}

	// Output:
	// <a>:
	//   href="http://example.com/1"
	// <video>:
	//   src="http://example.com/1.1"
	//   poster="http://example.com/1.2"
	// <a>:
	//   href="http://example.com/2"
	// <video>:
	//   src="http://example.com/2.1"
	//   poster="http://example.com/2.2"
}
```

## Benchmarks

`htmlselector.SelectTags()`:

```
BenchmarkDistance/10Tags/490.00B-8         	 2000000	       836 ns/op	    4304 B/op	       2 allocs/op
BenchmarkDistance/100Tags/4.87KiB-8        	 2000000	       840 ns/op	    4304 B/op	       2 allocs/op
BenchmarkDistance/1000Tags/50.58KiB-8      	 2000000	       829 ns/op	    4304 B/op	       2 allocs/op
BenchmarkDistance/10000Tags/525.19KiB-8    	 2000000	       902 ns/op	    4305 B/op	       2 allocs/op
BenchmarkDistance/100000Tags/5.32MiB-8     	 2000000	       855 ns/op	    4321 B/op	       2 allocs/op
BenchmarkDistance/1000000Tags/55.10MiB-8   	  300000	      3505 ns/op	    5436 B/op	      15 allocs/op
```

## License

The MIT License (MIT)

Copyright &copy; 2020 thewizardplusplus
