package stock_trade

import "testing"

func TestMinMax(t *testing.T) {
	t.Log(Min(1, 2))
	t.Log(Max(2, 1))
}

func TestStockTrade0Brute(t *testing.T) {
	pricesList := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}
	for i, prices := range pricesList {
		t.Log(i, StockTrade0Brute(prices))
	}
}

func TestStockTrade0MonotonicStack(t *testing.T) {
	pricesList := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
	}
	for i, prices := range pricesList {
		t.Log(i, StockTrade0MonotonicStack(prices))
	}
}

func TestStockTrade0TrackMinPrice(t *testing.T) {
	pricesList := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}
	for i, prices := range pricesList {
		t.Log(i, StockTrade0TrackMinPrice(prices))
	}
}

func TestStockTrade0DynamicProgramming(t *testing.T) {
	pricesList := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}
	for i, prices := range pricesList {
		t.Log(i, StockTrade0DynamicProgramming(prices))
	}
}

func TestStockTrade1Greedy(t *testing.T) {
	pricesList := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
	}
	for i, prices := range pricesList {
		t.Log(i, StockTrade1Greedy(prices))
	}
}

func TestStockTrade1DynamicProgramming(t *testing.T) {
	pricesList := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
	}
	for i, prices := range pricesList {
		t.Log(i, StockTrade1DynamicProgramming(prices))
	}
}

func TestStockTrade2DynamicProgramming(t *testing.T) {
	pricesList := [][]int{
		{3, 3, 5, 0, 0, 3, 1, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
		{},
	}
	for i, prices := range pricesList {
		t.Log(i, StockTrade2DynamicProgramming(prices))
	}
}

func TestStockTrade3DynamicProgramming(t *testing.T) {
	tests := []struct {
		prices []int
		k      int
	}{
		{
			[]int{2, 4, 1},
			2,
		},
		{
			[]int{3, 2, 6, 5, 0, 3},
			2,
		},
	}
	for i, test := range tests {
		t.Log(i, StockTrade3DynamicProgramming(test.prices, test.k))
	}
}

func TestStockTrade4DynamicProgramming(t *testing.T) {
	tests := []struct {
		prices     []int
		freezeDays int
	}{
		{
			[]int{1, 2, 3, 0, 2},
			1,
		},
		{
			[]int{1},
			1,
		},
	}
	for i, test := range tests {
		t.Log(i, StockTrade4DynamicProgramming(test.prices, test.freezeDays))
	}
}

func TestStockTrade5DynamicProgramming(t *testing.T) {
	tests := []struct {
		prices []int
		fee    int
	}{
		{
			[]int{1, 3, 2, 8, 4, 9},
			2,
		},
		{
			[]int{1, 3, 7, 5, 10, 3},
			3,
		},
	}
	for i, test := range tests {
		t.Log(i, StockTrade5DynamicProgramming(test.prices, test.fee))
	}
}
