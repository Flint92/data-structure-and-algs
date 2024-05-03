package tree

import (
	"fmt"
)

type SegmentTree struct {
	trees []int // 线段树数组，为了防止越界，线段树数组是原始数组的4倍
	lazy  []int // 懒标记数组，更新和查询优化
	n     int   // 原始数组的长度
}

func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	seg := &SegmentTree{
		n:     n,
		trees: make([]int, n<<2),
		lazy:  make([]int, n<<2),
	}

	seg.buildTree(nums, 0, 0, n-1)

	return seg
}

func (t *SegmentTree) QueryRange(left, right int) int {
	return t.queryRange(0, 0, t.n-1, left, right)
}

func (t *SegmentTree) UpdateRange(left, right, val int) {
	t.updateRange(0, 0, t.n-1, left, right, val)
}

func (t *SegmentTree) Print() {
	for i := 0; i < len(t.trees); i++ {
		fmt.Printf("trees[%d]=%d\t", i, t.trees[i])
	}
	fmt.Println()
}

func (t *SegmentTree) updateRange(treeNode, start, end, left, right, val int) {
	t.pushDown(treeNode, start, end)
	if start > right || end < left {
		return
	}
	if start >= left && end <= right {
		t.trees[treeNode] += (end - start + 1) * val
		if start != end {
			t.lazy[leftNode(treeNode)] += val
			t.lazy[rightNode(treeNode)] += val
		}
		return
	} else {
		mid := midNode(start, end)
		if left <= mid {
			t.updateRange(leftNode(treeNode), start, mid, left, right, val)
		}
		if right > mid {
			t.updateRange(rightNode(treeNode), mid+1, end, left, right, val)
		}
		t.pushUp(treeNode)
	}
}

func (t *SegmentTree) queryRange(treeNode, start, end, left, right int) int {
	t.pushDown(treeNode, start, end)
	if start > right || end < left {
		return 0
	}

	if start >= left && end <= right {
		return t.trees[treeNode]
	}
	sum := 0
	mid := midNode(start, end)
	if left <= mid {
		sum += t.queryRange(leftNode(treeNode), start, mid, left, right)
	}
	if right > mid {
		sum += t.queryRange(rightNode(treeNode), mid+1, end, left, right)
	}
	return sum
}

func (t *SegmentTree) buildTree(nums []int, treeNode, start, end int) {
	if start == end {
		t.trees[treeNode] = nums[start]
	} else {
		mid := midNode(start, end)
		left := leftNode(treeNode)
		right := rightNode(treeNode)

		t.buildTree(nums, left, start, mid)
		t.buildTree(nums, right, mid+1, end)
		t.pushUp(treeNode)
	}

}

func (t *SegmentTree) pushDown(treeNode, start, end int) {
	if t.lazy[treeNode] != 0 {
		t.trees[treeNode] += (end - start + 1) * t.lazy[treeNode]
		if start != end {
			t.lazy[leftNode(treeNode)] += t.lazy[treeNode]
			t.lazy[rightNode(treeNode)] += t.lazy[treeNode]
		}
		t.lazy[treeNode] = 0
	}
}

func (t *SegmentTree) pushUp(i int) {
	// 求和
	t.trees[i] = t.trees[leftNode(i)] + t.trees[rightNode(i)]
}

func leftNode(treeNode int) int {
	return treeNode<<1 | 1
}

func rightNode(treeNode int) int {
	return (treeNode + 1) << 1
}

func midNode(start, end int) int {
	return start + (end-start)>>1
}
