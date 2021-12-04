package dmc

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
