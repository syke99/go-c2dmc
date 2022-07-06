package lab

import (
	"strconv"

	"github.com/syke99/go-c2dmc/internal/pkg"

	"github.com/lucasb-eyer/go-colorful"
)

type Lab struct{}

func (d Lab) LabToDmc(cbm map[string]pkg.DefColor, l float64, a float64, b float64) (string, string) {

	var previousDistance float64
	var dmc string
	var floss string

	// ic is Color struct holding the Lab values passed in to LabToDmc
	ic := colorful.Lab(l, a, b)

	// Search for hex in hm to check for exact matches. If it exists, loop through
	// d.ColorBank for the color name and floss number
	hex := ic.Hex()

	if val, ok := cbm[hex]; ok {
		return val.ColorName, val.Floss
	}

	// Otherwise, use the hex value to create a new colorful.Color
	// and search through the cb for the closest match (in L*a*b colorspace)
	var i int
	for _, c := range cbm {

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
			i++
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

	return dmc, floss

}

// Convenience functions for conversion LAB colors to other color spaces
func (d Lab) LabToRgb(l float64, a float64, b float64) (float64, float64, float64) {
	c := colorful.Lab(l, a, b)

	_r := c.R
	_g := c.G
	_b := c.B

	return _r, _g, _b
}

func (d Lab) LabToHex(l float64, a float64, b float64) string {
	c := colorful.Lab(l, a, b)

	return c.Hex()
}

func (d Lab) LabToHsv(l float64, a float64, b float64) (float64, float64, float64) {
	c := colorful.Lab(l, a, b)

	return c.Hsv()
}
