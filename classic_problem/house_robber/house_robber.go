package house_robber

// 打家劫舍
// https://leetcode.cn/problems/house-robber/
func HouseRobberDynamicProgramming(properties []int) int {
	// 令dp[i][0]表示0...i个房间不偷第i间的最大金额，dp[i][1]表示0...i个房间偷第i间的最大金额
	// 状态转移方程：
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	// dp[i][1] = dp[i-1][0] + properties[i]
	if len(properties) == 0 {
		return 0
	}
	unfetch := 0
	fetch := properties[0]
	for i := 1; i < len(properties); i++ {
		tmp := unfetch
		unfetch = Max(unfetch, fetch)
		fetch = tmp + properties[i]
	}
	return Max(unfetch, fetch)
}
