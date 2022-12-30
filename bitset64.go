package bitset

// BitSet64 can hold up to 64 binary flags encoded in int64
type BitSet64 int64

// Reset unsets all flags to false
func (s *BitSet64) Reset() {
	*s = 0
}

// Any returns true if any of flags is true
func (s BitSet64) Any() bool {
	return s != 0
}

// All returns true if all of flags are true
func (s BitSet64) All() bool {
	return s == -1
}

// Set flags or unflags a bit
func (s *BitSet64) Set(n int, b bool) {
	if 0 <= n && n < 64 {
		bit := 1 << n
		if b == true {
			*s |= BitSet64(bit)
		} else {
			*s &= BitSet64(^bit)
		}
	}
}

// Test returns true if the specified bit is flagged
func (s BitSet64) Test(n int) bool {
	if 0 <= n && n < 64 {
		bit := BitSet64(1 << n)
		if s&bit != 0 {
			return true
		}
	}
	return false
}
