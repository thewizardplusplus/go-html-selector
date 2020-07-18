# go-html-selector

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-html-selector?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-html-selector)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-html-selector)](https://goreportcard.com/report/github.com/thewizardplusplus/go-html-selector)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-html-selector.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-html-selector)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-html-selector/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-html-selector)

The library that implements collecting specified HTML tags and their attributes from an HTML document.

## Features

- collecting from an HTML document:
  - HTML tags;
  - HTML attributes;
- options:
  - filtering a result:
    - by specified HTML tags;
    - by specified HTML attributes;
- representing a result:
  - grouping HTML attributes by their tags;
- optimizations by the number:
  - of memory allocations;
  - of string copies.

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

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		"a":     {"href"},
		"video": {"src", "poster"},
	})

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

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/simple_markup/10_tags/430B-8         	 2000000	       826 ns/op	    4304 B/op	       2 allocs/op
BenchmarkSelectTags/simple_markup/100_tags/4.4K-8        	 2000000	       825 ns/op	    4304 B/op	       2 allocs/op
BenchmarkSelectTags/simple_markup/1000_tags/45.7K-8      	 2000000	       828 ns/op	    4304 B/op	       2 allocs/op
BenchmarkSelectTags/simple_markup/10000_tags/476.3K-8    	 2000000	       896 ns/op	    4305 B/op	       2 allocs/op
BenchmarkSelectTags/simple_markup/100000_tags/4.8M-8     	 2000000	       797 ns/op	    4321 B/op	       2 allocs/op
BenchmarkSelectTags/simple_markup/1000000_tags/50.3M-8   	  500000	      2325 ns/op	    4983 B/op	      10 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/complex_markup/10_tags/1020B-8       	 2000000	       895 ns/op	     4304 B/op	       2 allocs/op
BenchmarkSelectTags/complex_markup/100_tags/10.4K-8      	 2000000	       827 ns/op	     4304 B/op	       2 allocs/op
BenchmarkSelectTags/complex_markup/1000_tags/108.8K-8    	 2000000	       829 ns/op	     4304 B/op	       2 allocs/op
BenchmarkSelectTags/complex_markup/10000_tags/1.1M-8     	 2000000	       925 ns/op	     4308 B/op	       2 allocs/op
BenchmarkSelectTags/complex_markup/100000_tags/11.6M-8   	 2000000	       828 ns/op	     4349 B/op	       2 allocs/op
BenchmarkSelectTags/complex_markup/1000000_tags/120.6M-8 	       1	2228727318 ns/op	881045216 B/op	14000048 allocs/op
```

## License

The MIT License (MIT)

Copyright &copy; 2020 thewizardplusplus
