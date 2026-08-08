package img

import (
	"image"
	"image/draw"
)

func ImageToRGBA(src image.Image) ([]uint8, int, int) {
	srcBounds := src.Bounds()
	srcW := srcBounds.Dx()
	srcH := srcBounds.Dy()

	dst := image.NewNRGBA(image.Rect(0, 0, srcW, srcH))
	draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Src)

	return dst.Pix, dst.Bounds().Dx(), dst.Bounds().Dy()
}
