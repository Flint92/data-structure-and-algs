package tree

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

// RBTree
// 节点是红色或黑色。
// 根节点是黑色。
// 所有叶子都是黑色。（叶子是nil节点）
// 从根节点到所有的叶子节点的路径上不能有两个连续的红色节点。
// 从任一节点到其每个叶子的路径都包含相同数目的黑色节点。
type RBTree struct {
	root *RBTreeNode
}

type RBTreeNode struct {
	val    int
	left   *RBTreeNode
	right  *RBTreeNode
	parent *RBTreeNode
	color  Color
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

// Search same to bst
func (rbTree *RBTree) Search(val int) bool {
	return searchRBNode(rbTree.root, val) != nil
}

// Delete from java TreeMap
func (rbTree *RBTree) Delete(val int) bool {
	p := searchRBNode(rbTree.root, val)
	if p == nil {
		return false
	}

	// has 2 children
	if p.left != nil && p.right != nil {
		s := successor(p)
		// s的值赋给p之后，只需要删除s即可
		p.val = s.val
		p = s // p替换为s
	}

	// replacement 是 p 的一个之节点，p 最多只有一个子节点
	var replacement *RBTreeNode
	if p.left != nil {
		replacement = p.left
	} else {
		replacement = p.right
	}

	if replacement != nil { //p不是叶子节点
		replacement.parent = p.parent
		if p.parent == nil {
			rbTree.root = replacement
		} else if p.parent.left == p {
			p.parent.left = replacement
		} else {
			p.parent.right = replacement
		}

		p.left = nil
		p.right = nil
		p.parent = nil

		// 如果删除的是黑色节点才调整，replacement节点必定是红色的
		if p.color == BLACK {
			rbTree.fixAfterDeletion(replacement)
		}

	} else if p.parent == nil { // p是根节点
		rbTree.root = nil
	} else { // p是叶子节点
		if p.color == BLACK {
			rbTree.fixAfterDeletion(p)
		}
		// 删除p
		if p.parent != nil {
			if p == p.parent.left {
				p.parent.left = nil
			} else if p == p.parent.right {
				p.parent.right = nil
			}
			p.parent = nil
		}
	}

	return true
}

// Insert from java TreeMap
func (rbTree *RBTree) Insert(val int) bool {
	exists := rbTree.Search(val)
	if exists {
		return false
	}

	if rbTree.root == nil {
		rbTree.root = &RBTreeNode{val: val, color: BLACK}
		return true
	}

	parent := rbTree.root
	var node *RBTreeNode
	for parent != nil {
		if parent.val < val {
			if parent.right == nil {
				node = &RBTreeNode{val: val, parent: parent, color: RED}
				parent.right = node
				break
			}
			parent = parent.right
		} else {
			if parent.left == nil {
				node = &RBTreeNode{val: val, parent: parent, color: RED}
				parent.left = node
				break
			}
			parent = parent.left
		}
	}

	// 只有父节点不为空且是红色才会操作
	for node != nil && node != rbTree.root && colorOf(parentOf(node)) == RED {
		if parentOf(node) == leftOf(parentOf(parentOf(node))) { // 父节点是爷爷的左子节点
			uncle := rightOf(parentOf(parentOf(node)))
			if colorOf(uncle) == RED {
				setColor(parentOf(node), BLACK)         // 父节点涂黑
				setColor(uncle, BLACK)                  // 叔叔节点涂黑
				setColor(parentOf(parentOf(node)), RED) // 爷爷节点涂红
				node = parentOf(parentOf(node))         // 继续向上判断爷爷节点
			} else {
				if node == rightOf(parentOf(node)) {
					node = parentOf(node)   // node节点指向父节点
					rbTree.rotateLeft(node) // 对于父节点进行左旋
				}
				setColor(parentOf(node), BLACK)              // 父节点涂黑。
				setColor(parentOf(parentOf(node)), RED)      // 爷爷节点涂红。
				rbTree.rotateRight(parentOf(parentOf(node))) // 对爷爷节点右旋。
			}
		} else { // 父节点是爷爷的右子节点
			uncle := leftOf(parentOf(parentOf(node)))
			if colorOf(uncle) == RED {
				setColor(parentOf(node), BLACK)
				setColor(uncle, BLACK)
				setColor(parentOf(parentOf(node)), RED)
				node = parentOf(parentOf(node))
			} else {
				if node == leftOf(parentOf(node)) {
					node = parentOf(node)
					rbTree.rotateRight(node)
				}
				setColor(parentOf(node), BLACK)
				setColor(parentOf(parentOf(node)), RED)
				rbTree.rotateLeft(parentOf(parentOf(node)))
			}
		}
	}

	rbTree.root.color = BLACK
	rbTree.root.parent = nil

	return true
}

func (rbTree *RBTree) rotateLeft(p *RBTreeNode) {
	if p != nil {
		r := p.right
		p.right = r.left
		if r.left != nil {
			r.left.parent = p
		}
		r.parent = p.parent
		if p.parent == nil {
			rbTree.root = r
		} else if p.parent.left == p {
			p.parent.left = r
		} else {
			p.parent.right = r
		}
		r.left = p
		p.parent = r
	}
}

func (rbTree *RBTree) rotateRight(p *RBTreeNode) {
	if p != nil {
		l := p.left
		p.left = l.right
		if l.right != nil {
			l.right.parent = p
		}
		l.parent = p.parent
		if p.parent == nil {
			rbTree.root = l
		} else if p.parent.right == p {
			p.parent.right = l
		} else {
			p.parent.left = l
		}
		l.right = p
		p.parent = l
	}
}

func (rbTree *RBTree) fixAfterDeletion(x *RBTreeNode) {
	for x != rbTree.root && colorOf(x) == BLACK {
		if x == leftOf(parentOf(x)) {
			sib := rightOf(parentOf(x))

			if colorOf(sib) == RED {
				setColor(sib, BLACK)
				setColor(parentOf(x), RED)
				rbTree.rotateLeft(parentOf(x))
				sib = rightOf(parentOf(x))
			}

			if colorOf(leftOf(sib)) == BLACK && colorOf(rightOf(sib)) == BLACK {
				setColor(sib, RED)
				x = parentOf(x)
			} else {
				if colorOf(rightOf(sib)) == BLACK {
					setColor(leftOf(sib), BLACK)
					setColor(sib, RED)
					rbTree.rotateRight(sib)
					sib = rightOf(parentOf(x))
				}
				setColor(sib, colorOf(parentOf(x)))
				setColor(parentOf(x), BLACK)
				setColor(rightOf(sib), BLACK)
				rbTree.rotateLeft(parentOf(x))
				x = rbTree.root
			}

		} else { // symmetric
			sib := leftOf(parentOf(x))

			if colorOf(sib) == RED {
				setColor(sib, BLACK)
				setColor(parentOf(x), RED)
				rbTree.rotateRight(parentOf(x))
				sib = leftOf(parentOf(x))
			}

			if colorOf(rightOf(sib)) == BLACK && colorOf(leftOf(sib)) == BLACK {
				setColor(sib, RED)
				x = parentOf(x)
			} else {
				if colorOf(leftOf(sib)) == BLACK {
					setColor(rightOf(sib), BLACK)
					setColor(sib, RED)
					rbTree.rotateLeft(sib)
					sib = leftOf(parentOf(x))
				}
				setColor(sib, colorOf(parentOf(x)))
				setColor(parentOf(x), BLACK)
				setColor(leftOf(sib), BLACK)
				rbTree.rotateRight(parentOf(x))
				x = rbTree.root
			}
		}
	}

	setColor(x, BLACK)
}

func searchRBNode(node *RBTreeNode, val int) *RBTreeNode {
	if node == nil {
		return nil
	}

	if node.val == val {
		return node
	} else if val < node.val {
		return searchRBNode(node.left, val)
	} else {
		return searchRBNode(node.right, val)
	}

}

func parentOf(p *RBTreeNode) *RBTreeNode {
	if p == nil {
		return nil
	}
	return p.parent
}

func rightOf(p *RBTreeNode) *RBTreeNode {
	if p == nil {
		return nil
	}
	return p.right
}

func leftOf(p *RBTreeNode) *RBTreeNode {
	if p == nil {
		return nil
	}
	return p.left
}

func colorOf(p *RBTreeNode) Color {
	if p == nil {
		return BLACK
	}
	return p.color
}

func setColor(p *RBTreeNode, color Color) {
	if p != nil {
		p.color = color
	}
}

func successor(t *RBTreeNode) *RBTreeNode {
	if t == nil {
		return nil
	}
	if t.right != nil {
		p := t.right
		for p.left != nil {
			p = p.left
		}
		return p
	} else {
		p := t.parent
		ch := t
		for p != nil && ch == p.right {
			ch = p
			p = p.parent
		}
		return p
	}
}
