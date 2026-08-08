package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/rebay1982/c64char/internal/config"
)

func parseFlags() config.Config {
	cfg := config.Config{}

	f := flag.String("file", "", "Input filename")
	e := flag.Bool("e", false, "Specify to encode a PNG to C64 data format")

	flag.Parse()

	cfg.Filename = *f
	cfg.Encode = *e

	return cfg
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
	cfg := parseFlags()

	f, _ := os.Open(cfg.Filename)
	defer f.Close()

	i, err := getImageFromFile(f)
	h, w := getImageSize(i)

	fmt.Printf("Width %d, Height %d, err %v\n", w, h, err)
}
