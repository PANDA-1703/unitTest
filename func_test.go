package main

import "testing"

func TestSum(t *testing.T) {
	testTable := []struct {
		a        int
		b        int
		expected int
	}{
		{
			a:        2,
			b:        30,
			expected: 32,
		},
		{
			a:        3,
			b:        300,
			expected: 303,
		},
	}

	for _, testA := range testTable {
		result := Sum(testA.a, testA.b)
		if result != testA.expected {
			t.Errorf("Incorrect result. Expect %d, got %d",
				testA.expected, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	testTable := []struct {
		a        int
		expected int
	}{
		{
			a:        2,
			expected: 4,
		},
		{
			a:        3,
			expected: 9,
		},
	}

	for _, testA := range testTable {
		result := Multiply(testA.a)
		if result != testA.expected {
			t.Errorf("Incorrect result. Expect %d, got %d",
				testA.expected, result)
		}
	}
}
