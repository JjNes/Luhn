package luhn

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		number   int
		expected bool
	}{
		{1230, true},
		{49927398716, true},
		{49927398717, false},
		{1234567812345678, false},
		{2201382000000005, true},
		{2201382000000104, true},
		{2201382000000013, true},
		{2204290100000006, true},
		{6011329933655299, true},
	}

	for _, test := range tests {
		if result := IsValid(test.number); result != test.expected {
			t.Errorf("Valid(%d) = %v; want %v", test.number, result, test.expected)
		}
	}
}

func TestCheckSumStr(t *testing.T) {
	tests := []struct {
		number   string
		expected bool
	}{
		{"1230", true},
		{"49927398716", true},
		{"49927398717", false},
		{"1234567812345678", false},
		{"123abc", false},
		{"1234567812345678", false},
		{"2201382000000005", true},
		{"2201382000000104", true},
		{"2201382000000013", true},
		{"2204290100000006", true},
		{"6011329933655299", true},
	}

	for _, test := range tests {
		if result := IsValidStr(test.number); result != test.expected {
			t.Errorf("checkSumStr(%s) = %v; want %v", test.number, result, test.expected)
		}
	}
}
