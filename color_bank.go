// Bank of colors who's rbg values already correspond to DMC string color values
// These values are used in dmc_calc.go to find the closest matching dmc color to
// the given rgb value
package dmc

func FillColorBank() [][]string {

	var colorBank [][]string

	// Salmon very light
	svl := []string{"255", "226", "225", "Salmon Very Light", "3713"}
	colorBank = append(colorBank, svl)

	// Salmon light
	sl := []string{"255", "201", "201", "Salmon Light", "761"}
	colorBank = append(colorBank, sl)

	// Salmon
	s := []string{"245", "173", "173", "Salmon", "760"}
	colorBank = append(colorBank, s)

	// Salmon Medium
	sm := []string{"241", "135", "135", "Salmon Medium", "3712"}
	colorBank = append(colorBank, sm)

	// Salmon Dark
	sd := []string{"277", "109", "109", "Salmon Dark", "3328"}
	colorBank = append(colorBank, sd)

	// Salmon Very Dark
	svd := []string{"191", "45", "45", "Salmon Very Dark", "347"}
	colorBank = append(colorBank, svd)

	// Peach
	p := []string{"254", "215", "204", "Peach", "353"}
	colorBank = append(colorBank, p)

	// Coral Light
	cl := []string{"253", "156", "151", "Coral Light", "352"}
	colorBank = append(colorBank, cl)

	// Coral
	c := []string{"233", "106", "103", "Coral", "351"}
	colorBank = append(colorBank, c)

	// Coral Medium
	cm := []string{"224", "72", "72", "Coral Medium", "350"}
	colorBank = append(colorBank, cm)

	// Coral Dark
	cd := []string{"210", "16", "53", "Coral Dark", "349"}
	colorBank = append(colorBank, cd)

	// Coral Red Very Dark
	crvd := []string{"187", "5", "31", "Coral Red Very Dark", "817"}
	colorBank = append(colorBank, crvd)

	// Melon Light
	ml := []string{"255", "203", "213", "Melon Light", "3708"}
	colorBank = append(colorBank, ml)

	// Melon Medium
	mm := []string{"255", "173", "188", "Melon Medium", "3706"}
	colorBank = append(colorBank, mm)

	// Melon Dark
	md := []string{"255", "121", "146", "Melon Dark", "3705"}
	colorBank = append(colorBank, md)

	// Melon Very Dark
	mvd := []string{"231", "73", "103", "Melon Very Dark", "3801"}
	colorBank = append(colorBank, mvd)

	// Bright Red
	br := []string{"227", "29", "66", "Bright Red", "666"}
	colorBank = append(colorBank, br)

	// Red
	r := []string{"199", "43", "59", "Red", "321"}
	colorBank = append(colorBank, r)

	// Red Medium
	rm := []string{"183", "31", "51", "Red Medium", "304"}
	colorBank = append(colorBank, rm)

	// Red Dark
	rd := []string{"167", "19", "43", "Red Dark", "498"}
	colorBank = append(colorBank, rd)

	return colorBank
}
