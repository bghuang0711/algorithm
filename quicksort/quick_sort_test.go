package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		elements := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		QuickSort(elements)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, elements)
	})
}
