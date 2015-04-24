package numbers

import (
	"errors"
	"github.com/bep/cast"
)

// Seq creates a sequence of integers.
// It's named and used as GNU's seq.
// Examples:
// 3 => 1, 2, 3
// 1 2 4 => 1, 3
// -3 => -1, -2, -3
// 1 4 => 1, 2, 3, 4
// 1 -2 => 1, 0, -1, -2
func Seq(args interface{}) ([]int, error) {

	intArgs := cast.ToIntSlice(args)

	if len(intArgs) < 1 || len(intArgs) > 3 {
		return nil, errors.New("invalid number of arguments")
	}

	var inc = 1
	var last int
	var first = intArgs[0]

	if len(intArgs) == 1 {
		last = first
		if last == 0 {
			return []int{}, nil
		} else if last > 0 {
			first = 1
		} else {
			first = -1
			inc = -1
		}
	} else if len(intArgs) == 2 {
		last = intArgs[1]
		if last < first {
			inc = -1
		}
	} else {
		inc = intArgs[1]
		last = intArgs[2]
		if inc == 0 {
			return nil, errors.New("<increment> must not be equal to 0")
		}
		if first < last && inc < 0 {
			return nil, errors.New("<increment> must be greater than 0")
		}
		if first > last && inc > 0 {
			return nil, errors.New("<increment> must be less than 0")
		}
	}

	size := int(((last - first) / inc) + 1)

	seq := make([]int, size)
	val := first
	for i := 0; ; i++ {
		seq[i] = val
		val += inc
		if (inc < 0 && val < last) || (inc > 0 && val > last) {
			break
		}
	}

	return seq, nil
}
