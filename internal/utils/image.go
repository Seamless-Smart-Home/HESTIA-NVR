package utils

import (
	"HESTIA/internal/compreface"
	imageProcessing "HESTIA/internal/opencv"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"
)

func FaceDetectionImgProcessing(snapshot image.Image) error {
	// Run Face Detection Pre-Processing
	img, err := imageProcessing.DetectFaces(snapshot)
	if err != nil {
		return err
	}

	// Send to Compreface for recognition
	results, err := compreface.Client.Recognition.RecognizeFaces(img)
	if err != nil {
		return err
	}

	// Draw rectangle around detected faces
	var rects []imageProcessing.ImgRect
	for _, result := range results.Result {
		rect := imageProcessing.ImgRect{
			Rect: image.Rectangle{
				Min: image.Point{
					X: result.Box.XMin,
					Y: result.Box.YMin,
				},
				Max: image.Point{
					X: result.Box.XMax,
					Y: result.Box.YMax,
				},
			},
			Label: result.Subjects[0].Subject,
		}

		rects = append(rects, rect)
	}

	detectedImg, err := imageProcessing.DrawRectanges(img, rects)
	if err != nil {
		return err
	}

	// Save Image
	err = SaveImage(detectedImg, "storage")
	if err != nil {
		return err
	}

	return nil
}

func SaveImage(image image.Image, path string) error {
	f, err := os.Create(filepath.Join(path, time.Now().String()+".jpg"))
	if err != nil {
		return err
	}
	defer f.Close()

	if err = jpeg.Encode(f, image, nil); err != nil {
		return fmt.Errorf("failed to encode: %v", err)
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}
