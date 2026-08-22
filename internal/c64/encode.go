package c64

func Encode(buf []uint8, w, h int) ([]byte, error) {
	err := validateSize(buf, w, h)
	if err != nil {
		return nil, err
	}

	cellsX := w >> 3
	cellsY := h >> 3
	output := []byte{}

	// Cycle over X, Y cells.
	for cY := range cellsY { // Cells (rows)
		for cX := range cellsX { // Cells (columns)

			for py := range 8 {
				var out byte // Output byte

				for px := range 8 {
					x := (cX<<3 + px)
					y := (cY<<3 + py)

					bpos := (y*w + x) << 2

					pixel := uint32(buf[bpos])
					pixel |= uint32(buf[bpos+1]) << 8
					pixel |= uint32(buf[bpos+2]) << 16
					pixel |= uint32(buf[bpos+3]) << 24

					if isPixelOn(pixel) {
						out |= 1 << (7 - px)
					}
				}
				output = append(output, out)
			}
		}
	}

	return output, nil
}
