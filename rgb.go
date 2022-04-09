package dmc

import (
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

func (d *DmcColors) RgbToDmc(r float64, g float64, b float64) (string, string) {

	var previousDistance float64
	var dmc string
	var floss string

	// ic is Color struct holding the RGB values passed in to RgbToDmc
	ic := colorful.Color{R: r, G: g, B: b}

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

			if previousDistance == 0 || ic.DistanceLuv(tc) < previousDistance {
				previousDistance = ic.DistanceLuv(tc)
				dmc = c.ColorName
				floss = c.Floss
			}
		}
	}

	return dmc, floss

}

func (d *DmcColors) RgbToLab(r float64, g float64, b float64) (float64, float64, float64) {
	c := colorful.Color{R: r, G: g, B: b}

	return c.Lab()
}

func (d *DmcColors) RgbToHex(r float64, g float64, b float64) string {
	c := colorful.Color{R: r, G: g, B: b}

	return c.Hex()
}

func (d *DmcColors) RgbToHsv(r float64, g float64, b float64) (float64, float64, float64) {
	c := colorful.Color{R: r, G: g, B: b}

	return c.Hsv()
}
