package dmc

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DmcTestSuite struct {
	suite.Suite
	Logger *Logger
}

func (suite *DmcTestSuite) SetupTest() {
	suite.Logger = NewLogger()
}

func TestDmcTestSuite(t *testing.T) {
	suite.Run(t, new(DmcTestSuite))
}

func (suite *DmcTestSuite) TestNewColorBank() {
	suite.IsType(&DmcColors{}, NewColorBank())
}

func (suite *DmcTestSuite) TestRgb() {
	cb := NewColorBank()
	n, f := cb.Rgb(245.0, 173.0, 173.0)
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
}

func (suite *DmcTestSuite) TestHex() {
	cb := NewColorBank()
	n, f := cb.Hex("#0b5353")
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
}

func (suite *DmcTestSuite) TestHsv() {
	cb := NewColorBank()
	n, f := cb.Hsv(0.000000, 0.293878, 245.000000)
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
}

func (suite *DmcTestSuite) TestLab() {
	cb := NewColorBank()
	n, f := cb.Lab(74.286797, 22.545819, 9.144645)
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
}

func (suite *DmcTestSuite) TestRgbHex() {
	cb := NewColorBank()
	hex := cb.RgbHex(245.0, 173.0, 173.0)
	suite.Equal("#0b5353", hex)
}

func (suite *DmcTestSuite) TestRgbHsv() {
	cb := NewColorBank()
	h, s, v := cb.RgbHsv(245.0, 173.0, 173.0)
	suite.Equal(0.000000, h)
	suite.Equal(0.2938775510204082, s)
	suite.Equal(245.000000, v)
}

func (suite *DmcTestSuite) TestRgbLab() {
	cb := NewColorBank()
	l, a, b := cb.RgbLab(245.0, 173.0, 173.0)
	suite.Equal(74.28679665732454, l)
	suite.Equal(22.545819185702953, a)
	suite.Equal(9.144645443127331, b)
}

func (suite *DmcTestSuite) TestHsvRgb() {
	cb := NewColorBank()
	r, g, b := cb.HsvRgb(0.000000, 0.2938775510204082, 245.000000)
	suite.Equal(245.0, r)
	suite.Equal(173.0, g)
	suite.Equal(173.0, b)
}

func (suite *DmcTestSuite) TestHsvHex() {
	cb := NewColorBank()
	hex := cb.HsvHex(0.000000, 0.2938775510204082, 245.000000)
	suite.Equal("#0b5353", hex)
}

func (suite *DmcTestSuite) TestHsvLab() {
	cb := NewColorBank()
	l, a, b := cb.HsvLab(0.000000, 0.2938775510204082, 245.000000)
	suite.Equal(74.28679665732454, l)
	suite.Equal(22.545819185702953, a)
	suite.Equal(9.144645443127331, b)
}

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

// Below tests will be implementd in v2 whenever HEX to RGB/LAB/HSV is functional

// func (suite *DmcTestSuite) TestHexRgb() {
// 	cb := NewColorBank()
// 	i, j, k := cb.HexRgb("#0b5353")
// 	x := fmt.Sprintf("%.1f", i)
// 	y := fmt.Sprintf("%.1f", j)
// 	z := fmt.Sprintf("%.1f", k)
// 	r, _ := strconv.ParseFloat(x, 64)
// 	g, _ := strconv.ParseFloat(y, 64)
// 	b, _ := strconv.ParseFloat(z, 64)
// 	suite.Equal(245.0, r)
// 	suite.Equal(173.0, g)
// 	suite.Equal(173.0, b)
// }

// func (suite *DmcTestSuite) TestHexHsv() {
// 	cb := NewColorBank()
// 	h, s, x := cb.HexHsv("#0b5353")
// 	y := fmt.Sprintf("%.1f", x)
// 	v, _ := strconv.ParseFloat(y, 64)
// 	suite.Equal(0.000000, h)
// 	suite.Equal(0.2938775510204082, s)
// 	suite.Equal(245.0, v)
// }

// func (suite *DmcTestSuite) TestHexLab() {
// 	cb := NewColorBank()
// 	l, a, b := cb.HexLab("#0b5353")
// 	suite.Equal(74.28679665732454, l)
// 	suite.Equal(22.545819185702953, a)
// 	suite.Equal(9.144645443127331, b)
// }
