package bitset

// BitSet can hold up to 32 binary flags encoded in int32
type BitSet int32

// Reset unsets all flags to false
func (s *BitSet) Reset() {
	*s = 0
}

// Any returns true if any of flags is true
func (s BitSet) Any() bool {
	return s != 0
}

// All returns true if all of flags are true
func (s BitSet) All() bool {
	return s == -1
}

// Set flags or unflags a bit
func (s *BitSet) Set(n int, b bool) {
	if 0 <= n && n < 32 {
		bit := 1 << n
		if b == true {
			*s |= BitSet(bit)
		} else {
			*s &= BitSet(^bit)
		}
	}
}

// Test returns true if the specified bit is flagged
func (s BitSet) Test(n int) bool {
	if 0 <= n && n < 32 {
		bit := BitSet(1 << n)
		if s&bit != 0 {
			return true
		}
	}
	return false
}
