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
	"github.com/rebay1982/c64char/internal/output"
)

func parseFlags() config.Config {
	cfg := config.Config{}

	f := flag.String("f", "", "Image filename.")
	v := flag.Bool("v", false, "Show version and exit.")

	flag.Parse()

	// Default until the decode function is  implemented.
	cfg.Encode = true
	cfg.Filename = *f
	cfg.ShowVersion = *v

	return cfg
}

func getImageFromFile(f *os.File) (i image.Image, err error) {
	i, _, err = image.Decode(f)

	return
}

func main() {
	cfg := parseFlags()

	if cfg.ShowVersion {
		// Show version and GTFO.
	}

	if cfg.Encode {
		if o, err := Encode(cfg); err != nil {
			os.Exit(1)

		} else {
			fmt.Print(o)
		}

	} else {
		fmt.Println("Decoding is not implemented yet, exiting...")
	}
}

func Encode(c config.Config) (string, error) {
	f, err := os.Open(c.Filename)
	if err != nil {
		fmt.Printf("unable to open specified image file %s. %v\n", c.Filename, err)
		return "", err
	}
	defer f.Close()

	i, err := getImageFromFile(f)
	if err != nil {
		fmt.Printf("unable to load image from file. %v\n", err)
		return "", err
	}

	data, err := c64.Encode(img.ImageToRGBA(i))
	if err != nil {
		fmt.Printf("unable to encode image from file. %v\n", err)
		return "", err
	}

	formatter := output.NewFormatter(output.AcmeFormatter)
	o := formatter.Output(data)

	return o, nil
}
