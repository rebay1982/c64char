package c64

import (
	"fmt"
)

func validateSize(buf []uint8, w, h int) error {
	l := len(buf)

	// Buffer length must match width and height.
	if l != (w*h)<<2 {
		return fmt.Errorf("invalid image buffer size %d, has to match w*h*4 %d", l, (w*h)<<2)
	}

	// Buffer length, width and height HAVE to be multiples of 8.
	if l&0x7 > 0 {
		return fmt.Errorf("invalid image buffer size %d, has to be multiple of 8", l)
	}

	if w&0x7 > 0 {
		return fmt.Errorf("invalid image width %d, has to be multiple of 8", w)
	}

	if h&0x7 > 0 {
		return fmt.Errorf("invalid image height %d, has to be multiple of 8", h)
	}

	return nil
}

// isPixelOn validates that the pixel (32bit RGBA) is non-black. This function ignores the ALPHA channel.
func isPixelOn(p uint32) bool {
	// Pixels need to be packed as following:
	//   p := dst.Pix(0) << 24 | dst.Pix(1) << 16 | dst.Pix(2) << 8 | dst.Pix(3).

	// Assuming 0xRRGGBBAA is packed this way.
	return p&0xFFFFFF00 > 0
}
