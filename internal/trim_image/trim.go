package trim_image

import (
	"log"

	imagick "gopkg.in/gographics/imagick.v3/imagick"
)

func TrimImage(inputFilename string) string {
	outputFilename := inputFilename
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	if err := mw.ReadImage(inputFilename); err != nil {
		log.Fatal(err)
	}

	mw.SetImageFuzz(0)

	if err := mw.TrimImage(0.0); err != nil {
		log.Fatal(err)
	}

	if err := mw.SetImagePage(0, 0, 0, 0); err != nil {
		log.Fatal(err)
	}

	if err := mw.WriteImage(outputFilename); err != nil {
		log.Fatal(err)
	}

	return outputFilename
}
