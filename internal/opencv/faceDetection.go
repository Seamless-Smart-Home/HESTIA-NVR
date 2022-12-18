package imageProcessing

import (
	"errors"
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func DetectFaces(src image.Image) (image.Image, error) {
	// prepare image matrix
	img, err := gocv.ImageToMatRGBA(src)
	if err != nil {
		return nil, err
	}
	defer img.Close()

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("hair_classifier_face.xml") {
		return nil, fmt.Errorf("error reading cascade file: %v", "hair_classifier_face.xml")
	}

	// detect faces
	rects := classifier.DetectMultiScale(img)
	if len(rects) == 0 {
		return nil, errors.New("no faces detected")
	}

	facesDetected, err := img.ToImage()
	if err != nil {
		return nil, err
	}

	return facesDetected, nil
}
