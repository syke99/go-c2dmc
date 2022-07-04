package rgb

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//-----------
// mocks
//-----------

type mockColorBank struct {
	fail bool
}

func (m mockColorBank) MockRgbToHex() (string, error) {
	if m.fail {
		return "", errors.New("failure")
	}
	return "#0b5353", nil
}

func (m mockColorBank) MockRgbToHsv() ([]float64, error) {
	if m.fail {
		return []float64{}, errors.New("failure")
	}
	return []float64{0.000000, 0.2938775510204082, 245.000000}, nil
}

func (m mockColorBank) MockRgbToLab() ([]float64, error) {
	if m.fail {
		return []float64{}, errors.New("failure")
	}
	return []float64{74.28679665732454, 22.545819185702953, 9.144645443127331}, nil
}

//-----------
// test mocks
//-----------

func TestRgbHex_NoFail(t *testing.T) {
	cb := mockColorBank{}
	hex, err := cb.MockRgbToHex()
	assert.Equal(t, "#0b5353", hex)
	assert.NoError(t, err)
}

func TestRgbHex_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockRgbToHex()
	assert.Error(t, err)
}

func TestRgbHsv_NoFail(t *testing.T) {
	cb := mockColorBank{}
	hsv, err := cb.MockRgbToHsv()
	assert.Equal(t, 0.000000, hsv[0])
	assert.Equal(t, 0.2938775510204082, hsv[1])
	assert.Equal(t, 245.000000, hsv[2])
	assert.NoError(t, err)
}

func TestRgbHsv_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockRgbToHsv()
	assert.Error(t, err)
}

func TestRgbLab_NoFail(t *testing.T) {
	cb := mockColorBank{}
	lab, err := cb.MockRgbToLab()
	assert.Equal(t, 74.28679665732454, lab[0])
	assert.Equal(t, 22.545819185702953, lab[1])
	assert.Equal(t, 9.144645443127331, lab[2])
	assert.NoError(t, err)
}

func TestRgbLab_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockRgbToLab()
	assert.Error(t, err)
}
