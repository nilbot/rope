package rope

import "testing"

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

var ropeConcatCases = []struct {
	in1  Rope
	in2  Rope
	want string
}{
	{nilRope, nilRope, ""},
	{okRope, nilRope, "ok"},
	{okRope, strRope, "okstr"},
	{commaRope, okRope, ", ok"},
}

func TestRopesConcat(t *testing.T) {
	for id, tc := range ropeConcatCases {
		if tc.in1.concat(tc.in2).String() != tc.want {
			t.Errorf("concat::%d failed, expected %q, got %q",
				id, tc.want, tc.in1.concat(tc.in2))
		}
	}
}

var ropeIndexCases = []struct {
	in   Rope
	t    int
	want rune
}{
	{okRope, 0, 'o'},
	{okRope, 1, 'k'},
	{strRope, 2, 'r'},
}

func TestRopeIndex(t *testing.T) {
	for id, tc := range ropeIndexCases {
		if tc.in.Index(tc.t) != tc.want {
			t.Errorf("index::%d failed, expected %q, got %q",
				id, tc.want, tc.in.Index(tc.t))
		}
	}
}

// ---- test objects ---- //

var nilRope *ropeNode
var okRope = &ropeNode{
	left:   nil,
	right:  nil,
	length: 2,
	data:   "ok",
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

var commaRope = &ropeNode{
	2,
	nil,
	nil,
	", ",
}

var totes = "totes_"
var safe = "safe_"
var and = "and_"
var efficient = "efficient"

var tsae = &ropeNode{
	0,
	&ropeNode{
		11,
		&ropeNode{
			6,
			nil,
			nil,
			totes,
		},
		&ropeNode{
			5,
			nil,
			nil,
			safe,
		},
		"",
	},
	&ropeNode{
		0,
		&ropeNode{
			4,
			nil,
			nil,
			and,
		},
		&ropeNode{
			9,
			nil,
			nil,
			efficient,
		},
		"",
	},
	"",
}
