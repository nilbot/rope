package rope

import "fmt"

// special thanks to @francesc for `understanding of nil`, it reinforced my own
// exploration on the topic http://goo.gl/VfpQW4

// Rope interface
type Rope interface {
	fmt.Stringer
	isLeaf() bool
	nonLeaf() bool
	valid() bool
	Len() int
	concat(Rope) Rope
}

// types
type ropeNode struct {
	length int
	left   *ropeNode
	right  *ropeNode
	data   string
}

func (r *ropeNode) String() string {
	if r == nil {
		return ""
	}
	return "NOTIMPLEMENTED"
}

func (r *ropeNode) isLeaf() bool {
	// rely on shortcut to not panic
	return r != nil && r.left == nil && r.right == nil
}

// Len returns the length of the rope
func (r *ropeNode) Len() int {
	if r == nil {
		return 0
	}
	return r.length
}

func (r *ropeNode) nonLeaf() bool {
	if r == nil {
		return false
	}
	return r.length == r.left.Len()+r.right.Len()
}

func (r *ropeNode) valid() bool {
	return r == nil || r.isLeaf() || r.nonLeaf()
}

// WIP - unbalanced tree + no compression etc.
// [1,nil,nil,'s'] + [2,[1,nil,nil,'t'],[1,nil,nil,'r'],ignored] = (unbalanced)
// [3,[1,nil,nil,'s'],[2,[1,nil,nil,'t'],[1,nil,nil,'r'],nil],ignored]
func (r *ropeNode) concat(other Rope) Rope {
	if r == nil {
		return other
	}
	if other == nil {
		return r
	}
	newlen := r.Len() + other.Len()
	return &ropeNode{newlen, r, other.(*ropeNode), ""}
}
