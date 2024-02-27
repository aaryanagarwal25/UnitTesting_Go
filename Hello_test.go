package Hello

import (
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Case1",
			a:        4,
			b:        5,
			expected: 9,
		},
		{
			name:     "Case2",
			a:        -4,
			b:        5,
			expected: 1,
		},
		{
			name:     "Case1",
			a:        -4,
			b:        -5,
			expected: -9,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if tc.expected != result {
				t.Fatalf("expected %v", tc.expected)
			}
		})
	}
}
