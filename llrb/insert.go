package llrb

func (t *Tree) Insert(item Item) {
	t.root = insert(t.root, t.cmp, item)
	t.root.color = BLACK
}

func newNode(item Item) *Node {
	node := Node{Item: item, color: RED}
	return &node
}

func insert(h *Node, cmp CmpFunc, item Item) *Node {
	if h == nil {
		return newNode(item)
	}

	if isRed(h.left) && isRed(h.right) {
		colorFlip(h)
	}

	switch {
	case cmp(item, h.Item):
		h.left = insert(h.left, cmp, item)
	case cmp(h.Item, item):
		h.right = insert(h.right, cmp, item)
	default:
		h.Item = item
	}

	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}

	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	
	return h
}
