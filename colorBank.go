// Bank of colors who's rbg values already correspond to DMC string color values
// These values are used in dmc.go to find the closest matching dmc color to
// the given rgb value
package dmc

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type DefColor struct {
	ColorName string
	Floss     string
	Hex       string
	R         string
	G         string
	B         string
}

type DmcColors struct {
	ColorBank []DefColor
	HexMap    map[string]string
}

func fillColorBank() *DmcColors {

	var colorBank []DefColor

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
			dc := DefColor{}
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

	colorMap := make(map[string]string)

	for _, c := range colorBank {
		colorMap[c.Hex] = c.ColorName
	}

	return &DmcColors{
		ColorBank: colorBank,
		HexMap:    colorMap,
	}
}
