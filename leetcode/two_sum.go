package leetcode

// TwoSum return indices of the two numbers such that they add up
// to a specific target, when given an array of integers.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
// Example:
//	Given nums = [2, 7, 11, 15], target = 9,
//	Because nums[0] + nums[1] = 2 + 7 = 9,
//	return [0, 1].
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		if e, ok := m[val]; ok {
			// Keep ascending.
			return []int{e, i}
		}
		m[target-val] = i
	}

	// If not found, return null slice.
	return []int{}
}
