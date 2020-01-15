package beginendmatcher

import (
	"strconv"
	"testing"
)

const (
	base = 36

	dataProviderShift = base * base
	dataProviderLimit = 1 << 10

	randomSuffix = "x"
)

var (
	dataProvider = generateDataProvider(dataProviderShift, dataProviderLimit)
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

func TestImmutableRadixTreeMatcher(t *testing.T) {
	testMatcher(t, func(values []string) Matcher {
		return NewImmutableRadixTreeMatcher(values)
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
	var matcher = NewPureMatcher(dataProvider)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match(dataProvider[i%dataProviderLimit] + randomSuffix)
	}
}

func BenchmarkSortMatcher_Match(b *testing.B) {
	var matcher = NewSortMatcher(dataProvider)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match(dataProvider[i%dataProviderLimit] + randomSuffix)
	}
}

func BenchmarkImmutableRadixTreeMatcher_Match(b *testing.B) {
	var matcher = NewImmutableRadixTreeMatcher(dataProvider)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match(dataProvider[i%dataProviderLimit] + randomSuffix)
	}
}

func BenchmarkRadixTreeMatcher_Match(b *testing.B) {
	var matcher = NewRadixTreeMatcher(dataProvider)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.Match(dataProvider[i%dataProviderLimit] + randomSuffix)
	}
}

func BenchmarkNewImmutableRadixTreeMatcher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewImmutableRadixTreeMatcher(dataProvider)
	}
}

func BenchmarkNewRadixTreeMatcher(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewRadixTreeMatcher(dataProvider)
	}
}

func generateDataProvider(shift, limit int) []string {
	var (
		to = shift + limit
	)

	var result = make([]string, 0, limit)

	for i := shift; i < to; i++ {
		result = append(result, strconv.FormatUint(uint64(i), base))
	}

	return result
}
