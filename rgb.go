package dmc

import (
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

func (d *DmcColors) RgbToDmc(r float64, g float64, b float64) (string, string) {

	var dis float64
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

			// tc is Color struct holding the RGB values of each color in the
			// color bank to test how close it is to ic
			tc := colorful.Color{R: float64(red), G: float64(green), B: float64(blue)}

			if dis == 0 || (ic.DistanceLab(tc) < dis) {
				dis = ic.DistanceLab(tc)
				dmc = c.ColorName
				floss = c.Floss
			}
		}
	}

	return dmc, floss

}

// Convenience functions for convertion RGB colors to other color spaces
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
