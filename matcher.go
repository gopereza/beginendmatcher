package beginendmatcher

import (
	"github.com/armon/go-radix"
	"github.com/hashicorp/go-immutable-radix"
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

	ImmutableRadixTreeMatcher struct {
		tree *iradix.Tree
	}

	RadixTreeMatcher struct {
		tree *radix.Tree
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

	var findIndex = 0

	for i := 1; i <= length; i++ {
		var (
			values  = m.values[findIndex:]
			current = value[0:i]
		)

		findIndex = sort.SearchStrings(values, current)

		if findIndex == length {
			return false
		}

		if m.values[findIndex] == current {
			return true
		}
	}

	return false
}

func NewImmutableRadixTreeMatcher(values []string) *ImmutableRadixTreeMatcher {
	var tree = iradix.New()

	for _, value := range values {
		tree, _, _ = tree.Insert([]byte(value), 0)
	}

	return &ImmutableRadixTreeMatcher{
		tree: tree,
	}
}

func (m *ImmutableRadixTreeMatcher) Match(value string) bool {
	_, _, exists := m.tree.Root().LongestPrefix([]byte(value))

	return exists
}

func NewRadixTreeMatcher(values []string) *RadixTreeMatcher {
	var tree = radix.New()

	for _, value := range values {
		tree.Insert(value, 0)
	}

	return &RadixTreeMatcher{
		tree: tree,
	}
}

func (m *RadixTreeMatcher) Match(value string) bool {
	_, _, exists := m.tree.LongestPrefix(value)

	return exists
}
