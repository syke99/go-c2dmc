package rgb

import (
	"strconv"

	"github.com/syke99/go-c2dmc/internal/pkg"

	"github.com/lucasb-eyer/go-colorful"
)

type Rgb struct{}

func (d Rgb) RgbToDmc(cb []pkg.DefColor, hm map[string]string, r float64, g float64, b float64) (string, string) {

	var previousDistance float64
	var dmc string
	var floss string

	// ic is Color struct holding the RGB values passed in to RgbToDmc
	ic := colorful.Color{R: r, G: g, B: b}

	// Search for hex in hm to check for exact matches. If it exists, loop through
	// cb for the color name and floss number
	hex := ic.Hex()

	if val, ok := hm[hex]; ok {
		for _, c := range cb {
			if c.ColorName == val {
				return c.ColorName, c.Floss
			}
		}
		// Otherwise, use the hex value to create a new colorful.Color
		// and search through the cb for the closest match (in L*a*b colorspace)
	} else {
		for i, c := range cb {

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

func (d Rgb) RgbToLab(r float64, g float64, b float64) (float64, float64, float64) {
	c := colorful.Color{R: r, G: g, B: b}

	return c.Lab()
}

func (d Rgb) RgbToHex(r float64, g float64, b float64) string {
	c := colorful.Color{R: r, G: g, B: b}

	return c.Hex()
}

func (d Rgb) RgbToHsv(r float64, g float64, b float64) (float64, float64, float64) {
	c := colorful.Color{R: r, G: g, B: b}

	return c.Hsv()
}
