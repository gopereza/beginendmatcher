package beginendmatcher

import "github.com/armon/go-radix"

type BeginEndMatcher struct {
	fullMatchMap    map[string]bool
	prefixRadixTree *radix.Tree
	suffixRadixTree *radix.Tree
}

func NewBeginEndMatcher(values []string) *BeginEndMatcher {
	var (
		fullMatchMap    = make(map[string]bool)
		prefixRadixTree = radix.New()
		suffixRadixTree = radix.New()
	)

	for _, value := range values {
		var length = len(value)

		if length == 0 {
			continue
		}

		switch {
		case value[0] == '*':
			suffixRadixTree.Insert(reverseAsciiString(value[1:]), 0)
		case value[length-1] == '*':
			prefixRadixTree.Insert(value[:length-1], 0)
		default:
			fullMatchMap[value] = true
		}
	}

	return &BeginEndMatcher{
		fullMatchMap:    fullMatchMap,
		prefixRadixTree: prefixRadixTree,
		suffixRadixTree: suffixRadixTree,
	}
}

func (m *BeginEndMatcher) Match(value string) bool {
	if m.fullMatchMap[value] {
		return true
	}

	var length = len(value)

	prefix, _, prefixExists := m.prefixRadixTree.LongestPrefix(value)
	if prefixExists && length > len(prefix) {
		return true
	}

	suffix, _, suffixExists := m.suffixRadixTree.LongestPrefix(reverseAsciiString(value))
	if suffixExists && length > len(suffix) {
		return true
	}

	return false
}

func reverseAsciiString(s string) string {
	var (
		length = len(s)
		result = make([]byte, 0, length)
	)

	for i := length - 1; i >= 0; i-- {
		result = append(result, s[i])
	}

	return string(result)
}
