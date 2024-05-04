package heap

type MinHeap struct {
	data []int
	size int
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{data: make([]int, size)}
}

func (mh *MinHeap) Add(val int) {
	if mh.size == len(mh.data) {
		mh.data = append(mh.data, 0)
	}
	if mh.size == 0 {
		mh.data[0] = val
	} else {
		mh.siftUp(val)
	}
	mh.size++
}

func (mh *MinHeap) Peek() (int, bool) {
	if mh.size == 0 {
		return -1, false
	}

	return mh.data[0], true
}

func (mh *MinHeap) Remove() (int, bool) {
	if mh.size == 0 {
		return -1, false
	}
	mh.size--
	result := mh.data[0]        // 获取堆顶元素
	x := mh.data[mh.size]       // 获取原数组最后一个元素
	mh.data = mh.data[:mh.size] // 将原数组最后一个元素移除
	if mh.size != 0 {
		mh.siftDown(x)
	}
	return result, true
}

func (mh *MinHeap) Size() int {
	return mh.size
}

func (mh *MinHeap) siftUp(val int) {
	s := mh.size

	for s > 0 {
		parent := (s - 1) >> 1
		parentVal := mh.data[parent]
		if val > parentVal {
			break
		}
		mh.data[s] = parentVal
		s = parent
	}

	mh.data[s] = val
}

func (mh *MinHeap) siftDown(x int) {
	half := x >> 1
	index := 0 // 从根节点开始往下调整

	for index < half {
		minChildIdx := index<<1 | 1 // 左子节点
		minChild := mh.data[minChildIdx]

		right := minChildIdx + 1

		if right < mh.size { // 如果有右子节点，肯定有左子节点
			if mh.data[right] < minChild {
				minChildIdx = right
				minChild = mh.data[minChildIdx]
			}
		}

		if x <= minChild {
			break
		}

		mh.data[index] = minChild
		index = minChildIdx
	}

	mh.data[index] = x
}
