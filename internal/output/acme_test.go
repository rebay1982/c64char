package output

import (
	"strings"
	"testing"
)

func Test_AcmeFormatter_Output(t *testing.T) {
	testcases := []struct {
		name     string
		buf      []byte
		expected []string
	}{
		{
			name: "single_block",
			buf: []byte{
				0b00111100,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b00111100,
			},
			expected: []string{
				"; CHAR 0\n",
				"!byte %00111100\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %00111100\n",
				"\n",
			},
		},
		{
			name: "multi_block",
			buf: []byte{
				0b00111100,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b00111100,
				0b00111100,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b01000010,
				0b00111100,
			},
			expected: []string{
				"; CHAR 0\n",
				"!byte %00111100\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %00111100\n",
				"\n",
				"; CHAR 1\n",
				"!byte %00111100\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %01000010\n",
				"!byte %00111100\n",
				"\n",
			},
		},
	}

	f := acmeFormatter{}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			o := f.Output(tc.buf)

			expected := buildExpectedString(tc.expected)

			if o != expected {
				t.Fatalf("expected\n%s\n\n got\n%s\n\n", expected, o)
			}
		})
	}
}

func buildExpectedString(expected []string) string {
	b := strings.Builder{}
	for _, s := range expected {
		b.WriteString(s)
	}

	return b.String()
}
