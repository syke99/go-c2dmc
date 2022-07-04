package dmc

import (
	"github.com/syke99/go-c2dmc/colorBank"
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
	suite.IsType(&colorBank.DmcColors{}, NewColorBank())
}

func (suite *DmcTestSuite) TestRgb() {
	cb := NewColorBank()
	n, f := Rgb(cb, 245.0, 173.0, 173.0)
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
}

// func (suite *DmcTestSuite) TestHex() {
// 	cb := NewColorBank()
// 	n, f := cb.Hex("#0b5353")
// 	suite.Equal("Salmon", n)
// 	suite.Equal("760", f)
// }

func (suite *DmcTestSuite) TestHsv() {
	cb := NewColorBank()
	n, f := Hsv(cb, 0.000000, 0.293878, 245.000000)
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
}

func (suite *DmcTestSuite) TestLab() {
	cb := NewColorBank()
	n, f := Lab(cb, 74.286797, 22.545819, 9.144645)
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
}
