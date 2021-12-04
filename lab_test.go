package dmc

import (
	"fmt"
	"strconv"
)

func (suite *DmcTestSuite) TestLabRgb() {
	cb := NewColorBank()
	// Due to the nature of floats, the Lab to Rgb calculation doesn't return super precise
	// values (for example, the L value of 74.28679665732454 could return 245.0000000000006
	// or 245.00000000006 as the R value, a difference of 2 decimal points). So for
	// simplicity's sake, this test rounds the returned float to a precizion of 1 decimal
	// place
	i, j, k := cb.LabRgb(74.28679665732454, 22.545819185702953, 9.144645443127331)
	x := fmt.Sprintf("%.1f", i)
	y := fmt.Sprintf("%.1f", j)
	z := fmt.Sprintf("%.1f", k)
	r, _ := strconv.ParseFloat(x, 64)
	g, _ := strconv.ParseFloat(y, 64)
	b, _ := strconv.ParseFloat(z, 64)
	suite.Equal(245.0, r)
	suite.Equal(173.0, g)
	suite.Equal(173.0, b)
}

func (suite *DmcTestSuite) TestLabHsv() {
	cb := NewColorBank()
	// Due to the nature of floats, the Lab to Rgb calculation doesn't return super precise
	// values (for example, the B value of 9.144645443127331 could return 245.0000000000006
	// or 245.00000000006 as the R value, a difference of 2 decimal points). So for
	// simplicity's sake, this test rounds the returned float for V to a precizion of 1 decimal
	// place
	h, s, x := cb.LabHsv(74.28679665732454, 22.545819185702953, 9.144645443127331)
	y := fmt.Sprintf("%.1f", x)
	v, _ := strconv.ParseFloat(y, 64)
	suite.Equal(0.000000, h)
	suite.Equal(0.2938775510204082, s)
	suite.Equal(245.0, v)
}

func (suite *DmcTestSuite) TestLabHex() {
	cb := NewColorBank()
	hex := cb.LabHex(74.28679665732454, 22.545819185702953, 9.144645443127331)
	suite.Equal("#0b5353", hex)
}
