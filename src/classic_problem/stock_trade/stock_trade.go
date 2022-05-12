package stock_trade

import "math"

// 股票买卖问题通解：
// 令dp[i][k][0]表示第i天剩余k次交易机会不持有股票的最大利润，dp[i][k][1]表示第i天剩余k次交易机会持有股票的最大利润
// 状态转移方程为：
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])
// 最大利润为：dp[n-1][k][0]
// 初始化：dp[0][k][0]=0 dp[0][k][1]=-prices[0]
// 参考
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/solution/yi-chong-jie-fa-tuan-mie-mai-mai-gu-piao-36px/

// 一次买入一次卖出
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func StockTrade0Brute(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				continue
			}
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func StockTrade0MonotonicStack(prices []int) int {
	maxProfit := 0
	prices = append(prices, -1) // sentinel
	stack := make([]int, 0)
	for i := 0; i < len(prices); i++ {
		if len(stack) > 0 && prices[i] < stack[len(stack)-1] {
			profit := stack[len(stack)-1] - stack[0]
			if profit > maxProfit {
				maxProfit = profit
			}
			for len(stack) > 0 && prices[i] < stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, prices[i])
	}
	return maxProfit
}

func StockTrade0TrackMinPrice(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}

func StockTrade0DynamicProgramming(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	unhold, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		unhold = Max(unhold, hold+prices[i])
		hold = Max(hold, -prices[i])
	}
	return unhold
}

// 多次买入多次卖出
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode-s/
func StockTrade1Greedy(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func StockTrade1DynamicProgramming(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	unhold, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		unhold = Max(unhold, hold+prices[i])
		hold = Max(hold, unhold-prices[i])
	}
	return unhold
}

// 限制最多2次交易
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
func StockTrade2DynamicProgramming(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	unhold_1, unhold_2 := 0, 0
	hold_1, hold_2 := -prices[0], -prices[0]
	for i := 1; i < len(prices); i++ {
		unhold_2 = Max(unhold_2, hold_2+prices[i])
		unhold_1 = Max(unhold_1, hold_1+prices[i])
		hold_2 = Max(hold_2, unhold_1-prices[i])
		hold_1 = Max(hold_1, -prices[i])
	}
	return unhold_2
}

// 限制最多交易k次
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
func StockTrade3DynamicProgramming(prices []int, k int) int {
	if len(prices) <= 1 {
		return 0
	}
	unhold := make([]int, k+1)
	hold := make([]int, k+1)
	for j := 0; j <= k; j++ {
		unhold[j] = 0
		hold[j] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			unhold[j] = Max(unhold[j], hold[j]+prices[i])
			hold[j] = Max(hold[j], unhold[j-1]-prices[i])
		}
	}
	return unhold[k]
}

// 限制冷冻期
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func StockTrade4DynamicProgramming(prices []int, freezeDays int) int {
	if len(prices) <= 1 {
		return 0
	}
	n := len(prices)
	unhold := make([]int, n)
	hold := make([]int, n)
	unhold[0] = 0
	hold[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		unhold[i] = Max(unhold[i-1], hold[i-1]+prices[i])
		tmp := 0
		if i-1 >= freezeDays {
			tmp = unhold[i-1-freezeDays]
		}
		hold[i] = Max(hold[i-1], tmp-prices[i])
	}
	return unhold[n-1]
}

// 买入或卖出收手续费
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
func StockTrade5DynamicProgramming(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}
	unhold, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		// 卖出收手续费
		unhold = Max(unhold, hold+prices[i]-fee)
		hold = Max(hold, unhold-prices[i])
		// 买入收手续费
		// unhold = Max(unhold, hold+prices[i])
		// hold = Max(hold, unhold-prices[i]-fee)
	}
	return unhold
}
