package numbers

import (
	"reflect"
	"testing"
)

func TestSeq(t *testing.T) {
	for i, this := range []struct {
		in     interface{}
		expect interface{}
	}{
		{[]int{-2, 5}, []int{-2, -1, 0, 1, 2, 3, 4, 5}},
		{[]interface{}{1, 2, 4}, []int{1, 3}},
		{[]interface{}{"1", "2", "4"}, []int{1, 3}},
		{[]string{"1", "2", "4"}, []int{1, 3}},
		{[]interface{}{1}, []int{1}},
		{[]interface{}{3}, []int{1, 2, 3}},
		{[]float64{3.2}, []int{1, 2, 3}},
		{[]interface{}{0}, []int{}},
		{[]interface{}{-1}, []int{-1}},
		{[]interface{}{-3}, []int{-1, -2, -3}},
		{[]interface{}{3, -2}, []int{3, 2, 1, 0, -1, -2}},
		{[]interface{}{6, -2, 2}, []int{6, 4, 2}},
		{[]interface{}{1, 0, 2}, false},
		{[]interface{}{1, -1, 2}, false},
		{[]interface{}{2, 1, 1}, false},
		{[]interface{}{2, 1, 1, 1}, false},
		{[]interface{}{}, false},
		{[]interface{}{t}, false},
		{nil, false},
	} {

		result, err := Seq(this.in)

		if b, ok := this.expect.(bool); ok && !b {
			if err == nil {
				t.Errorf("[%d] Seq didn't return an expected error", i)
			}
		} else {
			if err != nil {
				t.Errorf("[%d] failed: %s", i, err)
				continue
			}
			if !reflect.DeepEqual(result, this.expect) {
				t.Errorf("[%d] Seq got %v but expected %v", i, result, this.expect)
			}
		}
	}
}
