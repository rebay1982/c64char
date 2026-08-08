package main

import (
	"os"

	"image"
	"image/png"
	"testing"
)

func getImageForTest(fn string) (i image.Image, err error) {
	f, _ := os.Open(fn)
	defer f.Close()

	i, err = png.Decode(f)

	return
}

func Test_getImageFileSize(t *testing.T) {
	testcases := []struct {
		name     string
		filename string
		h, w     int
	}{
		{
			name:     "green_path",
			filename: "../assets/sample.png",
			w:        256,
			h:        512,
		},
	}

	for _, tc := range testcases {
		i, err := getImageForTest(tc.filename)

		if err != nil {
			t.Errorf("Failed to load image sample for test: %s\ncause: %v", tc.filename, err)
		}

		w, h := getImageSize(i)

		if w != tc.w || h != tc.h || err != nil {
			t.Errorf("expected %d, %d got %d, %d", tc.w, tc.h, w, h)
		}
	}
}
