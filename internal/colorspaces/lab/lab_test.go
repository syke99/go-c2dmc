package lab

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

func (m mockColorBank) MockLabToRgb() ([]float64, error) {
	if m.fail {
		return []float64{}, errors.New("failure")
	}
	return []float64{245.0, 173.0, 173.0}, nil
}

func (m mockColorBank) MockLabToHsv() ([]float64, error) {
	if m.fail {
		return []float64{}, errors.New("failure")
	}
	return []float64{0.000000, 0.2938775510204082, 245.0}, nil
}

func (m mockColorBank) MockLabToHex() (string, error) {
	if m.fail {
		return "", errors.New("failure")
	}
	return "#0b5353", nil
}

func TestLabRgb_NoFail(t *testing.T) {
	cb := mockColorBank{}
	rgb, err := cb.MockLabToRgb()
	assert.Equal(t, 245.0, rgb[0])
	assert.Equal(t, 173.0, rgb[1])
	assert.Equal(t, 173.0, rgb[2])
	assert.NoError(t, err)
}

func TestLabRgb_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockLabToRgb()
	assert.Error(t, err)
}

func TestLabHsv_NoFail(t *testing.T) {
	cb := mockColorBank{}
	hsv, err := cb.MockLabToHsv()
	assert.Equal(t, 0.000000, hsv[0])
	assert.Equal(t, 0.2938775510204082, hsv[1])
	assert.Equal(t, 245.0, hsv[2])
	assert.NoError(t, err)
}

func TestLabHsv_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockLabToHsv()
	assert.Error(t, err)
}

func TestLabHex_NoFail(t *testing.T) {
	cb := mockColorBank{}
	hex, err := cb.MockLabToHex()
	assert.Equal(t, "#0b5353", hex)
	assert.NoError(t, err)
}

func TestLabHex_Fail(t *testing.T) {
	cb := mockColorBank{fail: true}
	_, err := cb.MockLabToHex()
	assert.Error(t, err)
}
