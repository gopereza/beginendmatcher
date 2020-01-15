package beginendmatcher

import "strings"

type (
	Matcher struct {
		values []string
	}
)

func NewMatcher(values []string) *Matcher {
	return &Matcher{values: values}
}

func (m *Matcher) Match(value string) bool {
	for _, prefix := range m.values {
		if strings.HasPrefix(value, prefix) {
			return true
		}
	}

	return false
}
