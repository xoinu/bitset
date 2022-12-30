package bitset

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	var bs BitSet

	if bs.Any() == true {
		t.Fatal("Any() must be false")
	}

	if bs.All() == true {
		t.Fatal("All() must be false")
	}

	for i := 0; i < 32; i++ {
		bs.Set(i, true)
		if bs.Test(i) == false {
			t.Fatal("error")
		}
		if bs.Any() == false {
			t.Fatal("Any() must be true")
		}
		if i < 31 && bs.All() == true {
			t.Fatal("All() must be false")
		}
	}

	if bs.All() == false {
		t.Fatal("All() must be true")
	}

	if bs != -1 {
		t.Fatal("bs must be -1")
	}

	for i := 0; i < 32; i++ {
		bs.Set(i, false)
		if bs.Test(i) == true {
			t.Fatal("error")
		}
		if i < 31 {
			if bs.Any() == false {
				t.Fatal("Any() must be true")
			}
		} else {
			if bs.Any() == true {
				t.Fatal("Any() must be false")
			}
		}
		if bs.All() == true {
			t.Fatal("All() must be false")
		}
	}

	if bs != 0 {
		t.Fatal("bs must be 0")
	}

	bs = -1
	bs.Reset()
	if bs != 0 {
		t.Fatal("bs must be 0 after Reset()")
	}
}
