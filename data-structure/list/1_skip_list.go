package list

import (
	"fmt"
	"math/rand"
)

const maxLevel int = 2 << 3

// SkipList https://www.baeldung.com/cs/skip-lists
type SkipList struct {
	head         *SkipListNode
	currentLevel int
}

type SkipListNode struct {
	val  int
	next *SkipListNode
	down *SkipListNode
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:         &SkipListNode{val: -1, next: nil},
		currentLevel: 1,
	}
}

func (sl *SkipList) Add(num int) {
	level := randomLevel()

	// 记录待插入节点的前一个节点
	preNodes := make([]*SkipListNode, level)

	// 第一步：如果跳表层数比较少，在上面添加，层数至少为 level
	if sl.currentLevel < level {
		beforeHead := sl.head

		// 更新head节点
		sl.head = &SkipListNode{val: -1, next: nil}
		curHead := sl.head

		// 在上面添加每层的头节点
		for i := sl.currentLevel; i < level-1; i++ {
			node := &SkipListNode{val: -1, next: nil}
			curHead.down = node
			curHead = node
		}

		// 最后创建的链表头节点和之前的头节点连在一起
		curHead.down = beforeHead
	}

	// 第二步：从上往下查找每层待插入节点的前一个节点
	pre := sl.head
	// 上层不需要插入的跳过。
	for i := sl.currentLevel - 1; i >= level; i-- {
		pre = pre.down
	}
	// 从当前层往下每层都要插入该节点，找出每层待插入节点的前一个节点。
	for i := level - 1; i >= 0; i-- {
		for pre.next != nil && pre.next.val < num {
			pre = pre.next
		}
		preNodes[i] = pre // 记录前一个节点。
		pre = pre.down
	}

	// 第三步：节点插入，插入的时候不光有 next 指针，而且还有 down 指针
	var topNode *SkipListNode
	for i := level - 1; i >= 0; i-- {
		node := &SkipListNode{val: num, next: preNodes[i].next}
		preNodes[i].next = node
		// 上下也要链接
		if topNode != nil {
			topNode.down = node
		}
		topNode = node
	}

	// 更新跳表的层级，用来记录当前跳表的层级
	if level > sl.currentLevel {
		sl.currentLevel = level
	}
}

func (sl *SkipList) Remove(num int) bool {
	// 删除链表和插入链表类似，都是需要先找到插入或删除链表的前一个节点。
	topIndex := -1 // 从当前层开始往下每层都要删除。
	// 查找待删除节点的前一个节点，从上面一层开始查找。
	pre := sl.head
	for i := sl.currentLevel - 1; i >= 0; i-- {
		for pre.next != nil && pre.next.val < num {
			pre = pre.next
		}
		// 如果找到就终止查找，表示在当前层以及他下面的所有层都要删除该节点
		if pre.next != nil && pre.next.val == num {
			topIndex = i
			break
		}
		if pre.down == nil {
			// 如果跳表中没有要删除的节点，返回 false 。
			return false
		}
		pre = pre.down // 当前层没找到就往下一层继续查找。
	}

	if topIndex == -1 {
		// 如果跳表中没找到要删除的节点，返回 false 。
		return false
	}

	// 从 topIndex 层开始，他下面的每一层都要删除。
	for i := topIndex; i >= 0; i-- {
		if pre == nil {
			break
		}
		if pre.next != nil {
			pre.next = pre.next.next
		}
		pre = pre.down // 继续下一层的删除。
		if pre != nil {
			// 找到待删除节点的前一个节点。
			for pre.next != nil && pre.next.val != num {
				pre = pre.next
			}
		}
	}
	// 如果上面一层的节点被删除完了，要更新 curLevelCount 的值 ，还要更新 head节点。
	curr := sl.head
	for sl.currentLevel > 1 && curr.next == nil {
		curr = curr.down
		sl.head = curr
		sl.currentLevel--
	}
	return true
}

func (sl *SkipList) Search(target int) bool {
	pre := sl.head
	for pre != nil {
		// 如果当前节点值小于 target ，需要到右边查找，如果右边没有节点就到下边查找。
		if pre.val < target {
			if pre.next == nil {
				// 右边没有节点，到下边查找
				pre = pre.down
			} else if pre.next.val > target {
				pre = pre.down
			} else {
				pre = pre.next
			}
		} else if pre.val == target {
			// 如果找到直接返回。
			return true
		} else {
			// 如果当前节点值大于 target ，说明没有，直接返回 false
			return false
		}
	}
	return false
}

func (sl *SkipList) Print() {
	cur := sl.head
	for cur != nil {
		tmp := cur.next
		for tmp != nil {
			if tmp.next != nil {
				fmt.Print(tmp.val, "->")
			} else {
				fmt.Print(tmp.val)
			}

			tmp = tmp.next
		}
		fmt.Println()
		cur = cur.down
	}
}

func randomLevel() int {
	level := 1
	for i := 1; i < maxLevel; i++ {
		if rand.Int31()%2 == 1 {
			level++
		}
	}
	return level
}
