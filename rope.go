package rope

import "fmt"
import "unicode/utf8"

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
	Index(int) rune
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
	var stringSlice []string
	preorderTraversal(r, &stringSlice)
	concat := ""
	for _, v := range stringSlice {
		// just for testing. this is wronga
		// TODO get this right
		concat = concat + v
	}
	return concat
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

// Index returns the rune at index
func (r *ropeNode) Index(idx int) rune {
	if idx < 0 || !r.valid() {
		return 0
	}
	// if idx is greater than left.length, idx = idx - left.length
	if r.isLeaf() {
		data := r.data
		for curr := 0; curr < idx; curr++ {
			_, s := utf8.DecodeRuneInString(data)
			data = data[s:]
		}
		ru, _ := utf8.DecodeRuneInString(data)
		return ru
	}
	if idx < r.left.Len() {
		return r.left.Index(idx)
	}
	return r.right.Index(idx - r.left.Len())
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
	newlen := r.Len() + other.Len() // TODO handle overflow danger
	return &ropeNode{newlen, r, other.(*ropeNode), ""}
}

func preorderTraversal(r *ropeNode, inputSlice *[]string) {
	if r == nil {
		return
	}

	preorderTraversal(r.left, inputSlice)
	if r.isLeaf() {
		*inputSlice = append(*inputSlice, r.data)
		return
	}
	preorderTraversal(r.right, inputSlice)
}
