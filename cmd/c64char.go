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
	f, _ := os.Open(c.Filename)
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
