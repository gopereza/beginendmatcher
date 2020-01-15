# Begin-end matcher [![Build Status](https://travis-ci.org/gopereza/beginendmatcher.svg?branch=master)](https://travis-ci.org/gopereza/beginendmatcher)

Match string value by prefixes and suffixes

### Example
```go
package main

import "github.com/gopereza/beginendmatcher"

func main() {
	var matcher = beginendmatcher.NewBeginEndMatcher([]string{
		"a1",
		"b1",
		// ...
        "*abc",
        // ...
        "abc*",
        // ...
	})

	if matcher.Match("some") {
		// do
	}
}
```

### Testing
```text
BenchmarkPureMatcher_Match                 	  514190	      2156 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortMatcher_Match                 	 2483388	       477 ns/op	       0 B/op	       0 allocs/op
BenchmarkImmutableRadixTreeMatcher_Match   	 8477890	       140 ns/op	       0 B/op	       0 allocs/op
BenchmarkRadixTreeMatcher_Match            	 9127501	       131 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewImmutableRadixTreeMatcher      	     582	   2174686 ns/op	 2567483 B/op	   28346 allocs/op
BenchmarkNewRadixTreeMatcher               	    2310	    504617 ns/op	  175576 B/op	    3343 allocs/op
```