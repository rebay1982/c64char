package img

import (
	"os"

	"image"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

func getImageForTest(fn string) (i image.Image, err error) {
	f, _ := os.Open(fn)
	defer f.Close()

	i, _, err = image.Decode(f)

	return
}

func TestImageToRGBA(t *testing.T) {
	testcases := []struct {
		name         string
		filename     string
		bufLen, h, w int
	}{
		{
			name:     "green_path_png",
			filename: "../../assets/test.png",
			bufLen:   524288, // 4* 256x512 (4 bytes per pixel)
			w:        256,
			h:        512,
		},
		{
			name:     "green_path_jpg",
			filename: "../../assets/test.png",
			bufLen:   524288, // 4* 256x512 (4 bytes per pixel)
			w:        256,
			h:        512,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			i, err := getImageForTest(tc.filename)
			if err != nil {
				t.Fatalf("Failed to load image sample for test: %s\ncause: %v", tc.filename, err)
			}

			b, w, h := ImageToRGBA(i)

			if w != tc.w || h != tc.h || len(b) != tc.bufLen {
				t.Errorf("expected %d, %d, %d got %d, %d, %d", tc.bufLen, tc.w, tc.h, len(b), w, h)
			}
		})
	}
}
