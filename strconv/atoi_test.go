package strconv

import "testing"

func TestAtoIHappyPath(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"-123", -123},
		{"0", 0},
		{"-0", 0},
		{"1", 1},
		{"-345345345", -345345345},
	}

	for _, test := range tests {
		actual, err := AtoI(test.input)
		if err != nil {
			t.Errorf("Error received for input: %s", test.input)
		} else if actual != test.expected {
			t.Errorf("Error invalid output for input: %s", test.input)
		}
	}
}

func TestAtoIUnHappyPath(t *testing.T) {
	tests := []string{
		"-fsdf3434",
		"--4323434",
		"dasdasd",
		"dasdasd234234",
		"=234234",
	}

	for _, input := range tests {
		_, err := AtoI(input)
		if err == nil {
			t.Errorf("Did not receive error for input: %s", input)
		}
	}
}
