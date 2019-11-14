package bits

import "testing"

func TestBitString(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{21, "10101"},
		{3, "11"},
		{5, "101"},
		{12, "1100"},
		{101, "1100101"},
	}

	for _, test := range tests {
		actual := String(test.input)
		if actual != test.expected {
			t.Errorf("actual %s didnt match expected %s", actual, test.expected)
		}
	}
}
