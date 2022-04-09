package dmc

import (
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

func (d *DmcColors) LabToDmc(l float64, a float64, b float64) (string, string) {

	var previousDistance float64
	var dmc string
	var floss string

	// ic is Color struct holding the Lab values passed in to LabToDmc
	ic := colorful.Lab(l, a, b)

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
		for i, c := range d.ColorBank {

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

			if i == 0 {
				previousDistance = ic.DistanceLuv(tc)
				dmc = c.ColorName
				floss = c.Floss
				continue
			}

			curDis := ic.DistanceLuv(tc)

			if curDis == 0 {
				dmc = c.ColorName
				floss = c.Floss
				break
			}

			if curDis < previousDistance {
				previousDistance = ic.DistanceLuv(tc)
				dmc = c.ColorName
				floss = c.Floss
			}
		}
	}

	return dmc, floss

}

// Convenience functions for convertion LAB colors to other color spaces
func (d *DmcColors) LabToRgb(l float64, a float64, b float64) (float64, float64, float64) {
	c := colorful.Lab(l, a, b)

	_r := c.R
	_g := c.G
	_b := c.B

	return _r, _g, _b
}

func (d *DmcColors) LabToHex(l float64, a float64, b float64) string {
	c := colorful.Lab(l, a, b)

	return c.Hex()
}

func (d *DmcColors) LabToHsv(l float64, a float64, b float64) (float64, float64, float64) {
	c := colorful.Lab(l, a, b)

	return c.Hsv()
}
