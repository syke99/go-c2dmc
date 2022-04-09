package dmc

import (
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

func (d *DmcColors) HsvToDmc(h float64, s float64, v float64) (string, string) {

	var previousDistance float64
	var dmc string
	var floss string

	// ic is colorful.Color struct holding the HSV values passed in to RgbToDmc
	// so it can be converted to lab colorspace and tested
	ic := colorful.Hsv(h, s, v)

	// Search for hex in d.HexMap to check for exact matches. If it exists, loop through
	// d.ColorBank for the color name and floss number
	hex := ic.Hex()

	if val, ok := d.HexMap[hex]; ok {
		for _, c := range d.ColorBank {
			if c.ColorName == val {
				return c.ColorName, c.Floss
			}
		}
		// Otherwise, use the hex value to create a new colorful.Color
		// and search through the d.ColorBank for the closest match (in L*a*b colorspace)
	} else {
		for _, c := range d.ColorBank {

			red, _ := strconv.Atoi(c.R)
			green, _ := strconv.Atoi(c.G)
			blue, _ := strconv.Atoi(c.B)

			// tempc is a temporary colorful.Color struct holding the rgb values of
			// each color in the the colorBank so they can be converted to lab colorspace
			// values and tested for distance between colors
			tempc := colorful.Color{R: float64(red), G: float64(green), B: float64(blue)}

			l, a, b := tempc.Lab()

			// tc is colorful.Color struct holding the Lab values of each color in the
			// color bank to test how close it is to ic
			tc := colorful.Lab(l, a, b)

			if previousDistance == 0 || ic.DistanceCIEDE2000(tc) < previousDistance {
				dmc = c.ColorName
				floss = c.Floss
			}
		}
	}

	return dmc, floss

}

// Convenience functions for convertion HSV colors to other color spaces
func (d *DmcColors) HsvToRgb(h float64, s float64, v float64) (float64, float64, float64) {
	c := colorful.Hsv(h, s, v)

	_r := c.R
	_g := c.G
	_b := c.B

	return _r, _g, _b
}

func (d *DmcColors) HsvToHex(h float64, s float64, v float64) string {
	c := colorful.Hsv(h, s, v)

	return c.Hex()
}

func (d *DmcColors) HsvToLab(h float64, s float64, v float64) (float64, float64, float64) {
	c := colorful.Hsv(h, s, v)

	return c.Lab()
}
