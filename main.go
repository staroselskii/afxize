package main

import (
	"flag"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/staroselskii/afxize/personify"
	"github.com/staroselskii/afxize/facefinder"
)

var haarCascade = flag.String("haar", "haarcascade_frontalface_alt.xml", "The location of the Haar Cascade XML configuration to be provided to OpenCV.")
var facesDir = flag.String("faces", "faces", "The directory to search for faces.")

func main() {
	flag.Parse()

	var personFaces facefinder.FaceList

	var facesPath string
	var err error

	if *facesDir != "" {
		facesPath, err = filepath.Abs(*facesDir)
		if err != nil {
			panic(err)
		}
	}

	err = personFaces.Load(facesPath)
	if err != nil {
		panic(err)
	}
	if len(personFaces) == 0 {
		panic("no faces found")
	}

	file := flag.Arg(0)

	baseImage := facefinder.LoadImage(file)

	var canvas = personify.Personify(baseImage, haarCascade, personFaces)

	jpeg.Encode(os.Stdout, canvas, &jpeg.Options{jpeg.DefaultQuality})
}
