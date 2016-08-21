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

var lenConcatCases = []struct {
	in1  Rope
	in2  Rope
	want int
}{
	{nilRope, nilRope, 0},
	{okRope, strRope, 5},
}

func TestRopesConcatLen(t *testing.T) {
	for id, tc := range lenConcatCases {
		if tc.in1.concat(tc.in2).Len() != tc.want {
			t.Errorf("len::%d failed, expected %v, got %v",
				id, tc.want, tc.in1.concat(tc.in2).Len())
		}
	}
}

var strRope = &ropeNode{
	3,
	&ropeNode{
		1,
		nil,
		nil,
		"s",
	},
	&ropeNode{
		2,
		&ropeNode{
			1,
			nil,
			nil,
			"t",
		},
		&ropeNode{
			1,
			nil,
			nil,
			"r",
		},
		"",
	},
	"",
}
