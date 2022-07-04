package rgb

import (
	"github.com/syke99/go-c2dmc/colorBank"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRgbHex(t *testing.T) {
	cb := colorBank.New()
	hex := cb.Rgb.RgbToHex(245.0, 173.0, 173.0)
	assert.Equal(t, "#0b5353", hex)
}

func TestRgbHsv(t *testing.T) {
	cb := colorBank.New()
	h, s, v := cb.Rgb.RgbToHsv(245.0, 173.0, 173.0)
	assert.Equal(t, 0.000000, h)
	assert.Equal(t, 0.2938775510204082, s)
	assert.Equal(t, 245.000000, v)
}

func TestRgbLab(t *testing.T) {
	cb := colorBank.New()
	l, a, b := cb.Rgb.RgbToLab(245.0, 173.0, 173.0)
	assert.Equal(t, 74.28679665732454, l)
	assert.Equal(t, 22.545819185702953, a)
	assert.Equal(t, 9.144645443127331, b)
}
