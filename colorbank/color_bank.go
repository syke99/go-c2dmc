// Package colorbank provides a bank of colors who's rbg values already correspond to DMC string color values
// These values are used in dmc.go to find the closest matching dmc color to
// the given rgb value
package colorBank

import (
	"net/http"

	"github.com/syke99/go-c2dmc/filter"
	"github.com/syke99/go-c2dmc/internal/colorspaces/hex"
	"github.com/syke99/go-c2dmc/internal/colorspaces/hsv"
	"github.com/syke99/go-c2dmc/internal/colorspaces/lab"
	"github.com/syke99/go-c2dmc/internal/colorspaces/rgb"
	"github.com/syke99/go-c2dmc/internal/pkg"

	"github.com/PuerkitoBio/goquery"
)

type DmcColors struct {
	ColorBank map[string]pkg.DefColor
	Rgb       rgb.Rgb
	Hex       hex.Hex
	Hsv       hsv.Hsv
	Lab       lab.Lab
}

func New(filter filter.FilterOption) *DmcColors {
	cb := fillColorBank(filter)
	cbm := fillHexMap(cb)
	colors := &DmcColors{
		ColorBank: cbm,
		Rgb:       rgb.Rgb{},
		Hex:       hex.Hex{},
		Hsv:       hsv.Hsv{},
		Lab:       lab.Lab{},
	}
	return colors
}

func fillColorBank(fltr filter.FilterOption) []pkg.DefColor {
	var colorBank []pkg.DefColor

	page := pkg.Page
	resp, err := http.Get(page)
	if err != nil {
		println(err.Error())
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		println(err.Error())
	}

	if fltr == nil {
		colorBank = fillColorBankNoFilter(doc, colorBank)
		return colorBank
	}

	colorBank = fillColorBankWithFilter(doc, colorBank, fltr)

	return colorBank
}

func fillColorBankNoFilter(doc *goquery.Document, colorBank []pkg.DefColor) []pkg.DefColor {
	doc.Find("tbody").Children().Each(func(i int, s *goquery.Selection) {
		if i >= 1 {
			dc := pkg.DefColor{
				ColorName: "",
				Floss:     "",
				Hex:       "",
				R:         "",
				G:         "",
				B:         "",
			}
			s.Children().Each(func(i int, s *goquery.Selection) {
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
			})
			colorBank = append(colorBank, dc)
		}
	})

	return colorBank
}

func fillColorBankWithFilter(doc *goquery.Document, colorBank []pkg.DefColor, fltr filter.FilterOption) []pkg.DefColor {
	filterFunc := func(i int, sel *goquery.Selection) bool {
		if _, ok := fltr[sel.Children().First().Text()]; ok {
			return true
		}
		return false
	}

	doc.Find("tbody").Children().Each(func(i int, s *goquery.Selection) {
		if i >= 1 {
			dc := pkg.DefColor{
				ColorName: "",
				Floss:     "",
				Hex:       "",
				R:         "",
				G:         "",
				B:         "",
			}
			s.Children().FilterFunction(filterFunc).Each(func(i int, s *goquery.Selection) {
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
			})
			colorBank = append(colorBank, dc)
		}
	})

	return colorBank
}

func fillHexMap(colorBank []pkg.DefColor) map[string]pkg.DefColor {

	colorMap := make(map[string]pkg.DefColor)

	for _, color := range colorBank {
		colorMap[color.Hex] = color
	}

	return colorMap
}
