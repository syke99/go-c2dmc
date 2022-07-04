package dmc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DmcTestSuite struct {
	suite.Suite
}

//-----------
// mocks
//-----------

type mockColorBank struct {
	fail bool
}

func MockNewColorBank() *mockColorBank {
	return &mockColorBank{}
}

func (m mockColorBank) MockRgb() (string, string, error) {
	if m.fail {
		return "", "", errors.New("failure")
	}
	return "Salmon", "760", nil
}

func (m mockColorBank) MockHsv() (string, string, error) {
	if m.fail {
		return "", "", errors.New("failure")
	}
	return "Salmon", "760", nil
}

func (m mockColorBank) MockLab() (string, string, error) {
	if m.fail {
		return "", "", errors.New("failure")
	}
	return "Salmon", "760", nil
}

//-----------
// test mocks
//-----------

func TestDmcTestSuite(t *testing.T) {
	suite.Run(t, new(DmcTestSuite))
}

func (suite *DmcTestSuite) TestNewColorBank() {
	suite.IsType(&mockColorBank{}, MockNewColorBank())
}

func (suite *DmcTestSuite) TestRgb_NoFail() {
	cb := mockColorBank{}
	n, f, err := cb.MockRgb()
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
	suite.NoError(err)
}

func (suite *DmcTestSuite) TestRgb_Fail() {
	cb := mockColorBank{fail: true}
	n, f, err := cb.MockRgb()
	suite.Equal("", n)
	suite.Equal("", f)
	suite.Error(err)
}

// func (suite *DmcTestSuite) TestHex() {
// 	cb := NewColorBank()
// 	n, f := cb.Hex("#0b5353")
// 	suite.Equal("Salmon", n)
// 	suite.Equal("760", f)
// }

func (suite *DmcTestSuite) TestHsv_NoFail() {
	cb := mockColorBank{}
	n, f, err := cb.MockHsv()
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
	suite.NoError(err)
}

func (suite *DmcTestSuite) TestHsv_Fail() {
	cb := mockColorBank{fail: true}
	n, f, err := cb.MockRgb()
	suite.Equal("", n)
	suite.Equal("", f)
	suite.Error(err)
}

func (suite *DmcTestSuite) TestLab_NoFail() {
	cb := mockColorBank{}
	n, f, err := cb.MockLab()
	suite.Equal("Salmon", n)
	suite.Equal("760", f)
	suite.NoError(err)
}
