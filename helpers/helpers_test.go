package helpers

import (
	"testing"
)

func TestMaxInt(t *testing.T) {
	for i, this := range []struct {
		i1       int
		i2       int
		expected int
	}{
		{1, 1, 1},
		{1, 2, 2},
		{-1, 1, 1},
		{-10, -4, -4},
	} {
		got := MaxInt(this.i1, this.i2)
		if got != this.expected {
			t.Errorf("[%d] MaxInt error for %d/%d: Got %d, expected %d", i, this.i1, this.i2, got, this.expected)
		}
	}
}

func TestNumDigits(t *testing.T) {
	for i, this := range []struct {
		in       int
		expected int
	}{
		{1, 1},
		{0, 1},
		{10, 2},
		{9999, 4},
		{-1, 1},
		{-100, 3},
	} {
		got := NumDigits(this.in)

		if got != this.expected {
			t.Errorf("[%d] NumDigits error for %d: Got %d, expected %d", i, this.in, got, this.expected)
		}
	}
}
