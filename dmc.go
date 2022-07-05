package dmc

import (
	"image/color"

	cb "github.com/syke99/go-c2dmc/colorbank"
	"github.com/syke99/go-c2dmc/filter"
	"github.com/syke99/go-c2dmc/internal/colorspaces"
)

// ColorBank is the color bank used for comparing colors, either with
// the full list of DMC thread colors, or filtered with a FilterOption
type ColorBank struct {
	// ColorBank contains filtered or unexported fields
	bank *cb.DmcColors
}

// NewColorBank returns a ColorBank without a filter
func NewColorBank() ColorBank {
	return ColorBank{
		bank: cb.New(nil),
	}
}

// NewColorBankWithFilter returns a ColorBank with the filter that is passed in
func NewColorBankWithFilter(fltr filter.FilterOption) ColorBank {
	return ColorBank{
		bank: cb.New(fltr),
	}
}

// RgbA Determine DMC color from RBG (Red, Green, Blue) values
func (cb ColorBank) RgbA(col color.Color) (float64, float64, float64) {
	return colorspaces.RgbaToRgb(col)
}

// Rgb Determine DMC color from RBG (Red, Green, Blue) values
func (cb ColorBank) Rgb(r float64, g float64, b float64) (string, string) {
	return cb.bank.Rgb.RgbToDmc(cb.bank.ColorBank, cb.bank.HexMap, r, g, b)
}

// Hex Determine DMC color from Hex value
func (cb ColorBank) Hex(hex string) (string, string) {
	return cb.bank.Hex.HexToDmc(cb.bank.ColorBank, cb.bank.HexMap, hex)
}

// Hsv Determine DMC color from HSV (Hue, Saturation, Value) values
func (cb ColorBank) Hsv(h float64, s float64, v float64) (string, string) {
	return cb.bank.Hsv.HsvToDmc(cb.bank.ColorBank, cb.bank.HexMap, h, s, v)
}

// Lab Determine DMC color from HSV (Hue, Saturation, Value) values
func (cb ColorBank) Lab(l float64, a float64, b float64) (string, string) {
	return cb.bank.Lab.LabToDmc(cb.bank.ColorBank, cb.bank.HexMap, l, a, b)
}

// For convenience's sake, dmc provides the below functions for converting the four
// used colorspaces/hex value into the three other colorspaces used in this package/hex values.
// For a larger availability of colorspaces to change to and from, reference
// the package github.com/lucasb-eyer/go-colorful

// RgbHex Converting from Rgb to Hex
func (cb ColorBank) RgbHex(r float64, g float64, b float64) string {
	return cb.bank.Rgb.RgbToHex(r, g, b)
}

// RgbHsv Converting from Rgb to Hsv
func (cb ColorBank) RgbHsv(r float64, g float64, b float64) (float64, float64, float64) {
	return cb.bank.Rgb.RgbToHsv(r, g, b)
}

// RgbLab Converting from Rgb to Lab
func (cb ColorBank) RgbLab(r float64, g float64, b float64) (float64, float64, float64) {
	return cb.bank.Rgb.RgbToLab(r, g, b)
}

// HsvRgb Converting from Hsv to Rgb
func (cb ColorBank) HsvRgb(h float64, s float64, v float64) (float64, float64, float64) {
	return cb.bank.Hsv.HsvToRgb(h, s, v)
}

// HsvLab Converting from Hsv to Lab
func (cb ColorBank) HsvLab(h float64, s float64, v float64) (float64, float64, float64) {
	return cb.bank.Hsv.HsvToLab(h, s, v)
}

// HsvHex Converting from Hsv to Hex
func (cb ColorBank) HsvHex(h float64, s float64, v float64) string {
	return cb.bank.Hsv.HsvToHex(h, s, v)
}

// LabRgb Converting from Lab to Rgb
func (cb ColorBank) LabRgb(l float64, a float64, b float64) (float64, float64, float64) {
	return cb.bank.Lab.LabToRgb(l, a, b)
}

// LabHsv Converting from Lab to Hsv
func (cb ColorBank) LabHsv(l float64, a float64, b float64) (float64, float64, float64) {
	return cb.bank.Lab.LabToHsv(l, a, b)
}

// LabHex Converting from Lab to Hex
func (cb ColorBank) LabHex(l float64, a float64, b float64) string {
	return cb.bank.Lab.LabToHex(l, a, b)
}

// The methods below for convertiong from Hex to RGB/LAB/HSV colorspaces
// currently not working. Will be implemented in v2

// Converting from Hex values
// func (d *DmcColors) HexRgb(hex string) (float64, float64, float64) {
// 	return d.HexToRgb(hex)
// }

// func (d *DmcColors) HexHsv(hex string) (float64, float64, float64) {
// 	return d.HexToHsv(hex)
// }

// func (d *DmcColors) HexLab(hex string) (float64, float64, float64) {
// 	return d.HexToLab(hex)
// }
