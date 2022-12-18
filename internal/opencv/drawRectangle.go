package imageProcessing

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

type ImgRect struct {
	Rect  image.Rectangle
	Label string
}

func DrawRectanges(src image.Image, rects []ImgRect) (image.Image, error) {
	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// prepare image matrix
	img, err := gocv.ImageToMatRGBA(src)
	if err != nil {
		return nil, err
	}
	defer img.Close()

	for _, r := range rects {
		gocv.Rectangle(&img, r.Rect, blue, 1)

		size := gocv.GetTextSize(r.Label, gocv.FontHersheyPlain, 1.2, 2)
		pt := image.Pt(r.Rect.Min.X+(r.Rect.Min.X/10)-(size.X/2), r.Rect.Min.Y-2)
		gocv.PutText(&img, r.Label, pt, gocv.FontHersheyPlain, 1.2, blue, 2)
	}

	processedImage, err := img.ToImage()
	if err != nil {
		return nil, err
	}

	return processedImage, nil
}
