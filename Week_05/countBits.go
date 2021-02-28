package Week_05

// 布莱恩尼克根算法 ； 清除n的最右一位1的值  num &= (num - 1)
// dp思想 类似状态压缩DP   i 一定大于  i&(i-1)  dp[i] = dp[i & (i-1)] + 1
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}

/*
func countBits(num int) []int {
    res := make([]int, num+1)

     for i := 0 ; i <= num ; i++{
         count := 0
         tmp := i
         for tmp > 0{
             tmp &= tmp -1
             count ++
         }
         res[i] = count
     }
     return res
}
// num &= num - 1

// o(n)
*/
