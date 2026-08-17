package c64

func Encode(buf []uint8, w, h int) ([]byte, error) {
	err := validateSize(buf, w, h)
	if err != nil {
		return nil, err
	}

	cellsX := w >> 3
	cellsY := h >> 3

	// Cycle over X, Y cells.
	for cY := 0; cY <= cellsY; cY++ {
		for cX := 0; cX <= cellsX; cX++ {

			for py := range 8 {
				var out byte // Output byte

				for px := range 8 {
					x := (cellsX << 3 + px) << 2
					y := ((cellsY << 3 + py) << 2) * w

					var pixel uint32 = uint32(buf[x + y])
					pixel |= uint32(buf[x + y + 1]) << 8
					pixel |= uint32(buf[x + y + 2]) << 16
					pixel |= uint32(buf[x + y + 3]) << 24

					if isPixelLit(pixel) {
						out |= 1 << (7 - px)
					}
				}
			}
		}
	}

	return nil, nil
}
