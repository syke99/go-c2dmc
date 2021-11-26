package dmc

import (
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

func (d *DmcColors) LabToDmc(l float64, a float64, b float64) (string, string) {

	var dis float64
	var dmc string
	var floss string

	// ic is Color struct holding the Lab values passed in to LabToDmc
	ic := colorful.Lab(l, a, b)

	for _, c := range d.ColorBank {

		red, _ := strconv.Atoi(c.R)
		green, _ := strconv.Atoi(c.G)
		blue, _ := strconv.Atoi(c.B)

		// tc is Color struct holding the Lab values of each color in the
		// color bank to test how close it is to ic
		tc := colorful.Lab(float64(red), float64(green), float64(blue))

		if dis == 0 || (ic.DistanceLab(tc) < dis) {
			dis = ic.DistanceLab(tc)
			dmc = c.ColorName
			floss = c.Floss
		}
	}

	return dmc, floss

}

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
