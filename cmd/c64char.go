package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/rebay1982/c64char/internal/config"
	"github.com/rebay1982/c64char/internal/img"
)

func parseFlags() config.Config {
	cfg := config.Config{}

	f := flag.String("file", "", "Input filename.")
	e := flag.Bool("e", true, "Specify to encode a PNG to C64 data format.")
	d := flag.Bool("d", false, "Specify to decode C64 data format to PNG.")

	flag.Parse()

	cfg.Filename = *f
	cfg.Encode = *e

	// Decode explicitly specified takes precedense on the default encoding behaviour.
	if *d {
		cfg.Encode = !*d
	}

	return cfg
}

func getImageFromFile(f *os.File) (i image.Image, err error) {
	i, err = png.Decode(f)

	return
}

func main() {
	cfg := parseFlags()

	f, _ := os.Open(cfg.Filename)
	defer f.Close()

	i, err := getImageFromFile(f)
	if err != nil {
		fmt.Printf("unable to load image from file. %v\n", err)
		os.Exit(1)
	}

	buff, w, h := img.ImageToRGBA(i)

	fmt.Printf("output: %d, %d, %d", len(buff), w, h)
}
