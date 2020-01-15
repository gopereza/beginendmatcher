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

	test(true, "abc", []string{"*bc", "abc1", "abc2"})
	test(true, "abc", []string{"*c", "abc1", "abc2"})
	test(true, "abc", []string{"a*", "abc1", "abc2"})
	test(true, "abc", []string{"ab*", "abc1", "abc2"})
}
