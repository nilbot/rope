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

func TestValidRopes(t *testing.T) {
	for id, tc := range validCases {
		if tc.in.valid() != tc.want {
			t.Errorf("valid::%d failed, expected %v", id, tc.want)
		}
	}
}
