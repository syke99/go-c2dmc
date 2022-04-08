package dmc

import (
	"log"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

func (d *DmcColors) HexToDmc(hex string) (string, string) {

	var dmc string
	var floss string

	// Search for hex in d.HexMap to check for exact matches. If it exists, loop through
	// d.ColorBank for the color name and floss number
	if val, ok := d.HexMap[hex]; ok {
		for _, c := range d.ColorBank {
			if c.ColorName == val {
				return c.ColorName, c.Floss
			}
		}
		// Otherwise, use the hex value to create a new colorful.Color
		// and search through the d.ColorBank for the closest match (in L*a*b colorspace)
	} else {
		hc, err := colorful.Hex(hex)
		if err != nil {
			log.Fatal(err)
		}

		// ic is Color struct holding the RGB values passed in to RgbToDmc
		ic := colorful.Color{R: hc.R, G: hc.G, B: hc.B}

		for _, c := range d.ColorBank {

			red, _ := strconv.Atoi(c.R)
			green, _ := strconv.Atoi(c.G)
			blue, _ := strconv.Atoi(c.B)

			// tc is Color struct holding the RGB values of each color in the
			// color bank to test how close it is to ic
			tc := colorful.Color{R: float64(red), G: float64(green), B: float64(blue)}

			if ic.DistanceLab(tc) == 0 {
				dmc = c.ColorName
				floss = c.Floss
			}
		}

	}

	return dmc, floss

}

// Below methods for converting Hex to RGB/LAB/HSV colorspaces currently not
// working. Will be implemented in v2

// Convenience functions for convertion Hexcode color values to other color spaces
func (d *DmcColors) HexToRgb(hex string) (float64, float64, float64) {
	c, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}

	_r := c.R
	_g := c.G
	_b := c.B

	return _r, _g, _b
}

func (d *DmcColors) HexToHsv(hex string) (float64, float64, float64) {
	c, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}

	return c.Hsv()
}

func (d *DmcColors) HexToLab(hex string) (float64, float64, float64) {
	c, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}

	return c.Lab()
}
