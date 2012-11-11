package llrb

func (t *Tree) Search(item Item) Item {

	x := t.root

	for x != nil {
		switch {
		case t.cmp(item, x.Item):
			x = x.left
		case t.cmp(x.Item, item):
			x = x.right
		default:
			return x.Item
		}
	}

	return nil

}

func (t *Tree) Max() Item {
	x := t.root

	for x.right != nil {
		x = x.right
	}

	return x.Item
}

func (t *Tree) Min() Item {
	x := t.root

	for x.left != nil {
		x = x.left
	}

	return x.Item
}
