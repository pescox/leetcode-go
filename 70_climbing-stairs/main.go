package main

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 提示：
// 1 <= n <= 45

// 解题思路：简单的 DP，经典的爬楼梯问题。一个楼梯可以由 n-1 和 n-2 的楼梯爬上来。
// 这一题求解的值就是斐波那契数列。
func main() {

}

// func climbStairs(n int) int {
// 	dp := make([]int, n+1)
// 	dp[0], dp[1] = 1, 1
// 	for i := 2; i <= n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	return dp[n]
// }

// 空间稍作优化
func climbStairs(n int) int {
	a, b := 1, 1
	var c int
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
