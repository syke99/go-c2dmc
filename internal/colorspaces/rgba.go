package colorspaces

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

func RgbaToRgb(col color.Color) (float64, float64, float64) {

	// use the go-colorful package to
	// return a colorful.Color that contains
	// the corresponding RGB values for the
	// color.Color (col) that is passed in
	c, s := colorful.MakeColor(col)
	if !s {
		return 0, 0, 0
	}

	// retrieve the RGB values from c and
	// return them
	return c.Clamped().R, c.Clamped().G, c.Clamped().B
}
