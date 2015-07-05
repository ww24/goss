/**
 * Go Screenshot Package
 */

package goss

import (
	"image"

	"github.com/ww24/goss/screenshot"
)

// Screenshot structure
type Screenshot struct {
	image image.Image
}

// Goss API
type Goss interface {
	ToImage() image.Image
	ToRGBA() *image.RGBA
}

// Capture method
func Capture() (img image.Image, err error) {
	img, err = screenshot.Capture()
	return
}

// New constructor
func New() (ss *Screenshot, err error) {
	ss = &Screenshot{}
	ss.image, err = screenshot.Capture()
	return
}

// ToImage method
func (ss *Screenshot) ToImage() (img image.Image) {
	img = ss.image
	return
}

// ToRGBA method
func (ss *Screenshot) ToRGBA() (rgba *image.RGBA) {
	rect := ss.image.Bounds()
	rgba = image.NewRGBA(rect)
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			col := ss.image.At(x, y)
			rgba.Set(x, y, col)
		}
	}
	return
}
