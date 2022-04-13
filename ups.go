package main

import (
	"image"

	"gocv.io/x/gocv"
)

var net gocv.Net

func main() {
	img := gocv.IMRead("./sample.jpg", 0)
	net := gocv.ReadNetFromONNX("./supres.onnx")

	// preprocessing the image
	gocv.Resize(img, &img, image.Pt(255, 255), 0, 0, 1)

	net.SetInput(img, "")
	outImg := net.Forward("")
	gocv.IMWrite("./output.png", outImg)
}
