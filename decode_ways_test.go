package coding

import "testing"

func TestNumDecoding(t *testing.T) {
	testData := map[string]int{
		"12":   2,
		"226":  3,
		"1":    1,
		"":     1,
		"2312": 4,
	}
	for str, num := range testData {
		resultNum := NumDecoding(str)
		if resultNum != num {
			t.Errorf("NumDecoding(%s) should be equal to %d, but is %d", str, num, resultNum)
		}
	}
}
