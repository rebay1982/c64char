package c64

import (
	"fmt"
)

func validateSize(buf []uint8, w, h int) error {
	l := len(buf)
	
	// Buffer length, width and height HAVE to be multiples of 8.
	if l & 0x7 > 0 {
		return fmt.Errorf("invalid image buffer size %d, has to be multiple of 8.", l)
	}

	if w & 0x7 > 0 {
		return fmt.Errorf("invalid image width %d, has to be multiple of 8.", w)
	}

	if h & 0x7 > 0 {
		return fmt.Errorf("invalid image height %d, has to be multiple of 8.", w)
	}

	return nil
}

func isPixelLit(p uint32) bool {
	// Pixels need to be packed as following:
	//   p := dst.Pix(0) << 24 | dst.Pix(1) << 16 | dst.Pix(2) << 8 | dst.Pix(3).

	// Assuming RGBA is packed this way.
	if p & 0xFFFFFF00 > 0 {
		return true
	}

	return false
}
