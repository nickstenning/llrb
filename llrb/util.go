package llrb

func rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	return x
}

func rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	return x
}

func colorFlip(h *Node) {
	h.color       = !h.color
	h.left.color  = !h.left.color
	h.right.color = !h.right.color
}

func isRed(x *Node) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}
