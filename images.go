package main
import tl "github.com/liconaj/brickcars/termloop"


type Image struct {
	*tl.Entity
}


func NewImage(canvas *tl.Canvas, x, y int) *Image {
	image := Image{
		Entity: tl.NewEntity(x, y, len(*canvas), len((*canvas)[0])),
	}
	image.ApplyCanvas(canvas)
	return &image
}
