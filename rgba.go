package dmc

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

func (d *DmcColors) RgbaToRgb(col color.Color) (float64, float64, float64) {

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
	r := c.R
	g := c.G
	b := c.B

	return r, g, b
}
