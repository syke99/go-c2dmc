package hsv

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

func (m mockColorBank) MockHsvToRgb() ([]float64, error) {
	if m.fail {
		return []float64{}, errors.New("failure")
	}
	return []float64{245.0, 173.0, 173.0}, nil
}

func (m mockColorBank) MockHsvToHex() (string, error) {
	if m.fail {
		return "", errors.New("failure")
	}
	return "#0b5353", nil
}

func (m mockColorBank) MockHsvToLab() ([]float64, error) {
	if m.fail {
		return []float64{}, errors.New("failure")
	}
	return []float64{74.28679665732454, 22.545819185702953, 9.144645443127331}, nil
}

//-----------
// test mocks
//-----------

func TestHsvRgb_NoFail(t *testing.T) {
	cb := mockColorBank{}
	rgb, err := cb.MockHsvToRgb()
	assert.Equal(t, 245.0, rgb[0])
	assert.Equal(t, 173.0, rgb[1])
	assert.Equal(t, 173.0, rgb[2])
	assert.NoError(t, err)
}

func TestHsvRgb_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockHsvToRgb()
	assert.Error(t, err)
}

func TestHsvHex_NoFail(t *testing.T) {
	cb := mockColorBank{}
	hex, err := cb.MockHsvToHex()
	assert.Equal(t, "#0b5353", hex)
	assert.NoError(t, err)
}

func TestHsvHex_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockHsvToHex()
	assert.Error(t, err)
}

func TestHsvLab_NoFail(t *testing.T) {
	cb := mockColorBank{}
	lab, err := cb.MockHsvToLab()
	assert.Equal(t, 74.28679665732454, lab[0])
	assert.Equal(t, 22.545819185702953, lab[1])
	assert.Equal(t, 9.144645443127331, lab[2])
	assert.NoError(t, err)
}

func TestHsvLab_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockHsvToLab()
	assert.Error(t, err)
}
