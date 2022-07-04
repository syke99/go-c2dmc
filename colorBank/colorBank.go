// Package colorBank provides a bank of colors who's rbg values already correspond to DMC string color values
// These values are used in dmc.go to find the closest matching dmc color to
// the given rgb value
package colorBank

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/syke99/go-c2dmc/colorspaces"
	"github.com/syke99/go-c2dmc/colorspaces/hex"
	"github.com/syke99/go-c2dmc/colorspaces/hsv"
	"github.com/syke99/go-c2dmc/colorspaces/lab"
	"github.com/syke99/go-c2dmc/colorspaces/rgb"
)

type DmcColors struct {
	ColorBank []colorspaces.DefColor
	HexMap    map[string]string
	Rgb       rgb.Rgb
	Hex       hex.Hex
	Hsv       hsv.Hsv
	Lab       lab.Lab
}

func New() *DmcColors {
	cb := fillColorBank()
	hm := fillHexMap(cb)
	colors := &DmcColors{
		ColorBank: cb,
		HexMap:    hm,
		Rgb:       rgb.Rgb{},
		Hex:       hex.Hex{},
		Hsv:       hsv.Hsv{},
		Lab:       lab.Lab{},
	}
	return colors
}

func fillColorBank() []colorspaces.DefColor {
	var colorBank []colorspaces.DefColor

	page := "https://floss.maxxmint.com/dmc_to_rgb.php"
	resp, err := http.Get(page)
	if err != nil {
		println(err.Error())
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		println(err.Error())
	}

	println(doc.Find("tbody").Children().Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			dc := colorspaces.DefColor{}
			println(s.Children().Each(func(i int, s *goquery.Selection) {
				switch i {
				case 1:
					dc.Floss = s.Text()
				case 2:
					dc.ColorName = s.Text()
				case 3:
					dc.Hex = s.Text()
				case 4:
					dc.R = s.Text()
				case 5:
					dc.G = s.Text()
				case 6:
					dc.B = s.Text()
				}
			}))
			colorBank = append(colorBank, dc)
		}
	}))

	return colorBank
}

func fillHexMap(colorBank []colorspaces.DefColor) map[string]string {

	colorMap := make(map[string]string)

	for _, c := range colorBank {
		colorMap[c.Hex] = c.ColorName
	}

	return colorMap
}
