/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/19
	判断这个值是不是2的幂次
 */

package tencent

func isPowerOfTwo(n int) bool {
	return n > 0 && (n & (n - 1)) == 0
}