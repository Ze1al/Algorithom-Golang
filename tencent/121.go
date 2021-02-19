/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	买卖股票的最佳时间
 */

package tencent

func maxProfit1(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i+1; j < len(prices); j++ {
			if prices[j] > prices[i] && (prices[j] - prices[i]) > res {
				res = prices[j] - prices[i]
			}
		}
	}
	return res
}



func maxProfit3(prices []int) int {
	minNum := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minNum {
			minNum = prices[i]
		} else {
			if (prices[i] - minNum) > res {
				res = prices[i] - minNum
			}
		}
	}
	return res
}