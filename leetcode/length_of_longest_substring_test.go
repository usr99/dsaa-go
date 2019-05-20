package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, LengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, LengthOfLongestSubstring("bbbbb"), 1)
	assert.Equal(t, LengthOfLongestSubstring("pwwkew"), 3)
}
