package beginendmatcher

import "testing"

func TestBeginEndMatcher_Match(t *testing.T) {
	t.Helper()

	var test = func(expect bool, value string, prefixes []string) {
		t.Helper()

		var matcher = NewBeginEndMatcher(prefixes)
		var got = matcher.Match(value)

		if expect != got {
			t.Errorf("expect %t, got %t", expect, got)
		}
	}

	test(false, "", nil)
	test(false, "abc", nil)
	test(false, "abc", []string{"abc1", "abc2"})

	// * = 1+ chars
	test(false, "abc", []string{"*abc", "abc1", "abc2"})
	test(false, "abc", []string{"abc*", "abc1", "abc2"})

	test(true, "abc", []string{"*bc", "nnn1", "nnn2"})
	test(true, "abc", []string{"*c", "nnn1", "nnn2"})
	test(true, "abc", []string{"a*", "nnn1", "nnn2"})
	test(true, "abc", []string{"ab*", "nnn1", "nnn2"})
}

func BenchmarkBeginEndMatcher_Match(b *testing.B) {
	var matcher = NewBeginEndMatcher(begidendMatcherDataProvider)
	const expect = true

	b.ResetTimer()
	for i := 0; i < b.N/3; i++ {
		var got = matcher.Match(equalDataProvider[i%dataProviderLimit])

		if expect != got {
			b.Errorf("expect %t, got %t", expect, got)
		}
	}

	for i := 0; i < b.N/3; i++ {
		var got = matcher.Match(random + prefixDataProvider[i%dataProviderLimit])

		if expect != got {
			b.Errorf("expect %t, got %t", expect, got)
		}
	}

	for i := 0; i < b.N/3; i++ {
		var got = matcher.Match(suffixDataProvider[i%dataProviderLimit] + random)

		if expect != got {
			b.Errorf("expect %t, got %t", expect, got)
		}
	}
}
