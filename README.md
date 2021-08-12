# go-html-selector

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-html-selector?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-html-selector)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-html-selector)](https://goreportcard.com/report/github.com/thewizardplusplus/go-html-selector)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-html-selector.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-html-selector)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-html-selector/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-html-selector)

The library that implements collecting text, specified HTML tags, and their attributes from an HTML document.

## Features

- collecting from an HTML document:
  - HTML tags;
  - HTML attributes;
  - text (optional);
- options:
  - filters:
    - filtering a result:
      - by specified HTML tags:
        - with support of the universal tag (`*`);
      - by specified HTML attributes;
    - friendly representation of filters:
      - for parsing from JSON;
      - for definition as a code literal;
  - skipping empty entities (separately optional):
    - tags without attributes;
    - attributes with an empty value;
    - whitespace-only text;
- representing a result:
  - using the builder interface for building a result;
  - built-in builders:
    - with grouping HTML attributes by their tags;
    - with collecting only values of HTML attributes;
    - with collecting only text:
      - with support of merging with other builders;
- optimizations:
  - of searching for a right filter among others:
    - with removing duplicate attribute filters, if there are the same filters for the universal tag;
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

`htmlselector.SelectTags()` with the text builder:

```go
package main

import (
	"bytes"
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
				link #1: <a href="http://example.com/1">one</a>
				<video
					src="http://example.com/1.1"
					poster="http://example.com/1.2">
					Unable to embed video #1.
				</video>
			</li>
			<li>
				link #2: <a href="http://example.com/2">two</a>
				<video
					src="http://example.com/2.1"
					poster="http://example.com/2.2">
					Unable to embed video #2.
				</video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		"a":     {"href"},
		"video": {"src", "poster"},
	})

	var builder builders.StructuralBuilder
	var textBuilder builders.TextBuilder
	err := htmlselector.SelectTags(
		reader,
		filters,
		htmlselector.MultiBuilder{Builder: &builder, TextBuilder: &textBuilder},
		htmlselector.SkipEmptyText(),
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

	fmt.Println("text parts:")
	for _, textPart := range textBuilder.TextParts() {
		textPart = bytes.TrimSpace(textPart)
		fmt.Printf("  %q\n", textPart)
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
	// text parts:
	//   "link #1:"
	//   "one"
	//   "Unable to embed video #1."
	//   "link #2:"
	//   "two"
	//   "Unable to embed video #2."
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

`htmlselector.SelectTags()` with the universal tag:

```
BenchmarkSelectTags/structural_builder/universal_tag/10_tags/1020B-8         	   50000	     25557 ns/op	    18370 B/op	     126 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/100_tags/10.4K-8        	   10000	    311815 ns/op	   168108 B/op	    1206 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/1000_tags/108.8K-8      	    1000	   2407620 ns/op	  1437633 B/op	   12006 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/10000_tags/1.1M-8       	     100	  22065997 ns/op	 11180076 B/op	  120006 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/100000_tags/11.6M-8     	      10	 445495615 ns/op	244945105 B/op	 1200006 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/1000000_tags/120.6M-8   	       1	2307457875 ns/op	383996112 B/op	12000006 allocs/op
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

`htmlselector.SelectTags()` with the universal tag:

```
BenchmarkSelectTags/flatten_builder/universal_tag/10_tags/1020B-8            	  100000	     17471 ns/op	     9823 B/op	      46 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/100_tags/10.4K-8           	   10000	    140121 ns/op	    57790 B/op	     406 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/1000_tags/108.8K-8         	    1000	   1577272 ns/op	   802345 B/op	    4006 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/10000_tags/1.1M-8          	     100	  15151013 ns/op	  5776548 B/op	   40006 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/100000_tags/11.6M-8        	      10	 164593709 ns/op	 70617641 B/op	  400006 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/1000000_tags/120.6M-8      	       1	1620865858 ns/op	869134992 B/op	 4000007 allocs/op
```

### With Structural & Text Builders

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/text_builder/simple_markup/10_tags/430B-8         	  122985	     10339 ns/op	     9466 B/op	      45 allocs/op
BenchmarkSelectTags/text_builder/simple_markup/100_tags/4.4K-8        	   15040	     78519 ns/op	    48475 B/op	     405 allocs/op
BenchmarkSelectTags/text_builder/simple_markup/1000_tags/45.7K-8      	    1990	    813642 ns/op	   510223 B/op	    4005 allocs/op
BenchmarkSelectTags/text_builder/simple_markup/10000_tags/476.3K-8    	     160	   6728975 ns/op	  3817896 B/op	   40005 allocs/op
BenchmarkSelectTags/text_builder/simple_markup/100000_tags/4.8M-8     	      19	  75982499 ns/op	 56745964 B/op	  400005 allocs/op
BenchmarkSelectTags/text_builder/simple_markup/1000000_tags/50.3M-8   	       2	 664071080 ns/op	290158432 B/op	 4000005 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/text_builder/complex_markup/10_tags/1020B-8       	   68089	     20964 ns/op	    21599 B/op	     116 allocs/op
BenchmarkSelectTags/text_builder/complex_markup/100_tags/10.4K-8      	    7014	    167202 ns/op	    41712 B/op	    1106 allocs/op
BenchmarkSelectTags/text_builder/complex_markup/1000_tags/108.8K-8    	     717	   1663950 ns/op	   387312 B/op	   11006 allocs/op
BenchmarkSelectTags/text_builder/complex_markup/10000_tags/1.1M-8     	      69	  23245678 ns/op	 27798263 B/op	  110006 allocs/op
BenchmarkSelectTags/text_builder/complex_markup/100000_tags/11.6M-8   	       6	 173781906 ns/op	 38403312 B/op	 1100006 allocs/op
BenchmarkSelectTags/text_builder/complex_markup/1000000_tags/120.6M-8 	       1	1740418254 ns/op	384003328 B/op	11000006 allocs/op
```

`htmlselector.SelectTags()` with the universal tag:

```
BenchmarkSelectTags/text_builder/universal_tag/10_tags/1020B-8        	   61375	     19060 ns/op	     8168 B/op	     136 allocs/op
BenchmarkSelectTags/text_builder/universal_tag/100_tags/10.4K-8       	    6638	    258572 ns/op	   407046 B/op	    1306 allocs/op
BenchmarkSelectTags/text_builder/universal_tag/1000_tags/108.8K-8     	     615	   1775232 ns/op	   387360 B/op	   13006 allocs/op
BenchmarkSelectTags/text_builder/universal_tag/10000_tags/1.1M-8      	      57	  27496277 ns/op	 43162804 B/op	  130006 allocs/op
BenchmarkSelectTags/text_builder/universal_tag/100000_tags/11.6M-8    	       6	 187202961 ns/op	 38403378 B/op	 1300006 allocs/op
BenchmarkSelectTags/text_builder/universal_tag/1000000_tags/120.6M-8  	       1	1854043284 ns/op	384003408 B/op	13000006 allocs/op
```

## License

The MIT License (MIT)

Copyright &copy; 2020-2021 thewizardplusplus
