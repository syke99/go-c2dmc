// TODO: Rename file to dmc_calc.go

package dmc

import (
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
)

var cb []DefColor

func init() {
	cb = FillColorBank()
}

func RgbToDmc(r float64, g float64, b float64) (string, int) {

	var dis float64
	var dmc string
	var floss int

	// ic is Color struct holding the RGB values passed in to RgbToDmc
	ic := colorful.Color{R: r, G: g, B: b}

	for _, c := range cb {

		red, _ := strconv.Atoi(c.R)
		green, _ := strconv.Atoi(c.G)
		blue, _ := strconv.Atoi(c.B)
		f, _ := strconv.Atoi(c.Floss)

		// tc is Color struct holding the RGB values of each color in the
		// color bank to test how close it is to ic
		tc := colorful.Color{R: float64(red), G: float64(green), B: float64(blue)}

		if dis == 0 || (ic.DistanceLab(tc) < dis) {
			dis = ic.DistanceLab(tc)
			dmc = c.ColorName
			floss = f
		}
	}

	return dmc, floss

}
