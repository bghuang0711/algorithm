package bubblesort

func BubbleSort(elements []int) {
	n := len(elements)
	for i := n - 1; i >= 0; i-- { // 为第i个位置冒泡，[i+1, n)已经保持顺序
		for j := 0; j < i; j++ {
			if elements[j] > elements[j+1] {
				elements[j], elements[j+1] = elements[j+1], elements[j]
			}
		}
	}
}
