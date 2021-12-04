package dmc

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
