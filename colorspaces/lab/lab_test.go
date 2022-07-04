package lab

import (
	"fmt"
	"github.com/syke99/go-c2dmc/colorBank"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabRgb(t *testing.T) {
	cb := colorBank.New()
	// Due to the nature of floats, the Lab to Rgb calculation doesn't return super precise
	// values (for example, the L value of 74.28679665732454 could return 245.0000000000006
	// or 245.00000000006 as the R value, a difference of 2 decimal points). So for
	// simplicity's sake, this test rounds the returned float to a precizion of 1 decimal
	// place
	i, j, k := cb.Lab.LabToRgb(74.28679665732454, 22.545819185702953, 9.144645443127331)
	x := fmt.Sprintf("%.1f", i)
	y := fmt.Sprintf("%.1f", j)
	z := fmt.Sprintf("%.1f", k)
	r, _ := strconv.ParseFloat(x, 64)
	g, _ := strconv.ParseFloat(y, 64)
	b, _ := strconv.ParseFloat(z, 64)
	assert.Equal(t, 245.0, r)
	assert.Equal(t, 173.0, g)
	assert.Equal(t, 173.0, b)
}

func TestLabHsv(t *testing.T) {
	cb := colorBank.New()
	// Due to the nature of floats, the Lab to Rgb calculation doesn't return super precise
	// values (for example, the B value of 9.144645443127331 could return 245.0000000000006
	// or 245.00000000006 as the R value, a difference of 2 decimal points). So for
	// simplicity's sake, this test rounds the returned float for V to a precizion of 1 decimal
	// place
	h, s, x := cb.Lab.LabToHsv(74.28679665732454, 22.545819185702953, 9.144645443127331)
	y := fmt.Sprintf("%.1f", x)
	v, _ := strconv.ParseFloat(y, 64)
	assert.Equal(t, 0.000000, h)
	assert.Equal(t, 0.2938775510204082, s)
	assert.Equal(t, 245.0, v)
}

func TestLabHex(t *testing.T) {
	cb := colorBank.New()
	hex := cb.Lab.LabToHex(74.28679665732454, 22.545819185702953, 9.144645443127331)
	assert.Equal(t, "#0b5353", hex)
}
