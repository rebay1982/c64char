package c64

import (
	"testing"
)

func Test_IsPixelOn(t *testing.T) {
	testCases := []struct {
		name     string
		pixel    uint32
		expected bool
	}{
		{
			name:     "red",
			pixel:    0xFF000000,
			expected: true,
		},
		{
			name:     "green",
			pixel:    0x00FF0000,
			expected: true,
		},
		{
			name:     "blue",
			pixel:    0x0000FF00,
			expected: true,
		},
		{
			name:     "black",
			pixel:    0x00000000,
			expected: false,
		},
		{
			name:     "black_alpha",
			pixel:    0x000000FF,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPixelOn(tc.pixel)

			if got != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, got)
			}
		})
	}
}
