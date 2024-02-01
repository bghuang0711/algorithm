package quicksort

func QuickSort(elements []int) {
	quickSort(elements, 0, len(elements)-1)
}

func quickSort(elements []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := partition2(elements, left, right)
	quickSort(elements, left, pivotIndex-1)
	quickSort(elements, pivotIndex+1, right)
}

// [left, right]
func partition(elements []int, left, right int) int {
	pivotIndex := left
	pivot := elements[left]
	sepIndex := pivotIndex + 1
	for i := pivotIndex + 1; i <= right; i++ {
		if elements[i] < pivot {
			elements[i], elements[sepIndex] = elements[sepIndex], elements[i]
			sepIndex++
		}
	}
	elements[pivotIndex], elements[sepIndex-1] = elements[sepIndex-1], elements[pivotIndex]
	return sepIndex - 1
}

// 夹逼 + 双指针 + 鸽笼原理
// 较难理解，较多细节，较易写错
func partition2(elements []int, left, right int) int {
	pivot := elements[left]
	emptyIndex := left
	for left < right {
		// 初始的时候，空格位置在最左边，所以需要先从右到左遍历，这样当找到小于或等于pivot的元素的时候（本应该位于基准元素的左边），可以将该元素放到空格位置
		for left < right && elements[right] > pivot {
			right--
		}
		elements[emptyIndex] = elements[right]        // 将找到的小于或等于pivot的元素放到空格位置
		emptyIndex = right                            // 找到的小于或等于pivot的元素位置成为新的空格位置
		for left < right && elements[left] <= pivot { // 必须从首个元素开始遍历。为防止特殊case：首个位置就是分隔位置
			left++
		}
		elements[emptyIndex] = elements[left] // 将找到的大于pivot的元素放到空格位置
		emptyIndex = left                     // 找到的大于pivot的元素位置成为新的空格位置
	}
	elements[emptyIndex] = pivot // 空格位置就是最后基准元素的位置
	return emptyIndex
}
