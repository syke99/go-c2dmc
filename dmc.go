// TODO: Rename file to dmc.go

package dmc

func NewColorBank() *DmcColors {

	var cb *DmcColors

	cb = fillColorBank()

	return cb
}

// Determine DMC color from RBG (Red, Green, Blue) values
func (d *DmcColors) Rgb(r float64, g float64, b float64) (string, string) {
	return d.RgbToDmc(r, g, b)
}

// Determine DMC color from Hex value
func (d *DmcColors) Hex(hex string) (string, string) {
	return d.HexToDmc(hex)
}

// Determine DMC color from HSV (Hue, Saturation, Value) values
func (d *DmcColors) Hsv(h float64, s float64, v float64) (string, string) {
	return d.HsvToDmc(h, s, v)
}

// For convenience's sake, dmc provides the below functions for converting the four
// used colorspaces/hex value into the three other colorspaces used in this package/hex values.
// For a larger availability of colorspaces to change to and from, reference
// the package github.com/lucasb-eyer/go-colorful

// Converting from RGB colorspace
func (d *DmcColors) RgbHex(r float64, g float64, b float64) string {
	return d.RgbToHex(r, g, b)
}

func (d *DmcColors) RgbHsv(r float64, g float64, b float64) (float64, float64, float64) {
	return d.RgbToHsv(r, g, b)
}

func (d *DmcColors) RgbLab(r float64, g float64, b float64) (float64, float64, float64) {
	return d.RgbToLab(r, g, b)
}

// Converting from HSV colorspace
func (d *DmcColors) HsvRgb(h float64, s float64, v float64) (float64, float64, float64) {
	return d.HsvToRgb(h, s, v)
}

func (d *DmcColors) HsvLab(h float64, s float64, v float64) (float64, float64, float64) {
	return d.HsvToLab(h, s, v)
}

func (d *DmcColors) HsvHex(h float64, s float64, v float64) string {
	return d.HsvToHex(h, s, v)
}

// Converting from Lab colorspace
func (d *DmcColors) LabRgb(l float64, a float64, b float64) (float64, float64, float64) {
	return d.LabToRgb(l, a, b)
}

func (d *DmcColors) LabHsv(l float64, a float64, b float64) (float64, float64, float64) {
	return d.LabToHsv(l, a, b)
}

func (d *DmcColors) LabHex(l float64, a float64, b float64) string {
	return d.LabToHex(l, a, b)
}

// Converting from Hex values
func (d *DmcColors) HexRgb(hex string) (float64, float64, float64) {
	return d.HexToRgb(hex)
}

func (d *DmcColors) HexHsv(hex string) (float64, float64, float64) {
	return d.HexToHsv(hex)
}

func (d *DmcColors) HexLab(hex string) (float64, float64, float64) {
	return d.HexToLab(hex)
}
