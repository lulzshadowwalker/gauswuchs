package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"os"

	c "github.com/lulzshadowwalker/gauswuchs/internal/config"
	"github.com/lulzshadowwalker/gauswuchs/pkg/gauswuchs"
)

func main() {
	file, err := os.Open(c.GetSrc())
	if err != nil {
		log.Fatalf("failed to open image file %q", err)
	}
	defer file.Close()

	reader := io.Reader(file)
	m, f, err := image.Decode(reader)
	if f != "jpeg" && f != "png" {
		log.Fatalf("%q format is not supported you can only use .jpeg/.png files", f)
	}
	if err != nil {
		log.Fatalf("could not decode image %q", err)
	}

	result, err := os.Create(c.GetDest())
	if err != nil {
		log.Fatalf("failed to create file at destination %q", err)
	}
	defer result.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	out := gauswuchs.Blur(m, c.GetKernelSize(), float64(c.GetStrength()))
	err = png.Encode(result, out)
	if err != nil {
		log.Fatalf("failed to save image at destination %q", err)
	}

	fmt.Printf("image converted successfully: %q\n", c.GetDest())
}
