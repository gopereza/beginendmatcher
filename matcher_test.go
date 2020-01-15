package beginendmatcher

import "testing"

func TestMatch(t *testing.T) {
	var got = Match("", []string{})
	var expect = false

	if expect != got {
		t.Errorf("expect %t, got %t", expect, got)
	}
}
