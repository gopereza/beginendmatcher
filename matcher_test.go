package beginendmatcher

import (
	"strconv"
	"testing"
)

const (
	base = 36

	equalDataProviderShift = base * base * base * base * base
	dataProviderLimit      = 1 << 10

	prefixDataProviderShift = equalDataProviderShift + dataProviderLimit
	suffixDataProviderShift = prefixDataProviderShift + dataProviderLimit

	random = "x"
)

var (
	equalDataProvider  = generateDataProvider("", "", equalDataProviderShift, dataProviderLimit)
	prefixDataProvider = generateDataProvider("", "", prefixDataProviderShift, dataProviderLimit)
	suffixDataProvider = generateDataProvider("", "", suffixDataProviderShift, dataProviderLimit)

	prefixMatcherDataProvider   = generateDataProvider("*", "", prefixDataProviderShift, dataProviderLimit)
	suffixMatcherDataProvider   = generateDataProvider("", "*", suffixDataProviderShift, dataProviderLimit)
	begidendMatcherDataProvider = append(
		append(
			append(
				make([]string, 0, 3*dataProviderLimit),
				equalDataProvider...,
			),
			prefixMatcherDataProvider...,
		),
		suffixMatcherDataProvider...,
	)
)

func TestPureMatcher(t *testing.T) {
	testMatcher(t, func(values []string) Matcher {
		return NewPureMatcher(values)
	})
}

func TestSortMatcher(t *testing.T) {
	testMatcher(t, func(values []string) Matcher {
		return NewSortMatcher(values)
	})
}

func TestRadixTreeMatcher(t *testing.T) {
	testMatcher(t, func(values []string) Matcher {
		return NewRadixTreeMatcher(values)
	})
}

func testMatcher(t *testing.T, newMatcher func([]string) Matcher) {
	t.Helper()

	var test = func(expect bool, value string, prefixes []string) {
		t.Helper()

		var matcher = newMatcher(prefixes)
		var got = matcher.Match(value)

		if expect != got {
			t.Errorf("expect %t, got %t", expect, got)
		}
	}

	test(false, "", nil)
	test(false, "any", nil)
	test(false, "b", []string{"a"})
	test(false, "b", []string{"a", "c", "d", "e", "f"})
	test(false, "abc", []string{"abcd"})
	test(false, "ABCD", []string{"abc"})
	test(true, "abc", []string{"ab"})
	test(true, "abc", []string{"abc"})
	test(true, "bcd", []string{"ab", "bcd"})
}

func BenchmarkPureMatcher_Match(b *testing.B) {
	var matcher = NewPureMatcher(equalDataProvider)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match(equalDataProvider[i%dataProviderLimit] + random)
	}
}

func BenchmarkSortMatcher_Match(b *testing.B) {
	var matcher = NewSortMatcher(equalDataProvider)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match(equalDataProvider[i%dataProviderLimit] + random)
	}
}

func BenchmarkRadixTreeMatcher_Match(b *testing.B) {
	var matcher = NewRadixTreeMatcher(equalDataProvider)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match(equalDataProvider[i%dataProviderLimit] + random)
	}
}

func BenchmarkNewRadixTreeMatcher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewRadixTreeMatcher(equalDataProvider)
	}
}

func generateDataProvider(prefix, suffix string, shift, limit int) []string {
	var (
		to = shift + limit
	)

	var result = make([]string, 0, limit)

	for i := shift; i < to; i++ {
		result = append(result, prefix+strconv.FormatUint(uint64(i), base)+suffix)
	}

	return result
}
