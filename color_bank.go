// Bank of colors who's rbg values already correspond to DMC string color values
// These values are used in dmc.go to find the closest matching dmc color to
// the given rgb value
package dmc

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

var colorBank []DefColor

func FillColorBank() *DmcColors {

	// TODO: Create scraper to scrape values from
	// threadcolors.com and unmarshal into colorBank
	// (All values below pulled from threadcolors.com)
	colorBank = []DefColor{
		{
			ColorName: "Salmon Very Light",
			Floss:     "3713",
			Hex:       "#ffe2e2",
			R:         "255",
			G:         "226",
			B:         "225",
		},
		{
			ColorName: "Salmon Light",
			Floss:     "761",
			Hex:       "#ffc9c9",
			R:         "255",
			G:         "201",
			B:         "201",
		},
		{
			ColorName: "Salmon",
			Floss:     "760",
			Hex:       "#f5adad",
			R:         "245",
			G:         "173",
			B:         "173",
		},
		{
			ColorName: "Salmon Medium",
			Floss:     "3712",
			Hex:       "#f18787",
			R:         "241",
			G:         "135",
			B:         "135",
		},
		{
			ColorName: "Salmon Dark",
			Floss:     "3328",
			Hex:       "#e36d6d",
			R:         "277",
			G:         "109",
			B:         "109",
		},
	}

	colorMap := make(map[string]string)

	for _, c := range colorBank {
		colorMap[c.Hex] = c.ColorName
	}

	return &DmcColors{
		ColorBank: colorBank,
		HexMap:    colorMap,
	}
}
