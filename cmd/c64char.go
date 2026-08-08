package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
)

var (
	filename *string
)

func parseFlags() {
	filename = flag.String("file", "", "Input filename")

	flag.Parse()
}

func getImageFromFile(f *os.File) (i image.Image, err error) {
	i, err = png.Decode(f)

	return
}

func getImageSize(i image.Image) (w, h int) {
	bounds := i.Bounds()
	w = bounds.Dx()
	h = bounds.Dy()

	return
}

func main() {
	parseFlags()

	f, _ := os.Open(*filename)
	defer f.Close()

	i, err := getImageFromFile(f)
	h, w := getImageSize(i)



	fmt.Printf("Width %d, Height %d, err %v\n", w, h, err)
}
