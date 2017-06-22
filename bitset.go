package bitset

// NewBitSet create BitSet with initial value
func NewBitSet(value int) *BitSet {
	return &BitSet{value}
}

/*
BitSet provide methods to operate the undertable value.
*/
type BitSet struct {
	value int
}

// Set will the assigned position as 1 to the undertable value
func (b *BitSet) Set(positions ...uint) {
	for _, position := range positions {
		b.value |= 1 << position
	}
}

// IsSet will check wheher all the assigned positions are 1
func (b *BitSet) IsSet(positions ...uint) bool {
	for _, position := range positions {
		isset := (b.value >> position) & 1
		if isset == 0 {
			return false
		}
	}
	return true
}

// Clear will clear the assigned positions as 0
func (b *BitSet) Clear(positions ...uint) {
	for _, position := range positions {
		b.value &= ^(1 << position)
	}
}

// Toggle will change the assigned positions from 1 to 0, or from 0 to 1
func (b *BitSet) Toggle(positions ...uint) {
	for _, position := range positions {
		b.value ^= 1 << position
	}
}

// Value will return the undertable value
func (b *BitSet) Value() int {
	return b.value
}
