package main

import (
	"github.com/llgcode/draw2d/draw2dpdf"
)

const (
	a4width  = 210
	a4length = 297
)

// mesurments in mm
var (
	borderDistance float64 = 16.5
	lineGap        float64 = 8
	fileNm                 = "hello.pdf"
)

func main() {
	dest := draw2dpdf.NewPdf("P", "mm", "A4")
	gc := draw2dpdf.NewGraphicContext(dest)

	// Set some properties

	gc.SetLineWidth(0.5)

	gc.MoveTo(borderDistance, borderDistance)
	gc.LineTo(a4width-borderDistance, borderDistance)
	gc.LineTo(a4width-borderDistance, a4length-borderDistance)
	gc.LineTo(borderDistance, a4length-borderDistance)
	gc.LineTo(borderDistance, borderDistance)

	for loop := borderDistance + lineGap; loop < a4length-borderDistance; loop += lineGap {
		gc.MoveTo(borderDistance, loop)
		gc.LineTo(a4width-borderDistance, loop)
	}
	// Draw a closed shape
	//gc.MoveTo(1, 1) // should always be called first for a new path
	//gc.LineTo(90, 1)
	//gc.QuadCurveTo(100, 10, 10, 10)
	//gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dpdf.SaveToPdfFile(fileNm, dest)
}
