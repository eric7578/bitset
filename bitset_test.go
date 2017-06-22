package bitset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitSetSet(t *testing.T) {
	bitSet := BitSet{0}
	bitSet.Set(0, 1, 2)
	assert.Equal(t, 7, bitSet.Value())
}

func TestBitSetClear(t *testing.T) {
	bitSet := BitSet{7}
	bitSet.Clear(1, 2)
	assert.Equal(t, 1, bitSet.Value())
}

func TestBitSetToggle(t *testing.T) {
	bitSet := BitSet{7}
	bitSet.Toggle(0, 2)
	assert.Equal(t, 2, bitSet.Value())
}

func TestBitSetIsSet(t *testing.T) {
	bitSet := BitSet{2}
	assert.False(t, bitSet.IsSet(0))
	assert.True(t, bitSet.IsSet(1))
	assert.False(t, bitSet.IsSet(2))
	assert.False(t, bitSet.IsSet(0, 1, 2))

	bitSet = BitSet{7}
	assert.True(t, bitSet.IsSet(0, 1, 2))
}
