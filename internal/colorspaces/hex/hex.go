package hex

import (
	"log"
	"strconv"

	"github.com/syke99/go-c2dmc/internal/pkg"

	"github.com/lucasb-eyer/go-colorful"
)

type Hex struct{}

func (d Hex) HexToDmc(cbm map[string]pkg.DefColor, hex string) (string, string) {

	var previousDistance float64
	var dmc string
	var floss string

	// Search for hex in hm to check for exact matches. If it exists, loop through
	// cb for the color name and floss number

	if val, ok := cbm[hex]; ok {
		return val.ColorName, val.Floss
	}

	// Otherwise, use the hex value to create a new colorful.Color
	// and search through the cb for the closest match (in L*a*b colorspace)

	hc, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}

	// ic is Color struct holding the RGB values passed in to RgbToDmc
	ic := colorful.Color{R: hc.R, G: hc.G, B: hc.B}

	var i int
	for _, c := range cbm {

		red, _ := strconv.Atoi(c.R)
		green, _ := strconv.Atoi(c.G)
		blue, _ := strconv.Atoi(c.B)

		// tc is Color struct holding the RGB values of each color in the
		// color bank to test how close it is to ic
		tc := colorful.Color{R: float64(red), G: float64(green), B: float64(blue)}

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

// Below methods for converting Hex to RGB/LAB/HSV colorspaces currently not
// working. Will be implemented in v2

// Convenience functions for convertion Hexcode color values to other color spaces
func (d Hex) HexToRgb(hex string) (float64, float64, float64) {
	c, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}

	_r := c.R
	_g := c.G
	_b := c.B

	return _r, _g, _b
}

func (d Hex) HexToHsv(hex string) (float64, float64, float64) {
	c, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}

	return c.Hsv()
}

func (d Hex) HexToLab(hex string) (float64, float64, float64) {
	c, err := colorful.Hex(hex)
	if err != nil {
		log.Fatal(err)
	}

	return c.Lab()
}
