# Change Log

## [v1.4](https://github.com/thewizardplusplus/go-html-selector/tree/v1.4) (2020-07-30)

- skipping empty attributes (i.e. with an empty value; optional);
- representing a result:
  - built-in builders:
    - with collecting only values of HTML attributes;
- optimizations:
  - optimize the `htmlselector.SelectTags()` function;
  - optimize the `byteutils.Copy()` function;
  - optimize the `builders.StructuralBuilder.AddAttribute()` method.

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
