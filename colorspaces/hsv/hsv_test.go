package hsv

import (
	"github.com/syke99/go-c2dmc/colorBank"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHsvRgb(t *testing.T) {
	cb := colorBank.New()
	r, g, b := cb.Hsv.HsvToRgb(0.000000, 0.2938775510204082, 245.000000)
	assert.Equal(t, 245.0, r)
	assert.Equal(t, 173.0, g)
	assert.Equal(t, 173.0, b)
}

func TestHsvHex(t *testing.T) {
	cb := colorBank.New()
	hex := cb.Hsv.HsvToHex(0.000000, 0.2938775510204082, 245.000000)
	assert.Equal(t, "#0b5353", hex)
}

func TestHsvLab(t *testing.T) {
	cb := colorBank.New()
	l, a, b := cb.Hsv.HsvToLab(0.000000, 0.2938775510204082, 245.000000)
	assert.Equal(t, 74.28679665732454, l)
	assert.Equal(t, 22.545819185702953, a)
	assert.Equal(t, 9.144645443127331, b)
}
