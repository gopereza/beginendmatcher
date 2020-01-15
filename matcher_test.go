package beginendmatcher

import (
	"strconv"
	"testing"
)

const (
	base = 36

	dataProviderShift = base * base
	dataProviderLimit = 1 << 10
)

var (
	dataProvider = generateDataProvider(dataProviderShift, dataProviderLimit)
)

func TestDirectMatch(t *testing.T) {
	var test = func(expect bool, value string, prefixes []string) {
		var matcher = NewMatcher(prefixes)
		var got = matcher.Match(value)

		if expect != got {
			t.Errorf("expect %t, got %t", expect, got)
		}
	}

	test(false, "", nil)
	test(false, "b", []string{"a"})
	test(false, "abc", []string{"abcd"})
	test(true, "abc", []string{"ab"})
}

func BenchmarkMatcher_Match(b *testing.B) {
	var matcher = NewMatcher(dataProvider)

	for i := 0; i < b.N; i++ {
		matcher.Match(dataProvider[i%dataProviderLimit])
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
