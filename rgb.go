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
