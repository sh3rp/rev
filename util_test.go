package rev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	a := []byte{0x01, 0x02, 0x03}
	b := []byte{0x01, 0x02, 0x03}
	c := []byte{0x02, 0x03, 0x01}
	d := []byte{0x01, 0x02, 0x03, 0x04}

	assert.True(t, eq(a, b))
	assert.False(t, eq(a, c))
	assert.False(t, eq(b, c))
	assert.False(t, eq(a, d))
}

func TestHash(t *testing.T) {}
