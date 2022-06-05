package greedy

import (
	"testing"
)

func Test_BuildHuffCodesByFreq(t *testing.T) {
	f := map[rune]int{
		'a': 45,
		'b': 13,
		'c': 12,
		'f': 5,
		'e': 9,
		'd': 16,
	}

	exp := map[rune]uint64{
		'a': 0b0,
		'b': 0b101,
		'c': 0b100,
		'f': 0b1100,
		'e': 0b1101,
		'd': 0b111,
	}

	r := BuildHuffCodesByFreq(f)

	for s, c := range exp {
		if r[s].Bits != c {
			t.Fatalf("wrong bits for %s: got %b; expected %b", string(s), r[s].Bits, c)
		}
	}
}
