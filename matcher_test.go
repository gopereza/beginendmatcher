package beginendmatcher

import "testing"

func TestMatch(t *testing.T) {
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
