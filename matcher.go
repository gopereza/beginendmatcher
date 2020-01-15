package beginendmatcher

import (
	"sort"
	"strings"
)

type (
	Matcher interface {
		Match(string) bool
	}

	PureMatcher struct {
		values []string
	}

	SortMatcher struct {
		values []string
	}
)

func NewPureMatcher(values []string) *PureMatcher {
	return &PureMatcher{values: values}
}

func (m *PureMatcher) Match(value string) bool {
	for _, prefix := range m.values {
		if strings.HasPrefix(value, prefix) {
			return true
		}
	}

	return false
}

func NewSortMatcher(values []string) *SortMatcher {
	sort.Strings(values)

	return &SortMatcher{values: values}
}

func (m *SortMatcher) Match(value string) bool {
	var length = len(value)
	if length == 0 {
		return false
	}

	var (
		current   = value[0:1]
		findIndex = sort.SearchStrings(m.values, current)
	)

	if findIndex == length {
		return false
	}

	if m.values[findIndex] == current {
		return true
	}

	for i := 1; i < length; i++ {
		var nextValues = m.values[findIndex:]

		current += string(value[i])

		findIndex = sort.SearchStrings(nextValues, current)

		if findIndex == length {
			return false
		}

		if m.values[findIndex] == current {
			return true
		}
	}

	return false
}
