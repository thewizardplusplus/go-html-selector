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
  - skipping empty entities (separately optional):
    - tags without attributes;
    - attributes with an empty value;
- representing a result:
  - using the builder interface for building a result;
  - built-in builders:
    - with grouping HTML attributes by their tags;
    - with collecting only values of HTML attributes;
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

`htmlselector.SelectTags()` with the universal tag:

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
				<a href="http://example.com/1" title="link #1">1</a>
				<video
					src="http://example.com/1.1"
					poster="http://example.com/1.2"
					title="video #1">
				</video>
			</li>
			<li>
				<a href="http://example.com/2" title="link #2">2</a>
				<video
					src="http://example.com/2.1"
					poster="http://example.com/2.2"
					title="video #2">
				</video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		htmlselector.UniversalTag: {"title", "href"},
		"video":                   {"src", "poster"},
	})

	var builder builders.StructuralBuilder
	err := htmlselector.SelectTags(
		reader,
		filters,
		&builder,
		htmlselector.SkipEmptyTags(),
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
	//   title="link #1"
	// <video>:
	//   src="http://example.com/1.1"
	//   poster="http://example.com/1.2"
	//   title="video #1"
	// <a>:
	//   href="http://example.com/2"
	//   title="link #2"
	// <video>:
	//   src="http://example.com/2.1"
	//   poster="http://example.com/2.2"
	//   title="video #2"
}
```

## Benchmarks

### With Structural Builder

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/structural_builder/simple_markup/10_tags/430B-8         	  200000	     12572 ns/op	     8124 B/op	      35 allocs/op
BenchmarkSelectTags/structural_builder/simple_markup/100_tags/4.4K-8        	   20000	     93708 ns/op	    49167 B/op	     305 allocs/op
BenchmarkSelectTags/structural_builder/simple_markup/1000_tags/45.7K-8      	    3000	    716388 ns/op	   367374 B/op	    3005 allocs/op
BenchmarkSelectTags/structural_builder/simple_markup/10000_tags/476.3K-8    	     300	   7021473 ns/op	  2772749 B/op	   30005 allocs/op
BenchmarkSelectTags/structural_builder/simple_markup/100000_tags/4.8M-8     	      20	  59393679 ns/op	  8104370 B/op	  300005 allocs/op
BenchmarkSelectTags/structural_builder/simple_markup/1000000_tags/50.3M-8   	       1	1166593360 ns/op	998999904 B/op	 3000006 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/structural_builder/complex_markup/10_tags/1020B-8       	  100000	     18614 ns/op	     8117 B/op	     106 allocs/op
BenchmarkSelectTags/structural_builder/complex_markup/100_tags/10.4K-8      	   10000	    194201 ns/op	   155986 B/op	    1006 allocs/op
BenchmarkSelectTags/structural_builder/complex_markup/1000_tags/108.8K-8    	    1000	   1725249 ns/op	   379632 B/op	   10006 allocs/op
BenchmarkSelectTags/structural_builder/complex_markup/10000_tags/1.1M-8     	     100	  23956279 ns/op	 18179332 B/op	  100006 allocs/op
BenchmarkSelectTags/structural_builder/complex_markup/100000_tags/11.6M-8   	      10	 184968462 ns/op	 38395632 B/op	 1000006 allocs/op
BenchmarkSelectTags/structural_builder/complex_markup/1000000_tags/120.6M-8 	       1	1857636429 ns/op	383995632 B/op	10000006 allocs/op
```

### With Flatten Builder

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/flatten_builder/simple_markup/10_tags/430B-8            	  200000	      6244 ns/op	     5971 B/op	      15 allocs/op
BenchmarkSelectTags/flatten_builder/simple_markup/100_tags/4.4K-8           	   30000	     54042 ns/op	    21038 B/op	     105 allocs/op
BenchmarkSelectTags/flatten_builder/simple_markup/1000_tags/45.7K-8         	    3000	    530844 ns/op	   191624 B/op	    1005 allocs/op
BenchmarkSelectTags/flatten_builder/simple_markup/10000_tags/476.3K-8       	     300	   5116874 ns/op	  1402546 B/op	   10005 allocs/op
BenchmarkSelectTags/flatten_builder/simple_markup/100000_tags/4.8M-8        	      20	  54404882 ns/op	 23420178 B/op	  100005 allocs/op
BenchmarkSelectTags/flatten_builder/simple_markup/1000000_tags/50.3M-8      	       2	 663629704 ns/op	284703000 B/op	 1000005 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/flatten_builder/complex_markup/10_tags/1020B-8          	  100000	     15029 ns/op	     11411 B/op	      46 allocs/op
BenchmarkSelectTags/flatten_builder/complex_markup/100_tags/10.4K-8         	   10000	    150423 ns/op	     90183 B/op	     406 allocs/op
BenchmarkSelectTags/flatten_builder/complex_markup/1000_tags/108.8K-8       	    1000	   1268568 ns/op	     74576 B/op	    4006 allocs/op
BenchmarkSelectTags/flatten_builder/complex_markup/10000_tags/1.1M-8        	     100	  19005918 ns/op	  10593772 B/op	   40006 allocs/op
BenchmarkSelectTags/flatten_builder/complex_markup/100000_tags/11.6M-8      	      10	 138218081 ns/op	   7442576 B/op	  400006 allocs/op
BenchmarkSelectTags/flatten_builder/complex_markup/1000000_tags/120.6M-8    	       1	2141494094 ns/op	1313346192 B/op	 4000007 allocs/op
```

## License

The MIT License (MIT)

Copyright &copy; 2020 thewizardplusplus
