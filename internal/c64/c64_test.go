package c64

import (
	"testing"
)

func Test_IsValidSize(t *testing.T) {
	testCases := []struct {
		name        string
		buf         []uint8
		w, h        int
		expectError bool
	}{
		{
			name:        "320x200",
			buf:         make([]uint8, 320*200*4),
			w:           320,
			h:           200,
			expectError: false,
		},
		{
			name:        "10x10",
			buf:         make([]uint8, 10*10*4),
			w:           10,
			h:           10,
			expectError: true,
		},
		{
			name:        "10x10",
			buf:         make([]uint8, 8*8*4),
			w:           10,
			h:           10,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateSize(tc.buf, tc.w, tc.h)
			if tc.expectError && err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !tc.expectError && err != nil {
				t.Fatalf("did not expect an error, got an error")
			}
		})
	}
}

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
