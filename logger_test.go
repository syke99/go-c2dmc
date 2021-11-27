package dmc

import (
	"errors"
	"io"
)

func (suite *DmcTestSuite) TestlogWritter() {
	suite.Implements((*io.Writer)(nil), logWriter())
}

func (suite *DmcTestSuite) TestNewLogger() {
	suite.IsType(&Logger{}, NewLogger())
}

func (suite *DmcTestSuite) TestLoggerLogError() {
	err := errors.New("test error")
	suite.Error(suite.Logger.LogError(err, ""))
}
