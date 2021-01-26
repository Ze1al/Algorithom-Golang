/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/26
	Pow(x, n)
 */

package Top


func myPow(x float64, n int) float64 {
	if n >= 0{
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1{
		return x
	}
	y := quickMul(x, n/2)
	if n % 2== 0 {
		return y * y
	}
	return y * y * x
}