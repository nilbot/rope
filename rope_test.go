package rope

import "testing"

var nilRope *ropeNode
var okRope = &ropeNode{
	left:   nil,
	right:  nil,
	length: 2,
	data:   "ok",
}
var validCases = []struct {
	in   Rope
	want bool
}{
	{nilRope, true},
	{okRope, true},
}
var lenCases = []struct {
	in   Rope
	want int
}{
	{nilRope, 0},
	{okRope, 2},
}

func TestValidRopes(t *testing.T) {
	for id, tc := range validCases {
		if tc.in.valid() != tc.want {
			t.Errorf("valid::%d failed, expected %v", id, tc.want)
		}
	}
}

func TestRopeLen(t *testing.T) {
	for id, tc := range lenCases {
		if tc.in.Len() != tc.want {
			t.Errorf("len::%d failed, expected %v, got %v",
				id, tc.want, tc.in.Len())
		}
	}
}
