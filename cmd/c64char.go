package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/rebay1982/c64char/internal/c64"
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
	i, _, err = image.Decode(f)

	return
}

func main() {
	cfg := parseFlags()

	if cfg.Encode {
		f, _ := os.Open(cfg.Filename)
		defer f.Close()

		i, err := getImageFromFile(f)
		if err != nil {
			fmt.Printf("unable to load image from file. %v\n", err)
			os.Exit(1)
		}

		buf, w, h := img.ImageToRGBA(i)
		data, err := c64.Encode(buf, w, h)

		if err != nil {
			fmt.Printf("unable to encode image from file. %v\n", err)
			os.Exit(1)
		}

		nbBlocks := len(data) >> 3
		for i := range nbBlocks {
			fmt.Printf("; block %d\n", i)

			for j := range 8 {
				b := data[(i << 3) + j] 

			//	fmt.Printf("!byte %%b\n", b)
			}
		}

	} else {
		fmt.Println("Decoding is not implemented yet, exiting...")
	}
}
