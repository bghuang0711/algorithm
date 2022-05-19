package house_robber

import "testing"

func TestHouseRobberDynamicProgramming(t *testing.T) {
	tests := []struct {
		properties []int
	}{
		{
			[]int{1, 2, 3, 1},
		},
		{
			[]int{2, 7, 9, 3, 1},
		},
		{
			[]int{1, 1, 8, 10},
		},
	}
	for i, test := range tests {
		t.Log(i, HouseRobberDynamicProgramming(test.properties))
	}
}
