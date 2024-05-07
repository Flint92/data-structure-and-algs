package tree

import "util"

// AVLTree 任何节点的左子树和右子树的高度差不能超过1
type AVLTree struct {
	root *AVLTreeNode
}

type AVLTreeNode struct {
	val    int
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (avlTree *AVLTree) Search(val int) bool {
	return searchAVLTreeNode(avlTree.root, val)
}

func (avlTree *AVLTree) Insert(val int) bool {
	exists := avlTree.Search(val)
	if exists {
		return false
	}
	avlTree.root = insertAVLTreeNode(avlTree.root, val)
	return true
}

func (avlTree *AVLTree) Delete(val int) bool {
	exists := avlTree.Search(val)
	if exists {
		avlTree.root = delAVLTreeNode(avlTree.root, val)
		return true
	}
	return false
}

// PreOrderTraversal Morris实现前序遍历
// https://wansuanfa.com/index.php/1336
func (avlTree *AVLTree) PreOrderTraversal() []int {
	preList := make([]int, 0)
	cur := avlTree.root
	for cur != nil {
		if cur.left == nil {
			preList = append(preList, cur.val)
			cur = cur.right
		} else {
			pre := cur.left
			for pre.right != nil && pre.right != cur {
				pre = pre.right
			}
			if pre.right == nil {
				preList = append(preList, cur.val) // 第一次访问根节点
				pre.right = cur
				cur = cur.left
			} else {
				pre.right = nil // 访问左子树后第二次访问根节点，这时候需要将树还原
				cur = cur.right
			}
		}
	}
	return preList
}

// InOrderTraversal Morris实现中序遍历
func (avlTree *AVLTree) InOrderTraversal() []int {
	inList := make([]int, 0)
	cur := avlTree.root
	for cur != nil {
		if cur.left == nil {
			inList = append(inList, cur.val)
			cur = cur.right
		} else {
			pre := cur.left
			for pre.right != nil && pre.right != cur {
				pre = pre.right
			}
			if pre.right == nil {
				pre.right = cur
				cur = cur.left
			} else {
				inList = append(inList, cur.val)
				pre.right = nil
				cur = cur.right
			}
		}
	}
	return inList
}

// PostOrderTraversal Morris实现后序遍历
func (avlTree *AVLTree) PostOrderTraversal() []int {
	postList := make([]int, 0)
	cur := avlTree.root
	for cur != nil {
		if cur.left == nil {
			cur = cur.right
		} else {
			pre := cur.left
			for pre.right != nil && pre.right != cur {
				pre = pre.right
			}
			if pre.right == nil {
				pre.right = cur
				cur = cur.left
			} else {
				pre.right = nil
				postList = append(postList, reverseTraversal(cur.left)...)
				cur = cur.right
			}
		}
	}
	postList = append(postList, reverseTraversal(avlTree.root)...)
	return postList
}

func reverseTraversal(node *AVLTreeNode) []int {
	r := make([]int, 0)
	for node != nil {
		r = append([]int{node.val}, r...)
		node = node.right
	}
	return r
}

func delAVLTreeNode(node *AVLTreeNode, val int) *AVLTreeNode {
	if node == nil {
		return nil
	}
	if node.val == val {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}

		post := postNode(node)

		// 移形换位，把后继节点的值赋值到node节点，删除后继节点
		node.val = post.val

		node.right = delAVLTreeNode(node.right, post.val)
	} else if val < node.val {
		node.left = delAVLTreeNode(node.left, val)
	} else {
		node.right = delAVLTreeNode(node.right, val)
	}

	return balanceAVLTreeNode(node)
}

func postNode(node *AVLTreeNode) *AVLTreeNode {
	rightSmallest := node.right
	for rightSmallest != nil && rightSmallest.left != nil {
		rightSmallest = rightSmallest.left
	}
	return rightSmallest
}

func insertAVLTreeNode(node *AVLTreeNode, val int) *AVLTreeNode {
	if node == nil {
		node = &AVLTreeNode{val: val}
	} else if val < node.val {
		node.left = insertAVLTreeNode(node.left, val)
	} else if val > node.val {
		node.right = insertAVLTreeNode(node.right, val)
	}
	return balanceAVLTreeNode(node)
}

// 左左类型（LL）：直接对不平衡节点右旋。
// 左右类型（LR）：先对不平衡节点的左子节点左旋，然后对不平衡节点右旋。
// 右右类型（RR）：直接对不平衡节点左旋。
// 右左类型（RL）：先对不平衡节点的右子节点右旋，然后对不平衡节点左旋。
func balanceAVLTreeNode(node *AVLTreeNode) *AVLTreeNode {
	if node == nil {
		return nil
	}

	if getNodeHeight(node.left)-getNodeHeight(node.right) > 1 {
		// 左侧失衡
		if getNodeHeight(node.left.left) >= getNodeHeight(node.left.right) {
			node = balanceLL(node) // LL
		} else {
			node = balanceLR(node) // LR
		}

	} else if getNodeHeight(node.right)-getNodeHeight(node.left) > 1 {
		// 右侧失衡
		if getNodeHeight(node.right.right) >= getNodeHeight(node.right.left) {
			node = balanceRR(node) // RR
		} else {
			node = balanceRL(node) // RL
		}
	}

	resetHeight(node)
	return node
}

func balanceLL(node *AVLTreeNode) *AVLTreeNode {
	left := node.left
	node.left = left.right
	left.right = node
	resetHeight(node)
	resetHeight(left)
	return left
}

func balanceRR(node *AVLTreeNode) *AVLTreeNode {
	right := node.right
	node.right = right.left
	right.left = node
	resetHeight(node)
	resetHeight(right)
	return right
}

func balanceLR(node *AVLTreeNode) *AVLTreeNode {
	node.left = balanceRR(node.left)
	return balanceLL(node)
}

func balanceRL(node *AVLTreeNode) *AVLTreeNode {
	node.right = balanceLL(node.right)
	return balanceRR(node)
}

func resetHeight(node *AVLTreeNode) {
	node.height = util.MaxInt(getNodeHeight(node.left), getNodeHeight(node.right)) + 1
}

func getNodeHeight(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func searchAVLTreeNode(node *AVLTreeNode, val int) bool {
	if node == nil {
		return false
	}

	if node.val == val {
		return true
	} else if val < node.val {
		return searchAVLTreeNode(node.left, val)
	} else {
		return searchAVLTreeNode(node.right, val)
	}

}
