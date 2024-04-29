package tree

import (
	"fmt"
	"sort"
	"strings"
)

// HuffmanTree https://en.wikipedia.org/wiki/Huffman_coding
type HuffmanTree struct {
	root *HuffmanTreeNode
}

type HuffmanTreeNode struct {
	val    string
	weight float64 // 权重
	left   *HuffmanTreeNode
	right  *HuffmanTreeNode
}

func NewHuffmanTree(freqs map[string]float64) *HuffmanTree {
	var nodes []*HuffmanTreeNode
	for word, freq := range freqs {
		nodes = append(nodes, &HuffmanTreeNode{val: word, weight: freq})
	}

	for len(nodes) > 1 {
		// 这里可以使用最小堆进行优化
		sortByWeight(nodes)

		first := nodes[0]
		second := nodes[1]

		nodes = nodes[2:]

		parent := &HuffmanTreeNode{
			weight: first.weight + second.weight,
			left:   first,
			right:  second,
		}

		nodes = append(nodes, parent)
	}

	return &HuffmanTree{root: nodes[0]}
}

func (hTree *HuffmanTree) Encode(val string) string {
	encodeMap := hTree.toEncodeMap()
	encodeStrs := make([]string, 0, len(val))
	for _, ch := range val {
		encodeStrs = append(encodeStrs, encodeMap[string(ch)])
	}
	return strings.Join(encodeStrs, "")
}

func (hTree *HuffmanTree) Decode(val string) (string, error) {
	node := hTree.root
	var lastNode *HuffmanTreeNode
	results := make([]string, 0)
	for _, ch := range val {
		if string(ch) == "1" {
			node = node.right
		} else if string(ch) == "0" {
			node = node.left
		} else {
			return "", fmt.Errorf("unknown char %v", string(ch))
		}

		if isLeaf(node) {
			results = append(results, node.val)
			lastNode = node
			node = hTree.root
		} else {
			lastNode = node
		}
	}

	if !isLeaf(lastNode) {
		return "", fmt.Errorf("invalid encode str %v", val)
	}

	return strings.Join(results, ""), nil
}

func (hTree *HuffmanTree) toEncodeMap() map[string]string {
	encodeMap := make(map[string]string)
	dfsToBuildEncodeMap(hTree.root, "", encodeMap)
	return encodeMap
}

func dfsToBuildEncodeMap(node *HuffmanTreeNode, bitStr string, encodeMap map[string]string) {
	if node == nil {
		return
	}

	if isLeaf(node) {
		if len(bitStr) > 0 {
			encodeMap[node.val] = bitStr
		}
	} else {
		dfsToBuildEncodeMap(node.left, bitStr+"0", encodeMap)
		dfsToBuildEncodeMap(node.right, bitStr+"1", encodeMap)
	}
}

func isLeaf(node *HuffmanTreeNode) bool {
	return node.left == nil && node.right == nil
}

func sortByWeight(nodes []*HuffmanTreeNode) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].weight < nodes[j].weight
	})
}
