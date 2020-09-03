package coding

import "strconv"

/*

A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

*/

var resultMap map[int]int

func NumDecoding(s string) int {
	resultMap = make(map[int]int)
	return numDecoding(s, 0, len(s))
}

func numDecoding(s string, start, end int) int {
	result, ok := resultMap[start]
	if ok {
		return result
	}
	if len(s) <= 1 || start >= end-1 {
		return 1
	}
	num, _ := strconv.ParseInt(s[start:start+2], 10, 64)
	if num > 26 {
		result = numDecoding(s, start+1, end)
	} else {
		result = numDecoding(s, start+1, end) + numDecoding(s, start+2, end)
	}
	resultMap[start] = result
	return result
}
