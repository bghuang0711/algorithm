package insertsort

func InsertSort(elements []int) {
	n := len(elements)

	for i := 1; i < n; i++ {
		cur := elements[i]
		for j := i - 1; j >= 0; j-- {
			if cur < elements[j] { // 如果待插入元素小于当前元素，当前元素后移一个位置
				elements[j+1] = elements[j]
			} else { // 否则，待插入元素插入当前位置的下一个位置，停止本轮插入
				elements[j+1] = cur
				break
			}
		}
	}
}
