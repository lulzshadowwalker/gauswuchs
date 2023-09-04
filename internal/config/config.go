package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var (
	src        string
	dest       string
	kernelSize int
	strength   int
)

var workingDir string

func init() {
	flag.StringVar(&src, "src", "", "source image file to be converted")
	flag.StringVar(&dest, "dest", "", "output destination directory")
	flag.IntVar(&kernelSize, "kernelSize", 5, "has to be an odd number")
	flag.IntVar(&strength, "strength", 10, "uhm strength")

	flag.Parse()

	if src == "" {
		wd, err := getWorkingDir()
		if err != nil {
			log.Fatal(err.Error())
		}

		src = path.Join(wd, "../../assets/images/hamster.jpeg")
	}

	if dest == "" {
		d, err := getWorkingDir()
		if err != nil {
			log.Fatal(err.Error())
		}

		dest = d
	}
	dest = path.Join(dest, "gauswuchs.png")

	if kernelSize%2 == 0 {
		log.Fatal("kernel size has to be an odd number")
	}
}

func getWorkingDir() (string, error) {
	if workingDir != "" {
		return workingDir, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf(
			`failed to get the current working dir
try assigning "--src" and/or "--dest" explicitly
error: %w`, err)
	}

	workingDir = wd
	return wd, nil
}

func GetSrc() string {
	return src
}

func GetDest() string {
	return dest
}

func GetKernelSize() int {
	return kernelSize
}

func GetStrength() int {
	return strength
}
