package dmc

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
