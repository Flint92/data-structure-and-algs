package tree

import "fmt"

type BST struct {
	root *BSTNode
}

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Search(val int) bool {
	return search(bst.root, val)
}

func (bst *BST) Insert(val int) bool {
	exists := bst.Search(val)
	if exists {
		return false
	}
	bst.root = insert(bst.root, val)
	return true
}

func (bst *BST) Delete(val int) bool {
	exists := bst.Search(val)
	if exists {
		del(bst.root, val)
		return true
	}
	return false
}

func (bst *BST) Print() {
	if bst.root == nil {
		fmt.Println("Empty Tree")
		return
	}
	inorderTraversal(bst.root)
}

func inorderTraversal(node *BSTNode) {
	if node == nil {
		return
	}
	inorderTraversal(node.left)
	fmt.Println(node.val)
	inorderTraversal(node.right)
}

// 查找节点node的前驱节点
// 前驱节点：对一棵二叉树进行中序遍历，遍历后的结果中，当前节点的前一个节点为该节点的前驱节点；
// 查找node节点的左子树的最大节点即可
func preNode(node *BSTNode) *BSTNode {
	leftBiggest := node.left
	for leftBiggest != nil && leftBiggest.right != nil {
		leftBiggest = leftBiggest.right
	}
	return leftBiggest
}

// 查找节点node的后继节点
// 后继节点：对一棵二叉树进行中序遍历，遍历后的结果中，当前节点的后一个节点为该节点的后继节点
// 查找node节点的右子树的最小节点即可
func postNode(node *BSTNode) *BSTNode {
	rightSmallest := node.right
	for rightSmallest != nil && rightSmallest.left != nil {
		rightSmallest = rightSmallest.left
	}
	return rightSmallest
}

func del(node *BSTNode, val int) *BSTNode {
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

		// 有两种方式。
		// 一种是让 p 的右子节点成为他前驱节点的右子节点。本例子中采用
		pre := preNode(node)
		pre.right = node.right
		return node.left

		// 还一种是让 p 的左子节点成为他后继节点的左子节点。
		//post := postNode(node)
		//post.left = node.left
		//return node.right
	} else if val < node.val {
		node.left = del(node.left, val)
		return node
	} else {
		node.right = del(node.right, val)
		return node
	}
}

func insert(node *BSTNode, val int) *BSTNode {
	if node == nil {
		return &BSTNode{val, nil, nil}
	}
	if val < node.val {
		node.left = insert(node.left, val)
	} else {
		node.right = insert(node.right, val)
	}
	return node
}

func search(node *BSTNode, val int) bool {
	if node == nil {
		return false
	}

	if node.val == val {
		return true
	} else if val < node.val {
		return search(node.left, val)
	} else {
		return search(node.right, val)
	}

}
