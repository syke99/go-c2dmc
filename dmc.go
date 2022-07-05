package dmc

import (
	"github.com/syke99/go-c2dmc/filter"
	"image/color"

	cb "github.com/syke99/go-c2dmc/colorbank"
	"github.com/syke99/go-c2dmc/internal/colorspaces"
)

func NewColorBank() *cb.DmcColors {
	return cb.New(nil)
}

func NewColorBankWithFilter(filter filter.FilterOption) *cb.DmcColors {
	return cb.New(filter)
}

// RgbA Determine DMC color from RBG (Red, Green, Blue) values
func RgbA(col color.Color) (float64, float64, float64) {
	return colorspaces.RgbaToRgb(col)
}

// Rgb Determine DMC color from RBG (Red, Green, Blue) values
func Rgb(d *cb.DmcColors, r float64, g float64, b float64) (string, string) {
	return d.Rgb.RgbToDmc(d.ColorBank, d.HexMap, r, g, b)
}

// Hex Determine DMC color from Hex value
func Hex(d *cb.DmcColors, hex string) (string, string) {
	return d.Hex.HexToDmc(d.ColorBank, d.HexMap, hex)
}

// Hsv Determine DMC color from HSV (Hue, Saturation, Value) values
func Hsv(d *cb.DmcColors, h float64, s float64, v float64) (string, string) {
	return d.Hsv.HsvToDmc(d.ColorBank, d.HexMap, h, s, v)
}

// Lab Determine DMC color from HSV (Hue, Saturation, Value) values
func Lab(d *cb.DmcColors, l float64, a float64, b float64) (string, string) {
	return d.Lab.LabToDmc(d.ColorBank, d.HexMap, l, a, b)
}

// For convenience's sake, dmc provides the below functions for converting the four
// used colorspaces/hex value into the three other colorspaces used in this package/hex values.
// For a larger availability of colorspaces to change to and from, reference
// the package github.com/lucasb-eyer/go-colorful

// RgbHex Converting from Rgb to Hex
func RgbHex(d *cb.DmcColors, r float64, g float64, b float64) string {
	return d.Rgb.RgbToHex(r, g, b)
}

// RgbHsv Converting from Rgb to Hsv
func RgbHsv(d *cb.DmcColors, r float64, g float64, b float64) (float64, float64, float64) {
	return d.Rgb.RgbToHsv(r, g, b)
}

// RgbLab Converting from Rgb to Lab
func RgbLab(d *cb.DmcColors, r float64, g float64, b float64) (float64, float64, float64) {
	return d.Rgb.RgbToLab(r, g, b)
}

// HsvRgb Converting from Hsv to Rgb
func HsvRgb(d *cb.DmcColors, h float64, s float64, v float64) (float64, float64, float64) {
	return d.Hsv.HsvToRgb(h, s, v)
}

// HsvLab Converting from Hsv to Lab
func HsvLab(d *cb.DmcColors, h float64, s float64, v float64) (float64, float64, float64) {
	return d.Hsv.HsvToLab(h, s, v)
}

// HsvHex Converting from Hsv to Hex
func HsvHex(d *cb.DmcColors, h float64, s float64, v float64) string {
	return d.Hsv.HsvToHex(h, s, v)
}

// LabRgb Converting from Lab to Rgb
func LabRgb(d *cb.DmcColors, l float64, a float64, b float64) (float64, float64, float64) {
	return d.Lab.LabToRgb(l, a, b)
}

// LabHsv Converting from Lab to Hsv
func LabHsv(d *cb.DmcColors, l float64, a float64, b float64) (float64, float64, float64) {
	return d.Lab.LabToHsv(l, a, b)
}

// LabHex Converting from Lab to Hex
func LabHex(d *cb.DmcColors, l float64, a float64, b float64) string {
	return d.Lab.LabToHex(l, a, b)
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
