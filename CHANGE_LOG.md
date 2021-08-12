# Change Log

## [v1.7](https://github.com/thewizardplusplus/go-html-selector/tree/v1.7) (2021-08-12)

- options:
  - support of selection terminating ahead of time (optional);
- representing a result:
  - built-in builders:
    - support of merging any builder with a selection terminator.

## [v1.6.1](https://github.com/thewizardplusplus/go-html-selector/tree/v1.6.1) (2020-12-26)

- improving tests:
  - adding tests with attributes overlapping between tags for the `htmlselector.SelectTags()` function;
- refactoring:
  - extracting the `htmlselector.selector` structure;
  - moving builders from the `htmlselector` package to a separate file.

## [v1.6](https://github.com/thewizardplusplus/go-html-selector/tree/v1.6) (2020-12-08)

- collecting from an HTML document:
  - text (optional);
- skipping whitespace-only text (optional);
- representing a result:
  - built-in builders:
    - with collecting only text:
      - with support of merging with other builders.

### Benchmarks

#### With Structural & Text Builders

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

## [v1.5](https://github.com/thewizardplusplus/go-html-selector/tree/v1.5) (2020-11-27)

- filtering a result:
  - support of the universal tag (`*`);
- optimizations:
  - removing duplicate attribute filters, if there are the same filters for the universal tag.

### Benchmarks

#### With Structural Builder

`htmlselector.SelectTags()` with the universal tag:

```
BenchmarkSelectTags/structural_builder/universal_tag/10_tags/1020B-8         	   50000	     25557 ns/op	    18370 B/op	     126 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/100_tags/10.4K-8        	   10000	    311815 ns/op	   168108 B/op	    1206 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/1000_tags/108.8K-8      	    1000	   2407620 ns/op	  1437633 B/op	   12006 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/10000_tags/1.1M-8       	     100	  22065997 ns/op	 11180076 B/op	  120006 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/100000_tags/11.6M-8     	      10	 445495615 ns/op	244945105 B/op	 1200006 allocs/op
BenchmarkSelectTags/structural_builder/universal_tag/1000000_tags/120.6M-8   	       1	2307457875 ns/op	383996112 B/op	12000006 allocs/op
```

#### With Flatten Builder

`htmlselector.SelectTags()` with the universal tag:

```
BenchmarkSelectTags/flatten_builder/universal_tag/10_tags/1020B-8            	  100000	     17471 ns/op	     9823 B/op	      46 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/100_tags/10.4K-8           	   10000	    140121 ns/op	    57790 B/op	     406 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/1000_tags/108.8K-8         	    1000	   1577272 ns/op	   802345 B/op	    4006 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/10000_tags/1.1M-8          	     100	  15151013 ns/op	  5776548 B/op	   40006 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/100000_tags/11.6M-8        	      10	 164593709 ns/op	 70617641 B/op	  400006 allocs/op
BenchmarkSelectTags/flatten_builder/universal_tag/1000000_tags/120.6M-8      	       1	1620865858 ns/op	869134992 B/op	 4000007 allocs/op
```

## [v1.4](https://github.com/thewizardplusplus/go-html-selector/tree/v1.4) (2020-07-30)

- skipping empty attributes (i.e. with an empty value; optional);
- representing a result:
  - built-in builders:
    - with collecting only values of HTML attributes;
- optimizations:
  - optimize the `htmlselector.SelectTags()` function;
  - optimize the `byteutils.Copy()` function;
  - optimize the `builders.StructuralBuilder.AddAttribute()` method.

### Benchmarks

#### With Structural Builder

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

#### With Flatten Builder

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

## [v1.3.1](https://github.com/thewizardplusplus/go-html-selector/tree/v1.3.1) (2020-07-30)

- fix benchmarks.

## [v1.3](https://github.com/thewizardplusplus/go-html-selector/tree/v1.3) (2020-07-30)

- representing a result:
  - using the builder interface for building a result;
  - built-in builders:
    - with grouping HTML attributes by their tags.

### Benchmarks

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

## [v1.2](https://github.com/thewizardplusplus/go-html-selector/tree/v1.2) (2020-07-21)

- skipping empty tags (i.e. without attributes; optional).

### Benchmarks

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/simple_markup/10_tags/430B-8         	  200000	      9796 ns/op	     6736 B/op	      50 allocs/op
BenchmarkSelectTags/simple_markup/100_tags/4.4K-8        	   20000	     66461 ns/op	    25408 B/op	     413 allocs/op
BenchmarkSelectTags/simple_markup/1000_tags/45.7K-8      	    2000	    614359 ns/op	   190624 B/op	    4016 allocs/op
BenchmarkSelectTags/simple_markup/10000_tags/476.3K-8    	     200	   6833663 ns/op	  3448433 B/op	   40026 allocs/op
BenchmarkSelectTags/simple_markup/100000_tags/4.8M-8     	      20	  77550370 ns/op	 35420159 B/op	  400036 allocs/op
BenchmarkSelectTags/simple_markup/1000000_tags/50.3M-8   	       2	 782489786 ns/op	339752720 B/op	 4000046 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/complex_markup/10_tags/1020B-8       	  100000	     19911 ns/op	    11216 B/op	     152 allocs/op
BenchmarkSelectTags/complex_markup/100_tags/10.4K-8      	   10000	    182668 ns/op	    67280 B/op	    1415 allocs/op
BenchmarkSelectTags/complex_markup/1000_tags/108.8K-8    	    1000	   1857313 ns/op	   740560 B/op	   14020 allocs/op
BenchmarkSelectTags/complex_markup/10000_tags/1.1M-8     	      50	  20172929 ns/op	  9136360 B/op	  140030 allocs/op
BenchmarkSelectTags/complex_markup/100000_tags/11.6M-8   	       5	 216623197 ns/op	 90800371 B/op	 1400040 allocs/op
BenchmarkSelectTags/complex_markup/1000000_tags/120.6M-8 	       1	2270568036 ns/op	881045264 B/op	14000050 allocs/op
```

## [v1.1](https://github.com/thewizardplusplus/go-html-selector/tree/v1.1) (2020-07-18)

- friendly representation of filters:
  - for parsing from JSON;
  - for definition as a code literal;
- optimizations:
  - of searching for the right one among filters;
  - of conversion from a byte slice to a string.

### Benchmarks

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/simple_markup/10_tags/430B-8         	  200000	     11213 ns/op	     6736 B/op	      49 allocs/op
BenchmarkSelectTags/simple_markup/100_tags/4.4K-8        	   20000	     66364 ns/op	    25408 B/op	     412 allocs/op
BenchmarkSelectTags/simple_markup/1000_tags/45.7K-8      	    2000	    612808 ns/op	   190624 B/op	    4015 allocs/op
BenchmarkSelectTags/simple_markup/10000_tags/476.3K-8    	     200	   6808284 ns/op	  3448434 B/op	   40025 allocs/op
BenchmarkSelectTags/simple_markup/100000_tags/4.8M-8     	      20	  77850576 ns/op	 35420164 B/op	  400035 allocs/op
BenchmarkSelectTags/simple_markup/1000000_tags/50.3M-8   	       2	 839525444 ns/op	339752704 B/op	 4000045 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/complex_markup/10_tags/1020B-8       	  100000	     22544 ns/op	    11216 B/op	     151 allocs/op
BenchmarkSelectTags/complex_markup/100_tags/10.4K-8      	   10000	    184979 ns/op	    67280 B/op	    1414 allocs/op
BenchmarkSelectTags/complex_markup/1000_tags/108.8K-8    	    1000	   1883995 ns/op	   740560 B/op	   14019 allocs/op
BenchmarkSelectTags/complex_markup/10000_tags/1.1M-8     	     100	  20011726 ns/op	  9136368 B/op	  140029 allocs/op
BenchmarkSelectTags/complex_markup/100000_tags/11.6M-8   	       5	 219608064 ns/op	 90800355 B/op	 1400039 allocs/op
BenchmarkSelectTags/complex_markup/1000000_tags/120.6M-8 	       1	2294753289 ns/op	881045200 B/op	14000049 allocs/op
```

## [v1.0](https://github.com/thewizardplusplus/go-html-selector/tree/v1.0) (2020-07-16)

### Benchmarks

`htmlselector.SelectTags()` with a simple markup:

```
BenchmarkSelectTags/simple_markup/10_tags/430B-8         	  200000	     10419 ns/op	     6736 B/op	      49 allocs/op
BenchmarkSelectTags/simple_markup/100_tags/4.4K-8        	   20000	     62599 ns/op	    25408 B/op	     412 allocs/op
BenchmarkSelectTags/simple_markup/1000_tags/45.7K-8      	    2000	    583876 ns/op	   190624 B/op	    4015 allocs/op
BenchmarkSelectTags/simple_markup/10000_tags/476.3K-8    	     200	   7063473 ns/op	  3448441 B/op	   40025 allocs/op
BenchmarkSelectTags/simple_markup/100000_tags/4.8M-8     	      20	  77255459 ns/op	 35420161 B/op	  400035 allocs/op
BenchmarkSelectTags/simple_markup/1000000_tags/50.3M-8   	       2	 758725576 ns/op	339752720 B/op	 4000045 allocs/op
```

`htmlselector.SelectTags()` with a complex markup:

```
BenchmarkSelectTags/complex_markup/10_tags/1020B-8       	  100000	     19263 ns/op	    11216 B/op	     151 allocs/op
BenchmarkSelectTags/complex_markup/100_tags/10.4K-8      	   10000	    181784 ns/op	    67280 B/op	    1414 allocs/op
BenchmarkSelectTags/complex_markup/1000_tags/108.8K-8    	    1000	   1813507 ns/op	   740560 B/op	   14019 allocs/op
BenchmarkSelectTags/complex_markup/10000_tags/1.1M-8     	     100	  20045364 ns/op	  9136367 B/op	  140029 allocs/op
BenchmarkSelectTags/complex_markup/100000_tags/11.6M-8   	       5	 221623379 ns/op	 90800358 B/op	 1400039 allocs/op
BenchmarkSelectTags/complex_markup/1000000_tags/120.6M-8 	       1	2162959416 ns/op	881045280 B/op	14000049 allocs/op
```
