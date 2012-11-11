package llrb

type Tree struct {
	root *Node
	cmp  CmpFunc
}

const RED = true
const BLACK = false

type Color bool

type Item interface{}

type Node struct {
	Item
	left, right *Node
	color       Color
}

type CmpFunc func(a, b interface{}) bool

func New(cmpfunc CmpFunc) *Tree {
	tree := Tree{nil, cmpfunc}
	return &tree
}
