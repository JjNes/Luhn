package luhn

import "testing"

func TestIsValidInt(t *testing.T) {
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
		{18, true},
		{6011329933655299, true},
	}

	for _, test := range tests {
		if result := IsValidInt(test.number); result != test.expected {
			t.Errorf("IsValidInt(%d) = %v; want %v", test.number, result, test.expected)
		}
	}
}

func TestIsValidStr(t *testing.T) {
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
		{"", false},
		{"6011329933655299", true},
	}

	for _, test := range tests {
		if result := IsValidStr(test.number); result != test.expected {
			t.Errorf("IsValidStr(%q) = %v; want %v", test.number, result, test.expected)
		}
	}
}

func TestNumberGenerator(t *testing.T) {
	tests := []struct {
		number int
	}{
		{123},
		{499273987},
		{4992739871},
		{123456781234567},
		{123456781234567},
		{220138200000000},
		{220138200000010},
		{220138200000001},
		{220429010000000},
		{601132993365529},
	}

	for _, test := range tests {
		full := GenerateLuhnNumber(test.number)
		if !IsValidInt(full) {
			t.Errorf("GenerateLuhnNumber(%d) produced invalid number %d", test.number, full)
		}
	}
}

func TestCheckDigitMath(t *testing.T) {
	body := 7992739871
	cd := luhnCheckDigit(body)
	if cd != 3 {
		t.Errorf("luhnCheckDigit(%d) = %d; want 3", body, cd)
	}
	full := GenerateLuhnNumber(body)
	if !IsValidInt(full) {
		t.Fatalf("full %d expected valid", full)
	}
}
