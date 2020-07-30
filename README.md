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
  - filters:
    - filtering a result:
      - by specified HTML tags;
      - by specified HTML attributes;
    - friendly representation of filters:
      - for parsing from JSON;
      - for definition as a code literal;
  - skipping empty tags (i.e. without attributes; optional);
- representing a result:
  - using the builder interface for building a result;
  - built-in builders:
    - with grouping HTML attributes by their tags;
- optimizations:
  - of searching for the right one among filters;
  - of conversion from a byte slice to a string;
  - by the number:
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

`htmlselector.SelectTags()` with the structural builder:

```go
package main

import (
	"fmt"
	"log"
	"strings"

	htmlselector "github.com/thewizardplusplus/go-html-selector"
	"github.com/thewizardplusplus/go-html-selector/builders"
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
			<li>
				<a href="">3</a>
				<video src="" poster=""></video>
			</li>
			<li>
				<a>4</a>
				<video></video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		"a":     {"href"},
		"video": {"src", "poster"},
	})

	var builder builders.StructuralBuilder
	err := htmlselector.SelectTags(
		reader,
		filters,
		&builder,
		htmlselector.SkipEmptyTags(),
		htmlselector.SkipEmptyAttributes(),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, tag := range builder.Tags() {
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

`htmlselector.SelectTags()` with the flatten builder:

```go
package main

import (
	"fmt"
	"log"
	"strings"

	htmlselector "github.com/thewizardplusplus/go-html-selector"
	"github.com/thewizardplusplus/go-html-selector/builders"
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
			<li>
				<a href="">3</a>
				<video src="" poster=""></video>
			</li>
			<li>
				<a>4</a>
				<video></video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		"a":     {"href"},
		"video": {"src", "poster"},
	})

	var builder builders.FlattenBuilder
	err := htmlselector.SelectTags(
		reader,
		filters,
		&builder,
		htmlselector.SkipEmptyTags(),
		htmlselector.SkipEmptyAttributes(),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, attributeValue := range builder.AttributeValues() {
		fmt.Printf("%q\n", attributeValue)
	}

	// Output:
	// "http://example.com/1"
	// "http://example.com/1.1"
	// "http://example.com/1.2"
	// "http://example.com/2"
	// "http://example.com/2.1"
	// "http://example.com/2.2"
}
```

## Benchmarks

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/simple_markup/10_tags/430B-8         	  200000	      9474 ns/op	     6784 B/op	      51 allocs/op
BenchmarkSelectTags/simple_markup/100_tags/4.4K-8        	   20000	     72306 ns/op	    25456 B/op	     414 allocs/op
BenchmarkSelectTags/simple_markup/1000_tags/45.7K-8      	    2000	    621791 ns/op	   190672 B/op	    4017 allocs/op
BenchmarkSelectTags/simple_markup/10000_tags/476.3K-8    	     200	   7247563 ns/op	  3448480 B/op	   40027 allocs/op
BenchmarkSelectTags/simple_markup/100000_tags/4.8M-8     	      20	  80736482 ns/op	 35420205 B/op	  400037 allocs/op
BenchmarkSelectTags/simple_markup/1000000_tags/50.3M-8   	       2	 802693264 ns/op	339752800 B/op	 4000047 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/complex_markup/10_tags/1020B-8       	  100000	     21746 ns/op	    11264 B/op	     153 allocs/op
BenchmarkSelectTags/complex_markup/100_tags/10.4K-8      	   10000	    187676 ns/op	    67328 B/op	    1416 allocs/op
BenchmarkSelectTags/complex_markup/1000_tags/108.8K-8    	    1000	   1823346 ns/op	   740608 B/op	   14021 allocs/op
BenchmarkSelectTags/complex_markup/10000_tags/1.1M-8     	     100	  21162013 ns/op	  9136409 B/op	  140031 allocs/op
BenchmarkSelectTags/complex_markup/100000_tags/11.6M-8   	       5	 227873490 ns/op	 90800432 B/op	 1400041 allocs/op
BenchmarkSelectTags/complex_markup/1000000_tags/120.6M-8 	       1	2402936519 ns/op	881045280 B/op	14000051 allocs/op
```

## License

The MIT License (MIT)

Copyright &copy; 2020 thewizardplusplus
