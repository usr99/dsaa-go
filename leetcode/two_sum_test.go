package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, TwoSum([]int{7, 2, 11, 15}, 9), []int{0, 1})
}
