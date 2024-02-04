package selectsort

func SelectSort(elements []int) {
	n := len(elements)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if elements[j] < elements[minIndex] {
				minIndex = j
			}
		}
		elements[i], elements[minIndex] = elements[minIndex], elements[i]
	}
}
