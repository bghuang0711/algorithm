package reverse_by_group

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func buildList(dataList []int) *Node {
	dummy := &Node{}
	tail := dummy
	for _, data := range dataList {
		node := &Node{
			data: data,
		}
		tail.next = node
		tail = node
	}
	return dummy.next
}

func TestSolve(t *testing.T) {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expectedList := []int{1, 2, 5, 4, 3, 8, 7, 6}
	head := buildList(inputList)
	K := 3
	head = Solve(head, K)
	outputList := make([]int, 0)
	cur := head
	for cur != nil {
		outputList = append(outputList, cur.data)
		cur = cur.next
	}
	assert.Equalf(t, len(expectedList), len(outputList), "len(expectedList) != len(outputList)")
	t.Log(outputList)
	for i := 0; i < len(outputList); i++ {
		assert.Equalf(t, expectedList[i], outputList[i], fmt.Sprintf("expectedList[%d] != outputList[%d]", i, i))
	}
}
