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

func (r *ropeNode) nonLeaf() bool {
	if r == nil {
		return false
	}
	return (r.left != nil || r.right != nil) &&
		r.length == r.left.length+r.right.length
}

func (r *ropeNode) valid() bool {
	return r == nil || r.isLeaf() || r.nonLeaf()
}
