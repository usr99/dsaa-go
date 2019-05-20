package leetcode

// LengthOfLongestSubstring return he length of the longest
// substring without repeating characters, when given a string.
// Example 1:
//   Input: "abcabcbb"
//   Output: 3
//   Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
//   Input: "bbbbb"
//   Output: 1
//   Explanation: The answer is "b", with the length of 1.
//
// Example 3:
//   Input: "pwwkew"
//   Output: 3
//   Explanation: The answer is "wke", with the length of 3.
//   Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func LengthOfLongestSubstring(s string) int {
	result := 0

	for i, _ := range s {
		str := getSubstring(s[i:])
		if result < len(str) {
			result = len(str)
		}
	}

	return result
}

func getSubstring(s string) string {
	e := make(map[rune]bool)

	for i, v := range s {
		if ok, _ := e[v]; ok {
			return s[:i]
		}

		e[v] = true
	}

	return s
}
